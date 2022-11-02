package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	POINTER_AXIS_DELAY = 100
)

type orangeStruct struct {
	sync.RWMutex
	progress bool
	chanp    *chan string
}

func (orangeStructPtr *orangeStruct) getNewChan() (x bool) {
	self := orangeStructPtr
	someChan := make(chan string, procWidth)
	self.Lock()
	defer self.Unlock()
	self.chanp = &someChan
	return self.chanp != nil
}

// Seen in October 2022, ubuntu 22.04
func (orangeStructPtr *orangeStruct) twoFingerOctober2022() {
	self := orangeStructPtr
	self.RLock()
	readFromPtr := self.chanp
	self.RUnlock()
	if readFromPtr == nil {
		return
	}

	// read just enough to read the intent
	go func() {
		time.Sleep(oct2FinDelay)
		self.Lock()
		defer self.Unlock()
		if self.chanp != nil {
			close(*self.chanp)
			self.chanp = nil
		}
		self.progress = false
	}()

	done := make(chan bool, 1)
	defer close(done)
	go func() {

		ourMoves := make(map[int]string)

		for i := range *readFromPtr {
			if !strings.Contains(i, octoberTwoFin) {
				return
			}
			splat := strings.Split(i, " ")
			map1 := make(map[int]string)
			for i := range splat {
				if splat[i] == "" {
					continue
				}
				map1[len(map1)] = splat[i]
			}
			if len(map1) < 7 {
				continue
			}
			moveA, ok := getFloatTwoFinger(map1[3])
			if !ok {
				continue
			}
			moveB, ok := getFloatTwoFinger(map1[5])
			if !ok {
				continue
			}
			x := ""
			switch {
			case moveA < 0 && moveB == 0:
				x = up
			case moveA == 0 && moveB > 0:
				x = right
			case moveA == 0 && moveB < 0:
				x = left
			default:
				x = down
			}
			ourMoves[len(ourMoves)] = x

		}

		rCt := 0
		lCt := 0
		uCt := 0
		dCt := 0

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

		// print("left: %d right: %d up: %d down: %d", lCt, rCt, uCt, dCt)

		movedTo := ""
		switch {
		case rCt >= lCt && rCt >= uCt && rCt >= dCt:
			movedTo = right
		case lCt >= rCt && lCt >= uCt && lCt >= dCt:
			movedTo = left
		case uCt >= lCt && uCt >= rCt && uCt >= dCt:
			movedTo = "Up - Muted"
		case dCt >= uCt && dCt >= rCt && dCt >= lCt:
			movedTo = "Down - Muted"
		}
		print("%s intent: %s", octoberTwoFin, movedTo)
		if len(movedTo) < 1 {
			return
		}
		const evtType = 2
		go eventLibStuff.handleEvent(movedTo, evtType)
		if deBug && notifyBool {
			letsNotify := fmt.Sprintf("%s %s-%s\n%d", notifyCmd, "2FTPad", movedTo, time.Now().Local().Unix())
			workChan <- letsNotify
		}

	}()

}

// 2 finger touchPad
func (orangeStructPtr *orangeStruct) processPointerAxis() {
	move := POINTER_AXIS
	self := orangeStructPtr
	self.RLock()
	readFromPtr := self.chanp
	self.RUnlock()

	if readFromPtr == nil {
		return
	}
	readFrom := *readFromPtr
	done := make(chan bool, 1)
	defer close(done)
	go func() {

		ourMoves := make(map[int]string)

		for i := range readFrom {
			splat := strings.Split(i, " ")
			map1 := make(map[int]string)
			for i := range splat {
				if splat[i] == "" {
					continue
				}
				map1[len(map1)] = splat[i]
			}
			if len(map1) < 7 {
				continue
			}
			moveA, ok := getFloatTwoFinger(map1[3])
			if !ok {
				continue
			}
			moveB, ok := getFloatTwoFinger(map1[5])
			if !ok {
				continue
			}
			x := ""
			switch {
			case moveA < 0 && moveB == 0:
				x = up
			case moveA == 0 && moveB > 0:
				x = right
			case moveA == 0 && moveB < 0:
				x = left
			default:
				x = down
			}
			ourMoves[len(ourMoves)] = x

		}

		rCt := 0
		lCt := 0
		uCt := 0
		dCt := 0

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

		// print("left: %d right: %d up: %d down: %d", lCt, rCt, uCt, dCt)

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
		print("processPointerMotion intent: %s", movedTo)
		if len(movedTo) < 1 {
			return
		}

		const evtType = 2
		go eventLibStuff.handleEvent(movedTo, evtType)
		if deBug && notifyBool {
			letsNotify := fmt.Sprintf("%s %s-%s\n%d", notifyCmd, "2FTPad", movedTo, time.Now().Local().Unix())
			workChan <- letsNotify
		}

	}()

	select {
	case <-done:
		return
	case <-time.After(POINTER_AXIS_DELAY * time.Millisecond):
		print("Ding ding ding! done with %s", move)
		self.Lock()
		defer self.Unlock()
		if self.chanp != nil {
			self.progress = false
			close(*self.chanp)
			self.chanp = nil
		}
		return
	}
}

