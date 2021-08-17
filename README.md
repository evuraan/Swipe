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
- Enable touchpad events
  ```bash
    gsettings set org.gnome.desktop.peripherals.touchpad send-events enabled
     ```
- Tool: Depending on whether you use `X11` or `Wayland`, you will need `xdotool` or `ydotool` etc. 
See [Config](#config-file) notes for more. 
  
## Setup
- Download `Swipe` to a folder. (Either download the latest [release](https://github.com/evuraan/Swipe/releases/download/1.06d/swipe), or clone this repo, or download and extract the <a href="https://github.com/evuraan/Swipe/archive/refs/heads/main.zip">Zip file</a>.)
- Launch `swipe` with an optional config file 

See your distro's documentation to setup `Swipe` as a [`Startup Application`](./images/Startup.png) - an [application](./images/Startup.png) that starts when a desktop user logs in. 

## Usage:

```bash
$ ./swipe -h
Usage: ./swipe
  -h  --help         print this usage and exit
  -v  --version      print version information and exit
  -s  --sampleCfg    show sample config
  -d  --debug        show verbose output
  -c  /etc/ku.conf   config file to use 
```
  
## Config file
If no config file is specified, `Swipe` would use a default configuration assuming `xdotool` usage. You can use another (like `ydotool` or `xte`), set your config file as appropriate - `Swipe` needs to know what command to run for each type of intercepted event. 

Generate a sample config file with `-s` option:
```bash
$ ./swipe -s

Sample Config: 
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
```
## Debug option
Run with `-d` option to have debug info onto the terminal:
```bash$ ./swipe -c swipe.conf -d 
Sun Aug 01 20:27:25 2021 Swipe/1.06d Howdy!
Sun Aug 01 20:27:25 2021 Swipe/1.06d pid: 2167
Sun Aug 01 20:27:27 2021 Swipe/1.06d left: 0 right: 7 up: 0 down: 0
Sun Aug 01 20:27:27 2021 Swipe/1.06d movedTo: RIGHT

```
Desktop notifications ([example](./images/Debug.png)) are also enabled in debug mode - which shows the details of the event intercepted. 
## Optional: Build 
If you prefer to build yourself, you will need the [Go Programming Language](https://golang.org/dl/) installed on your System. 

Go into the folder and build as: 
``` 
go build
```

