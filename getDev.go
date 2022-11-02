/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	devices  = "/proc/bus/input/devices"
	keyboard = "keyboard"
)



// func main() {
// 	fmt.Println("Hello!")
// 	x := getDeviceForPattern("Integrated")
// 	fmt.Printf("x: %#v\n", x)
// }

func getDeviceForPattern(lookForA string) (devicePath string) {
	if len(lookForA) < 1 {
		return
	}
	lookFor := strings.ToLower(lookForA)
	fmap := parseFileToMap(devices)
	if len(fmap) < 1 {
		return
	}
	var entry, lentr string
	pat := regexp.MustCompile(`event\d*`)
	for k := range fmap {
		entry = fmap[k]
		if len(entry) < 1 {
			continue
		}
		lentr = strings.ToLower(entry)
		if strings.Contains(lentr, lookFor) {
			pos := k + 4
			if len(fmap) < pos {
				continue
			}
			handler := fmap[pos]
			s := pat.FindString(handler)

			if len(s) > 0 {
				devicePath = fmt.Sprintf("/dev/input/%s", s)
				return devicePath
			}
		}
	}

	return devicePath
}

func parseFileToMap(configFile string) (someDict map[int]string) {
	someDict = make(map[int]string)

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
			if len(line) > 0 {
				someDict[len(someDict)] = line
			}
		}
	}()

	return someDict
}