func getFloatTwoFinger(input string) (float64, bool) {
	if len(input) < 1 {
		return 0, false
	}
	splat := strings.Split(input, "/")
	someFloat, err := strconv.ParseFloat(splat[0], 64)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "strconv %v\n", err)
		fmt.Printf("Inc: %s splat: %#v\n", input, splat)
		return 0, false
	}
	return someFloat, true
}

func (orangeStructPtr *orangeStruct) processLoop() {
	self := orangeStructPtr
	scanner := bufio.NewScanner(stdout)
	m := ""
	var progress bool

	for scanner.Scan() {
		m = scanner.Text()
		if len(m) < 1 {
			continue
		}

		self.RLock()
		progress = self.progress
		self.RUnlock()

		switch {
		case !progress && strings.Contains(m, POINTER_AXIS):
			nChanPro := self.getNewChan()
			self.Lock()
			self.progress = nChanPro
			self.Unlock()
			go self.processPointerAxis()
		// POINTER_MOTION one finger touchPad
		// we ignore those.
		// case !progress && strings.Contains(m, POINTER_MOTION):
		// 	nChanPro := self.getNewChan()
		// 	self.Lock()
		// 	self.progress = nChanPro
		// 	self.Unlock()
		// 	go self.processPointerMotion()
		case !progress && strings.Contains(m, swipeStart):
			nChanPro := self.getNewChan()
			self.Lock()
			self.progress = nChanPro
			self.Unlock()
			go swipeProcessor(self.chanp)
		case strings.Contains(m, swipeEnd):
			self.Lock()
			if self.chanp != nil {
				*self.chanp <- m
				close(*self.chanp)
				self.chanp = nil
			}
			self.progress = false
			self.Unlock()
		case strings.Contains(m, touchStart):
			nChanPro := self.getNewChan()
			self.Lock()
			self.progress = nChanPro
			self.Unlock()
			go touchProcessor(self.chanp)
		case strings.Contains(m, touchEnd):
			self.Lock()
			if self.chanp != nil {
				*self.chanp <- m
				close(*self.chanp)
				self.chanp = nil
			}
			self.progress = false
			self.Unlock()
		case !progress && strings.Contains(m, octoberTwoFin):
			// a new approach: start reading upon the first signs of the signature.
			// read just enough to understand the intent
			nChanPro := self.getNewChan()
			self.Lock()
			self.progress = nChanPro
			self.Unlock()
			go self.twoFingerOctober2022()

			/* intercept these
			   -event5   POINTER_SCROLL_FINGER   +11.170s	vert 0.00/0.0 horiz 9.67/0.0* (finger)
			    event5   POINTER_SCROLL_FINGER   +11.177s	vert 0.00/0.0 horiz 3.51/0.0* (finger)
			    event5   POINTER_SCROLL_FINGER   +11.183s	vert 0.00/0.0 horiz 10.54/0.0* (finger)
			    event5   POINTER_SCROLL_FINGER   +11.190s	vert 0.00/0.0 horiz 6.59/0.0* (finger)
			*/
		}

		self.Lock()
		if self.progress && self.chanp != nil {
			*self.chanp <- m
		}
		self.Unlock()
	}
}
