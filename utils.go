/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func checkExec(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func doRun(cmdIn string) error {

	if len(cmdIn) < 1 {
		err := errors.New("cmdIn len 0")
		return err
	}

	safetySplat := strings.Split(cmdIn, " ")
	cmdSplat := []string{}
	x := 0
	for i := range safetySplat {
		block := safetySplat[i]
		if block != "" {
			cmdSplat = append(cmdSplat, block)
			x++
		}
	}
	if x < 1 {
		err := errors.New("splat x 0")
		return err
	}

	cmd := exec.Command(cmdSplat[0], cmdSplat[1:]...) // #nosec G204
	err := cmd.Run()
	return err
}

// this has some special fu
func runThis(cmdIn string) error {

	if len(cmdIn) < 1 {
		err := errors.New("cmdIn len 0")
		return err
	}

	print("About to run %s", cmdIn)
	safetySplat := strings.Split(cmdIn, " ")
	cmdSplat := []string{}
	x := 0
	for i := range safetySplat {
		block := safetySplat[i]
		if block != "" {
			cmdSplat = append(cmdSplat, block)
			x++
		}
	}
	if x < 1 {
		err := errors.New("splat x 0")
		return err
	}

	cmd := exec.Command(cmdSplat[0], cmdSplat[1:]...) // #nosec G204
	cmd.Stderr = os.Stderr
	var err error
	stdout, err = cmd.StdoutPipe()

	if err != nil {
		return err
	}

	err = cmd.Start()

	if err == nil {
		print("pid: %v", cmd.Process.Pid)
		orange := orangeStruct{}
		orange.processLoop()
		// getOutput()
		go func() {
			_ = cmd.Wait()
		}()
	}

	return err
}

func swipeProcessor(chanPtr *chan string) {
	if chanPtr == nil {
		return
	}

	ourMoves := make(map[int]string)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	var incoming string
	evtType := 3
	for incoming = range *chanPtr {
		if len(incoming) < 1 {
			continue
		}
		gesture := incoming

		switch {
		case strings.Contains(gesture, swipeUpdate):
			wg.Add(1)

			go func() {
				defer wg.Done()
				gesture = strings.ReplaceAll(gesture, "(", "")
				gesture = strings.ReplaceAll(gesture, "/ ", "/")

				splat := strings.Split(gesture, " ")
				if len(splat) < 1 {
					return
				}

				map1 := make(map[int]string)
				for i := range splat {
					if strings.Contains(splat[i], "/") {
						map1[len(map1)] = splat[i]
					}
				}
				thisMove := getMoves(&map1)
				moove := thisMove.analyze()

				fourth := splat[4]

				mu.Lock()
				defer mu.Unlock()

				if len(fourth) > 0 && fourth[len(fourth)-1] == '4' {
					// change to 4 finger swipes
					evtType = 4
				}

				if len(moove) > 1 {
					ourMoves[len(ourMoves)] = moove
				}
			}()

		case strings.Contains(gesture, swipeEnd):
			rCt := 0
			lCt := 0
			uCt := 0
			dCt := 0

			wg.Wait()
			for k := range ourMoves {
				switch {
				case ourMoves[k] == up:
					uCt++
				case ourMoves[k] == down:
					dCt++
				case ourMoves[k] == left:
					lCt++
				case ourMoves[k] == right:
					rCt++
				}
			}

			print("left: %d right: %d up: %d down: %d", lCt, rCt, uCt, dCt)

			movedTo := ""
			switch {
			case rCt >= lCt && rCt >= uCt && rCt >= dCt:
				movedTo = right
			case lCt >= rCt && lCt >= uCt && lCt >= dCt:
				movedTo = left
			case uCt >= lCt && uCt >= rCt && uCt >= dCt:
				movedTo = up
			case dCt >= uCt && dCt >= rCt && dCt >= lCt:
				movedTo = down
			}

			go eventLibStuff.handleEvent(movedTo, evtType)
			if deBug && notifyBool {
				letsNotify := fmt.Sprintf("%s %s\n%d", notifyCmd, movedTo, time.Now().Local().Unix())
				workChan <- letsNotify
			}
			print("movedTo: %s", movedTo)
			return
		}

	}
}

func getMoves(mapPtr *map[int]string) *moves {
	if mapPtr == nil {
		return nil
	}

	cobra := *mapPtr
	lhs := cobra[0]
	rhs := cobra[1]

	if len(rhs) < 1 || len(lhs) < 1 {
		return nil
	}

	lhsSplat := strings.Split(lhs, "/")
	rhsSplat := strings.Split(rhs, "/")
	if len(lhsSplat) < 1 || len(rhsSplat) < 1 {
		return nil
	}
	if len(lhsSplat[0]) < 1 || len(lhsSplat[1]) < 1 || len(rhsSplat[0]) < 1 || len(rhsSplat[1]) < 1 {
		return nil
	}

	var thisMove moves
	var err error
	thisMove.a, err = strconv.ParseFloat(lhsSplat[0], 64)
	if err != nil {
		return nil
	}

	thisMove.b, err = strconv.ParseFloat(lhsSplat[1], 64)
	if err != nil {
		return nil
	}

	thisMove.c, err = strconv.ParseFloat(rhsSplat[0], 64)
	if err != nil {
		return nil
	}

	thisMove.c, err = strconv.ParseFloat(rhsSplat[1], 64)
	if err != nil {
		return nil
	}
	return &thisMove
}

func (movesPtr *moves) analyze() string {
	x := ""
	if movesPtr == nil {
		return x
	}

	self := movesPtr
	switch {
	case self.a <= 0 && self.b > 2:
		x = down
	case self.a < 1 && self.a > 0 && self.b < 0:
		x = up
	case self.a > 0 && self.b < self.a && self.b > 0:
		x = right
	case self.a < 0 && self.b < 0 && self.b > self.a:
		x = left

	}
	return x
}

func getAbs(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}

func print(strings string, args ...interface{}) {
	if !deBug {
		return
	}
	a := time.Now()
	msg := fmt.Sprintf(strings, args...)
	if msg[len(msg)-1] == '\n' {
		fmt.Print(a.Format(layout), tag, msg)
	} else {
		fmt.Println(a.Format(layout), tag, msg)
	}
}

func checkFile(fileName string) bool {
	return (getFileSize(fileName) > 0)
}

func getFileSize(fileName string) int64 {
	fi, err := os.Stat(fileName)
	if err != nil {
		return 0
	}
	return fi.Size()
}

func parseConfig(configFile string) {
	someDict := make(map[string]string)

	func() {
		f, err := os.Open(filepath.Clean(configFile))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open config file: %v\n", err)
			os.Exit(1)
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "err: %s", err)
			}
		}()

		fscanner := bufio.NewScanner(f)
		for fscanner.Scan() {
			line := fscanner.Text()
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "\"", "")
			if len(line) < 1 {
				continue
			}
			if strings.HasPrefix(line, "#") {
				continue
			}
			line = strings.ToUpper(line)

			splat := strings.Split(line, ":")
			if len(splat) < 2 {
				continue
			}
			key := strings.TrimSpace(splat[0])
			val := strings.TrimSpace(splat[1])
			if len(key) < 1 || len(val) < 1 {
				continue
			}
			someDict[key] = val
		}
	}()

	// if this fn was evoked, we expect someDict to be populated.
	if len(someDict) < 1 {
		fmt.Fprintf(os.Stderr, "Failed to parse %s\n", configFile)
		os.Exit(1)
	}

	evt2 = make(map[string]string)
	evt3 = make(map[string]string)
	evt4 = make(map[string]string)
	evt5 = make(map[string]string)

	j := 0
	for i := range directions {
		lookFor3 := "3" + directions[i]
		key := directions[i]

		x, ok := someDict[lookFor3]
		if ok {
			if len(x) > 1 {
				evt3[key] = x
				j++
			}
		}

		lookFor4 := "4" + directions[i]
		y, ok := someDict[lookFor4]
		if ok {
			if len(y) > 1 {
				evt4[key] = y
				j++
			}
		}

		// 5: Touchscreen events
		lookFor5 := "5" + directions[i]
		touchScreen, ok := someDict[lookFor5]
		if ok {
			if len(touchScreen) > 1 {
				evt5[key] = touchScreen
				j++
			}
		}

		// evt2  2 finger touchPad events
		lookFor2 := "2" + directions[i]
		twoPad, ok := someDict[lookFor2]
		if ok {
			if len(twoPad) > 1 {
				evt2[key] = twoPad
				j++
			}
		}

	}
	print("Read %d values from the config file", j)
	print("2 key touchpad events: %v", evt2)
	print("3 key touchpad events: %v", evt3)
	print("4 key touchpad events: %v", evt4)
	print("touchscreen events: %v", evt5)
}
