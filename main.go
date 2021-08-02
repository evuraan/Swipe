package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	stdout     io.ReadCloser
	mustHave   = []string{"/usr/libexec/libinput/libinput-debug-events", "libinput-debug-events"}
	deBug      = false
	notifyBool = false
	directions = []string{up, down, left, right}
	workChan   chan string
	configFile = ""
	swipes3    = map[string]string{
		down:  "xdotool key Home",
		up:    "xdotool key End",
		left:  "xdotool key alt+Right",
		right: "xdotool key alt+Left",
	}
	swipes4 = map[string]string{
		left:  "xdotool key ctrl+alt+Right",
		right: "xdotool key ctrl+alt+Left",
		up:    "xdotool key ctrl+alt+Down",
		down:  "xdotool key ctrl+alt+Up",
	}
)

const (
	progName    = "Swipe"
	ver         = "1.06d"
	stdBuf      = "stdbuf"
	swipeStart  = "GESTURE_SWIPE_BEGIN"
	swipeUpdate = "GESTURE_SWIPE_UPDATE"
	swipeEnd    = "GESTURE_SWIPE_END"
	pinchStart  = "GESTURE_PINCH_BEGIN"
	pinchEnd    = "GESTURE_PINCH_END"
	pinchUpdate = "GESTURE_PINCH_UPDATE"
	up          = "UP"
	down        = "DOWN"
	left        = "LEFT"
	right       = "RIGHT"
	tag         = progName + "/" + ver
	xdotool     = "xdotool"
	layout      = "Mon Jan 02 15:04:05 2006"
	MAXWORKERS  = 10
	procWidth   = 20
	notifyCmd   = "notify-send " + progName

	sampleConf = `
# lines starting with # are ignored. 

# 3 Button Gesture Mappping:
3up:    "xdotool key End"
3down:  "xdotool key Home"
3left:  "xdotool key alt+Right"
3right: "xdotool key alt+Left"

# 4 Button Gestures:
4left:  "xdotool key ctrl+alt+Right"
4right: "xdotool key ctrl+alt+Left"
4up:    "xdotool key ctrl+alt+Down"
4down:  "xdotool key ctrl+alt+Up"
	`
)

type moves struct {
	a float64
	b float64
	c float64
	d float64
}

func main() {
	parseArgs()
	print("Howdy!")

	workChan = make(chan string, 2)
	go func() {
		for cmdString := range workChan {
			if len(cmdString) > 0 {
				go doRun(cmdString)
			}
		}
	}()

	libinput()
	fmt.Println("Bye bye!")

}

func libinput() {

	if !checkExec(xdotool) {
		fmt.Fprintf(os.Stderr, "Error: Could not find %s\n", xdotool)
		os.Exit(1)
	}
	launcher := ""
	stdBufBool := checkExec(stdBuf)
	notifyBool = checkExec("notify-send")

	for i := range mustHave {
		if checkExec(mustHave[i]) {
			launcher = mustHave[i]
			if stdBufBool {
				launcher = fmt.Sprintf("%s -oL %s", stdBuf, mustHave[i])
			}
			break
		}
	}

	if len(launcher) < 1 {
		fmt.Fprint(os.Stderr, "Error: Cannot find libinput debug command\n")
		os.Exit(1)
	}

	err := runThis(launcher)
	print("err :%v", err)
}

func parseArgs() {
	argc := len(os.Args)
	if argc > 1 {
		for i, arg := range os.Args {
			if strings.Contains(arg, "help") || arg == "h" || arg == "--h" || arg == "-h" || arg == "?" {
				showhelp()
				os.Exit(0)
			}
			if strings.Contains(arg, "sampleCfg") || arg == "s" || arg == "--s" || arg == "-s" {
				fmt.Println("\nSample Config:", sampleConf)
				os.Exit(0)
			}
			if strings.Contains(arg, "version") || arg == "v" || arg == "--v" || arg == "-v" {
				fmt.Println("Version:", tag)
				os.Exit(0)
			}
			if strings.Contains(arg, "debug") || arg == "d" || arg == "--d" || arg == "-d" {
				deBug = true
			}
			if arg == "-c" {
				nextArg := i + 1
				if argc > nextArg {
					configFile = os.Args[nextArg]
					if len(configFile) < 1 {
						fmt.Println("Invalid usage")
						showhelp()
						os.Exit(1)
					}
					if !checkFile(configFile) {
						fmt.Fprintf(os.Stderr, "Could not find config file %s\n", configFile)
						os.Exit(1)
					}
					parseConfig(configFile)
					// our config maps must be present.
					if len(swipes3) < 1 || len(swipes4) < 1 {
						fmt.Fprint(os.Stderr, "Config maps are empty. Fatal!")
						os.Exit(1)
					}
				} else {
					fmt.Println("Invalid usage")
					showhelp()
					os.Exit(1)
				}
			}
		}
	}

}

func showhelp() {
	fmt.Printf("Usage: %s\n", os.Args[0])
	fmt.Println("  -h  --help         print this usage and exit")
	fmt.Println("  -v  --version      print version information and exit")
	fmt.Println("  -s  --sampleCfg    show sample config")
	fmt.Println("  -d  --debug        show verbose output")
	fmt.Println("  -c  /etc/ku.conf   config file to use ")
}
