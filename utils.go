package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
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

	cmd := exec.Command(cmdSplat[0], cmdSplat[1:]...)
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

	cmd := exec.Command(cmdSplat[0], cmdSplat[1:]...)
	cmd.Stderr = os.Stderr
	var err error
	stdout, err = cmd.StdoutPipe()

	if err != nil {
		return err
	}

	err = cmd.Start()

	if err == nil {
		print("pid: %v", cmd.Process.Pid)
		getOutput()
		go cmd.Wait()
	}

	return err
}

func getOutput() {
	scanner := bufio.NewScanner(stdout)
	m := ""
	var chanToUse *chan string = nil
	for scanner.Scan() {

		m = scanner.Text()
		switch {
		case strings.Contains(m, swipeStart):
			swipeChan := make(chan string, procWidth)
			chanToUse = &swipeChan
			go swipeProcessor(chanToUse)
		case strings.Contains(m, swipeEnd):
			if chanToUse != nil {
				// tricky stuff: send one last time and then close it.
				// also mark the ptr as nil to avoid fireworks.
				*chanToUse <- m
				close(*chanToUse)
				chanToUse = nil
			}
		}

		if chanToUse != nil {
			*chanToUse <- m
		}

	}

}

func swipeProcessor(chanPtr *chan string) {
	if chanPtr == nil {
		return
	}

	ourMoves := make(map[int]string)
	lookUpMap := swipes3
	mu := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	gesture := "swipe"

	for gesture = range *chanPtr {
		if len(gesture) < 1 {
			continue
		}
		jinjin := gesture

		switch {
		case strings.Contains(jinjin, swipeUpdate):
			wg.Add(1)

			go func() {
				defer wg.Done()
				jinjin = strings.ReplaceAll(jinjin, "(", "")
				jinjin = strings.ReplaceAll(jinjin, "/ ", "/")

				splat := strings.Split(jinjin, " ")
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
					lookUpMap = swipes4
				}

				if len(moove) > 1 {
					ourMoves[len(ourMoves)] = moove
				}
			}()

		case strings.Contains(jinjin, swipeEnd):
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

			print("movedTo: %s", movedTo)
			cmd, ok := lookUpMap[movedTo]
			if ok {
				workChan <- cmd
				if deBug && notifyBool {
					dornotify := fmt.Sprintf("%s %s\n%d", notifyCmd, movedTo, time.Now().Local().Unix())
					workChan <- dornotify
				}
				return
			}
			print("not sure how to process %s", movedTo)
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

// the wonkiest part: we have 4 float64s, which way was it?
// observe /usr/libexec/libinput/libinput-debug-events to increase reliability.
func (movesPtr *moves) analyze() string {
	x := ""
	if movesPtr == nil {
		return x
	}

	self := movesPtr
	switch {
	case self.a <= 0 && self.b > 2:
		x = down

	/* up:
	self: &main.moves{a:0, b:-0.59, c:-2.19, d:0}
	self: &main.moves{a:0, b:-0.88, c:-3.28, d:0}
	self: &main.moves{a:0, b:-0.22, c:-1.09, d:0}
	self: &main.moves{a:0, b:-0.37, c:-2.19, d:0}
	self: &main.moves{a:0, b:-0.18, c:-1.09, d:0}
	self: &main.moves{a:0, b:-0.18, c:-1.09, d:0}
	*/
	case self.a < 1 && self.a > 0 && self.b < 0:
		x = up

	/*right:
	self: &main.moves{a:4.98, b:0.59, c:2.19, d:0}
	self: &main.moves{a:7.04, b:0.59, c:2.19, d:0}
	self: &main.moves{a:4.98, b:0.88, c:3.28, d:0}
	self: &main.moves{a:6.74, b:0.59, c:2.19, d:0}
	self: &main.moves{a:7.62, b:0.59, c:2.19, d:0}
	self: &main.moves{a:5.86, b:0.29, c:1.09, d:0}
	self: &main.moves{a:8.79, b:0.29, c:1.09, d:0}
	self: &main.moves{a:2.93, b:0.59, c:2.19, d:0}
	*/
	case self.a > 0 && self.b < self.a && self.b > 0:
		x = right

	/* left:
		self: &main.moves{a:-3.81, b:-0.88, c:-3.28, d:0}
	self: &main.moves{a:-3.22, b:-1.47, c:-5.47, d:0}
	self: &main.moves{a:-2.64, b:-0.88, c:-3.28, d:0}
	self: &main.moves{a:-2.64, b:-1.17, c:-4.37, d:0}
	self: &main.moves{a:-2.35, b:-1.17, c:-4.37, d:0}
	self: &main.moves{a:-2.64, b:-1.47, c:-5.47, d:0}
	self: &main.moves{a:-1.47, b:-0.29, c:-1.09, d:0}
	self: &main.moves{a:-1.47, b:-0.59, c:-2.19, d:0}
	self: &main.moves{a:-1.47, b:-0.59, c:-2.19, d:0}
	self: &main.moves{a:-1.17, b:-0.29, c:-1.09, d:0}
	self: &main.moves{a:-0.88, b:-0.29, c:-1.09, d:0}
	*/
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
		f, err := os.Open(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open config file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
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

	j := 0
	for i := range directions {
		lookFor3 := "3" + directions[i]
		key := directions[i]

		x, ok := someDict[lookFor3]
		if ok {
			if len(x) > 1 {
				swipes3[key] = x
				j++
			}
		}

		lookFor4 := "4" + directions[i]
		y, ok := someDict[lookFor4]
		if ok {
			if len(y) > 1 {
				swipes4[key] = y
				j++
			}
		}
	}
	print("Read %d values from the config file", j)
	return
}
