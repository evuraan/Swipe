# Swipe

Gestures on Linux.

![Swipe](./images/Swipe_300x300.png)

<p>https://evuraan.info/Swipe/ 
 
## Screengrab:
-  [https://evuraan.info/evuraan/stuff/Swipe.mp4](https://evuraan.info/evuraan/stuff/Swipe.mp4) 
 
## Features
Swipe uses a novel yet simple correlation mechanism to determine directional intent from event coordinates.
- Wide range of devices supported 
- Touchscreens - single, double, triple, quad touches supported. 
- Touchpad - double, triple, quad touches 
- Wayland and X11 compatible
- Easy Installation - download (or git clone) and run. 
- No dependency on Python or Ruby
- Supports Config files
- Supports 480+ input events.

 
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
- Download `Swipe` to a folder. (Either download the latest build from the [bin folder](./bin/), or clone this repo, or download and extract the <a href="https://github.com/evuraan/Swipe/archive/refs/heads/main.zip">Zip file</a>.)
	<pre>
	$ wget https://github.com/evuraan/Swipe/blob/main/bin/swipe?raw=true -O swipe 
	$ chmod 775 swipe </pre>
- Launch `swipe`. In most cases `Swipe` will look for and find everything it needs, otherwise you can use the options and/or a config file.

See your distro's documentation to setup `Swipe` as a [`Startup Application`](./images/Startup.png) - an [application](./images/Startup.png) that starts when a desktop user logs in.

<img src="./images/Startup.png"  width='300'>

## Usage:

```bash
 $ ./swipe -h
Usage of ./swipe:
  -available
        Show available devices
  -c string
        Config file path
  -debug
        Enable debug
  -delay duration
        Delay between events (default 100ms)
  -help
        Show help
  -i string
        Input device, eg: /dev/input/event3
         (default "/dev/input/event3")
  -keys
        Show available keys
  -noIndicator
        Disable status icon
  -sampleCfg
        Show sample config
  -version
        Show version
```

- If no config file is specified, `Swipe` would use a default configuration.
- If an appropriate `kbd` device cannot be found, `Swipe` will ask you to specify a suitable device using the `-i` option.

## Config

Generate a sample config file with -sampleCfg option. [Here are some other config examples.](https://github.com/evuraan/Swipe/issues/7)

```bash
$ ./swipe -s

Sample Config:
# 2 button touchpad gestures
2right:     "KEY_LEFTALT + KEY_LEFT"
2left:      "KEY_LEFTALT + KEY_RIGHT"

# 3 button touchpad gestures
3right:     "KEY_LEFTALT + KEY_LEFT"
3left:      "KEY_LEFTALT + KEY_RIGHT"
3up:        "KEY_SPACE"
3down:      "KEY_LEFTSHIFT + KEY_SPACE"

# 4 button touchpad gestures
4right:     "KEY_MUTE"
4left:      "KEY_MUTE"
4up:        "KEY_VOLUMEUP"
4down:      "KEY_VOLUMEDOWN"

# Touchscreen gestures
touch1up:   "KEY_UP"
touch1down: "KEY_DOWN"
touch1left: "KEY_LEFTALT + KEY_LEFT"
touch1right:"KEY_LEFTALT + KEY_RIGHT"

touch2up:   "KEY_UP"
touch2down: "KEY_DOWN"
touch2left: "KEY_LEFTALT + KEY_LEFT"
touch2right:"KEY_LEFTALT + KEY_RIGHT"

touch3up:   "KEY_UP"
touch3down: "KEY_DOWN"
touch3left: "KEY_LEFTALT + KEY_LEFT"
touch3right:"KEY_LEFTALT + KEY_RIGHT"

touch4up:   "KEY_UP"
touch4down: "KEY_DOWN"
touch4left: "KEY_LEFTALT + KEY_LEFT"
touch4right:"KEY_LEFTALT + KEY_RIGHT"
```

Create and edit a custom config to suite your likings:

```bash
$ ./swipe -sampleCfg > mySwipe.conf
```

Make edits to `mySwipe.conf` and launch swipe as `$ ./swipe -c mySwipe.conf`

### Config file example

This is config the author currently uses:

<pre>
# 2 Button Touchpad 
2right: "KEY_LEFTALT + KEY_LEFT"
2left:  "KEY_LEFTALT + KEY_RIGHT"

# 3 Button Touchpad Gestures:
# Zoom in and out
3right: "KEY_LEFTCTRL + KEY_RIGHTSHIFT + KEY_EQUAL" 
3left:  "KEY_RIGHTCTRL + KEY_0"
3up:    "KEY_LEFTCTRL + KEY_RIGHTSHIFT + KEY_EQUAL" 
3down:  "KEY_RIGHTCTRL + KEY_MINUS"

# 4 Button Touchpad Gestures:
# Vol Up/Down/Mute
4right: "KEY_MUTE"
4left:  "KEY_LEFTALT + KEY_RIGHTCTRL + KEY_P"
4up:    "KEY_VOLUMEUP"
4down:  "KEY_VOLUMEDOWN"

# 5 - Touchscreens
touch1up:  "KEY_SPACE"
touch1down: "KEY_RIGHTSHIFT + KEY_SPACE"
touch1left:  "KEY_LEFTALT + KEY_RIGHT"
touch1right:  "KEY_LEFTALT + KEY_LEFT"

touch2up:  "KEY_RIGHTCTRL + KEY_KPPLUS"
touch2down: "KEY_RIGHTCTRL + KEY_MINUS"
touch2left: "KEY_LEFTALT + KEY_RIGHT"
touch2right: "KEY_LEFTALT + KEY_LEFT"
</pre>

## Debug option

Run with `-debug` option to have debug info onto the terminal:

```bash
$ ./swipe -debug -c /tmp/swipe.conf
Fri Sep 03 19:07:12 2021 Swipe/3.01e Read 14 values from the config file
Fri Sep 03 19:07:12 2021 Swipe/3.01e 2 key touchpad events: map[LEFT:KEY_LEFTALT + KEY_RIGHT RIGHT:KEY_RIGHTALT + KEY_LEFT]
Fri Sep 03 19:07:12 2021 Swipe/3.01e 3 key touchpad events: map[DOWN:KEY_LEFTSHIFT + KEY_SPACE LEFT:KEY_LEFTALT + KEY_RIGHT RIGHT:KEY_LEFTALT + KEY_LEFT UP:KEY_SPACE]
Fri Sep 03 19:07:12 2021 Swipe/3.01e 4 key touchpad events: map[DOWN:KEY_END LEFT:KEY_LEFTALT + KEY_RIGHT RIGHT:KEY_LEFTALT + KEY_LEFT UP:KEY_HOME]
Fri Sep 03 19:07:12 2021 Swipe/3.01e touchscreen events: map[DOWN:KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP + KEY_UP LEFT:KEY_LEFTALT + KEY_RIGHT RIGHT:KEY_LEFTALT + KEY_LEFT UP:KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN + KEY_DOWN]
```

Desktop notifications ([example](./images/Debug.png)) are also enabled in debug mode - which shows the details of the event intercepted.

## Keys and buttons supported:

Swipe supports about `482` keys/buttons - pretty much inline with Linux's [input-event-codes.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h). <br>  
Run `swipe -keys` to see a full list:

```bash
$ ./swipe -keys
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
