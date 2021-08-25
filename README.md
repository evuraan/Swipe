# Swipe
Gestures on Linux. 

![Swipe](./images/Swipe_300x300.png)
<p>https://evuraan.info/Swipe/ 

 
## Features
Swipe uses a novel yet simple correlation mechanism to determine directional intent from event coordinates.
- Wide range of devices supported 
- Touchscreens and Touchpads 
- Wayland and X11 compatible
- Easy Installation - download (or git clone) and run. 
- No dependency on Python or Ruby
- Supports Config files
## Available variants/branches 
- Branch [modular](https://github.com/evuraan/Swipe/tree/modular) - Swipe/1.06e - Use with `xdotool` or `ydotool` or `xte` etc. 
 
 ## Screengrab Video:
 `Swipe` in action on a Linux laptop:<br>
  [https://evuraan.info/evuraan/stuff/Swipe.mp4](https://evuraan.info/evuraan/stuff/Swipe.mp4)
 
 
## Requirements 
- libinput-tools  
  ```bash 
    sudo apt-get install libinput-tools 
   ```
- Your user must be a member of `input` group:
  ```bash 
    sudo gpasswd -a $USER input
    newgrp input
    ```
- Optional: Enable touchpad events
  ```bash
    gsettings set org.gnome.desktop.peripherals.touchpad send-events enabled
     ```
## Setup
- Download `Swipe` to a folder. (Either download the latest [release](https://github.com/evuraan/Swipe/releases/download/1.06d/swipe), or clone this repo, or download and extract the <a href="https://github.com/evuraan/Swipe/archive/refs/heads/main.zip">Zip file</a>.)
- Launch `swipe` with an optional config file 

See your distro's documentation to setup `Swipe` as a [`Startup Application`](./images/Startup.png) - an [application](./images/Startup.png) that starts when a desktop user logs in. 

## Usage:

```bash
$ ./swipe -h
Usage: ./swipe
  -h  --help             print this usage and exit
  -v  --version          print version information and exit
  -s  --sampleCfg        show sample config
  -d  --debug            show verbose output
  -c  /etc/ku.conf       config file to use 
  -k  --keys             show available keys
  -i  /dev/input/event1  kbd device to use
  
```
- If no config file is specified, `Swipe` would use a default configuration. 
- If an appropriate `kbd` device cannot be found, `Swipe` will ask you to specify a suitable device using the `-i` option.

## Config
Generate a sample config file with  -s option:

```bash
$ ./swipe -s
Sample Config: 
# 3 Button Touchpad Gestures:
3right: "KEY_LEFTALT + KEY_LEFT"
3left:  "KEY_LEFTALT + KEY_RIGHT"
3up:    "KEY_SPACE"
3down:  "KEY_LEFTSHIFT + KEY_SPACE"

# 4 Button Touchpad Gestures:
4right: "KEY_LEFTALT + KEY_LEFT"
4left:  "KEY_LEFTALT + KEY_RIGHT"
4up:    "KEY_HOME"
4down:  "KEY_END"

# 5 - Touchscreens
5right:      "KEY_LEFTALT + KEY_LEFT"
5left:       "KEY_LEFTALT + KEY_RIGHT"
5up:         "KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN"
5mediumUp:   "KEY_SPACE"
5mediumDown: "KEY_LEFTSHIFT + KEY_SPACE"
5down:       "KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP"
5fastup:     "KEY_SPACE"
5fastdown:   "KEY_LEFTSHIFT + KEY_SPACE"

```
## Debug option
Run with `-d` option to have debug info onto the terminal:
```bash$ ./swipe -c swipe.conf -d 
$ ./swipe -d
$ ./swipe -d -c /tmp/swipe.conf 
Wed Aug 25 07:14:12 2021 Swipe/3.01c Read 12 values from the config file
Wed Aug 25 07:14:12 2021 Swipe/3.01c 3 key touchpad events: map[DOWN:KEY_LEFTSHIFT + KEY_SPACE LEFT:KEY_LEFTALT + KEY_RIGHT RIGHT:KEY_LEFTALT + KEY_LEFT UP:KEY_SPACE]
Wed Aug 25 07:14:12 2021 Swipe/3.01c 4 key touchpad events: map[DOWN:KEY_END LEFT:KEY_LEFTALT + KEY_RIGHT RIGHT:KEY_LEFTALT + KEY_LEFT UP:KEY_HOME]
Wed Aug 25 07:14:12 2021 Swipe/3.01c touchscreen events: map[DOWN:KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP FAST_DOWN:KEY_LEFTSHIFT + KEY_SPACE FAST_UP:KEY_SPACE LEFT:KEY_LEFTALT + KEY_RIGHT MED_DOWN:KEY_LEFTSHIFT + KEY_SPACE MED_UP:KEY_SPACE RIGHT:KEY_LEFTALT + KEY_LEFT UP:KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN]
Copyright © 2021 Evuraan <evuraan@gmail.com>. All rights reserved.
This program comes with ABSOLUTELY NO WARRANTY.
Wed Aug 25 07:14:12 2021 Swipe/3.01c Howdy!
Wed Aug 25 07:14:12 2021 Swipe/3.01c keyboard device: /dev/input/event3
Wed Aug 25 07:14:12 2021 [C] [getFd] fd opened: 3
Wed Aug 25 07:14:12 2021 Swipe/3.01c About to run stdbuf -oL /usr/libexec/libinput/libinput-debug-events
Wed Aug 25 07:14:12 2021 Swipe/3.01c pid: 9145
Wed Aug 25 07:14:21 2021 Swipe/3.01c left: 0 right: 0 up: 0 down: 0
Wed Aug 25 07:14:21 2021 Swipe/3.01c movedTo: RIGHT
Wed Aug 25 07:14:21 2021 [C] [emit] emitted 24 bytes type 1 code 56
Wed Aug 25 07:14:21 2021 [C] [emit] emitted 24 bytes type 1 code 105
Wed Aug 25 07:14:21 2021 [C] [emit] emitted 24 bytes type 0 code 0
Wed Aug 25 07:14:21 2021 [C] [emit] emitted 24 bytes type 1 code 56
Wed Aug 25 07:14:21 2021 [C] [emit] emitted 24 bytes type 1 code 105
Wed Aug 25 07:14:21 2021 [C] [emit] emitted 24 bytes type 0 code 0
Wed Aug 25 07:14:21 2021 [C] [handleEvents] Handled 2 events
Wed Aug 25 07:14:21 2021 Swipe/3.01c Gesture type 3, intent: RIGHT, cmd: KEY_LEFTALT + KEY_LEFT
Wed Aug 25 07:14:23 2021 Swipe/3.01c touchLen: 3
Wed Aug 25 07:14:23 2021 Swipe/3.01c movedTo: DOWN
Wed Aug 25 07:14:23 2021 Swipe/3.01c fingers: 1
Wed Aug 25 07:14:23 2021 Swipe/3.01c startx :59.95, endx: 59.98
Wed Aug 25 07:14:23 2021 Swipe/3.01c starty: 60.01, endy: 60.34
Wed Aug 25 07:14:23 2021 Swipe/3.01c xdelta: 0.02999999999999403, ydelta: 0.3300000000000054, abs(xd): 0.02999999999999403, abs(yd): 0.3300000000000054
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 0 code 0
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 1 code 103
Wed Aug 25 07:14:23 2021 [C] [emit] emitted 24 bytes type 0 code 0
Wed Aug 25 07:14:23 2021 [C] [handleEvents] Handled 6 events
Wed Aug 25 07:14:23 2021 Swipe/3.01c Gesture type 5, intent: DOWN, cmd: KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP


```
Desktop notifications ([example](./images/Debug.png)) are also enabled in debug mode - which shows the details of the event intercepted. 

## Keys and buttons supported:
Swipe supports about `482` keys/buttons - pretty much inline with Linux's [input-event-codes.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h). <br>  
Run `swipe -k` to see a full list:
```bash
$ ./swipe -k
Available keys:
key -->  KEY_FN_D
key -->  KEY_BRL_DOT8
key -->  KEY_HANJA
key -->  KEY_FILE
key -->  KEY_PHONE
key -->  KEY_ATTENDANT_ON
key -->  KEY_MACRO_PRESET1
key -->  KEY_KP5
key -->  KEY_PAGEUP
key -->  KEY_RIGHT
key -->  KEY_PRESENTATION
key -->  KEY_KBDINPUTASSIST_NEXT
key -->  KEY_FASTREVERSE
key -->  KEY_KP1
<snip>  
```

## Optional: Build 
If you prefer to build yourself, you will need the [Go Programming Language](https://golang.org/dl/) installed on your System. 

Go into the folder and build as: 
``` 
go build
```
