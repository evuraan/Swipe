/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func touchProcessor(chanPtr *chan string) {
	if chanPtr == nil {
		return
	}
	var incoming string
	touchEvents := make(map[int]string)
	fingers := 0

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for incoming = range *chanPtr {
		if !strings.Contains(incoming, "TOUCH_MOTION") {
			continue
		}
		incoming := incoming
		wg.Add(1)
		go func() {
			defer wg.Done()
			if strings.Contains(incoming, "(0)") {
				splat := strings.Split(incoming, " ")
				for i := range splat {
					if strings.Contains(splat[i], "/") {
						mu.Lock()
						touchEvents[len(touchEvents)] = splat[i]
						mu.Unlock()
						break
					}
				}
			}
			mu.Lock()
			if strings.Contains(incoming, "(1)") && fingers == 0 {
				fingers++
			} else if strings.Contains(incoming, "(2)") && fingers == 1 {
				fingers++
			} else if strings.Contains(incoming, "(3)") && fingers == 2 {
				fingers++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	touchLen := len(touchEvents)
	go print("touchLen: %d", touchLen)
	if touchLen < touchMin {
		print("Dropping event, less than touchMin")
		return
	}
	startx, starty := getCounters(touchEvents[0])
	endx, endy := getCounters(touchEvents[1])
	xdelta := endx - startx
	ydelta := endy - starty
	xdAbs := getAbs(xdelta)
	ydAbs := getAbs(ydelta)

	go func() {
		print("startx :%v, endx: %v", startx, endx)
		print("starty: %v, endy: %v", starty, endy)
		print("xdelta: %v, ydelta: %v, abs(xd): %v, abs(yd): %v", xdelta, ydelta, xdAbs, ydAbs)
	}()

	fingers++
	touchEvent := fmt.Sprintf("TOUCH%d", fingers)
	switch {
	case xdelta >= ydelta && ydAbs >= xdAbs:
		touchEvent += "UP"
	case ydelta >= xdelta && xdAbs >= ydAbs:
		touchEvent += "LEFT"
	case xdelta >= ydelta && xdAbs >= ydAbs:
		touchEvent += "RIGHT"
	default:
		touchEvent += "DOWN"

	}

	const evtType = 5
	go eventLibStuff.handleEvent(touchEvent, evtType)
	if deBug && notifyBool {
		workChan <- fmt.Sprintf("%s %s\n%d", notifyCmd, touchEvent, time.Now().Local().Unix())
	}

	print("touchEvent: %s", touchEvent)
	print("fingers: %d", fingers)

}

func getCounters(touchy string) (x, y float64) {
	splat := strings.Split(touchy, "/")
	if len(splat) != 2 {
		return
	}
	var err error
	x, err = strconv.ParseFloat(splat[0], 64)
	if err != nil {
		return
	}

	y, err = strconv.ParseFloat(splat[1], 64)
	if err != nil {
		return
	}
	return
}
