# Swipe
Gestures on Linux. 

![Swipe](./images/Swipe_300x300.png)
<p>https://evuraan.info/Swipe/ 

## Features
Swipe uses a novel yet simple correlation mechanism to determine directional intent from event coordinates.
- Easy Installation - download (or git clone) and run. 
- No dependency on Python or Ruby
- Supports Config files
- Wide range of devices supported. 
- Wayland and X11 compatible
## Available variants/branches 
- Branch [modular](https://github.com/evuraan/Swipe/tree/modular) - Swipe/1.06e - Use with `xdotool` or `ydotool` or `xte` etc. 
- Branch [main](https://github.com/evuraan/Swipe) - Swipe/2.x - Native compatibility with `X11` and `Wayland` 
  
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
# swipe config file.
# lines starting with # are ignored. 

# 3 Button Gestures:
3up:     "KEY_LEFTALT + KEY_LEFT"
3down:   "KEY_LEFTALT + KEY_RIGHT"
3left:   "KEY_HOME"
3right:  "KEY_END"

# 4 Button Gestures:
4left:  "KEY_LEFTALT + KEY_LEFT"
4right: "KEY_LEFTALT + KEY_RIGHT"
4up:    "KEY_HOME"
4down:  "KEY_END"

```
## Debug option
Run with `-d` option to have debug info onto the terminal:
```bash$ ./swipe -c swipe.conf -d 
$ ./swipe -d
Copyright © 2021 Evuraan <evuraan@gmail.com>. All rights reserved.
This program comes with ABSOLUTELY NO WARRANTY.
Sun Aug 15 14:10:29 2021 Swipe/1.06d Howdy!
Sun Aug 15 14:10:29 2021 Swipe/1.06d keyboard device: /dev/input/event3
Sun Aug 15 14:10:29 2021 [C] [getFd] fd opened: 3
Sun Aug 15 14:10:29 2021 Swipe/1.06d pid: 92985
Sun Aug 15 14:10:41 2021 Swipe/1.06d left: 0 right: 2 up: 0 down: 0
Sun Aug 15 14:10:41 2021 Swipe/1.06d movedTo: RIGHT
Sun Aug 15 14:10:41 2021 [C] [emit] emitted 24 bytes type 1 code 56
Sun Aug 15 14:10:41 2021 [C] [emit] emitted 24 bytes type 1 code 106
Sun Aug 15 14:10:41 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sun Aug 15 14:10:41 2021 [C] [emit] emitted 24 bytes type 1 code 56
Sun Aug 15 14:10:41 2021 [C] [emit] emitted 24 bytes type 1 code 106
Sun Aug 15 14:10:41 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sun Aug 15 14:10:41 2021 [C] [handleEvents] Handled 2 events
Sun Aug 15 14:10:41 2021 Swipe/1.06d Gesture type 3, intent: RIGHT, cmd: KEY_LEFTALT + KEY_RIGHT

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
