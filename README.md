# Swipe
Gestures on Linux

![Swipe](./images/Swipe_300x300.png)

## Features
- Easy Installation - download (or git clone) and run. 
- No dependency on Python or Ruby
- Supports Config files
- Wide range of devices supported. 
## Requirements 
- libinput-tools and xdotools 
  ```bash 
    sudo apt-get install libinput-tools xdotool
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
## Setup
- Download `Swipe` to a folder. (Either download the [release](https://github.com/evuraan/Swipe/releases/download/1.06d/swipe), or clone this repo, or download and extract the <a href="https://github.com/evuraan/Swipe/archive/refs/heads/main.zip">Zip file</a>.)
- Launch `swipe` with an optional config file 

See your distro's documentation to setup `Swipe` as a [`Startup Application`](./images/Startup.png) - an [application](./images/Startup.png) that starts when a desktop user logs in. 

## Usage:
If no config file is specified, `Swipe` would use a default configuration.

```bash
$ ./swipe -h
Usage: ./swipe
  -h  --help         print this usage and exit
  -v  --version      print version information and exit
  -s  --sampleCfg    show sample config
  -d  --debug        show verbose output
  -c  /etc/ku.conf   config file to use 
```
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
Desktop notifications ([example](./images/Debug.png)) are also enabled in debug mode - which shows the details of the even intercepted. 
## Checksums:
```bash
$ ./swipe -v
Version: Swipe/1.06d
$ sha512sum swipe 
2a0047e2c3682243aec564c61169353d3fe6e2a2ff2796acb07572b10bc43a7931a8373d74e2069ee7b63b766cf707db0af5ceca7b59baa38be778c55b8371a0  swipe
```
