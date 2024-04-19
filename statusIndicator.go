package main

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"golang.org/x/sys/unix"
)

type conduitStruct struct {
	sync.RWMutex
	fifoPath   string
	isDisabled bool
	filePtr    *os.File
}

//go:embed images/Swipe_300x300.png indicator/panel.py images/Swipe.png
var embedFs embed.FS

// writes to disk and returns the path. type 0: icon, 1: py file 2: Swipe.png (icon change indicator)
func writeFile(typeInt int) (string, error) {

	retrieve := "images/Swipe_300x300.png"

	switch typeInt {
	case 1:
		retrieve = "indicator/panel.py"
	case 2:
		retrieve = "images/Swipe.png"
	}

	x, err := embedFs.ReadFile(retrieve)
	if err != nil {
		return "", err
	}

	writeTo := fmt.Sprintf("%s/Swipe%d.%d", os.TempDir(), time.Now().UnixNano(), typeInt)
	if err = os.WriteFile(writeTo, x, 0400); err != nil {
		return "", err
	}

	return writeTo, nil
}

func setupPanelConduit() {
	if statusIconDisabled {
		conduit.isDisabled = true
		return
	}

	fifoPath := fmt.Sprintf("%s/swipe%d", os.TempDir(), time.Now().UnixNano())
	if err := unix.Mkfifo(fifoPath, 0600); err != nil {
		fmt.Fprintf(os.Stderr, "mkfifo err: %s", err)
		return
	}
	if file, err := os.OpenFile(filepath.FromSlash(fifoPath), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600); err == nil {
		conduit.filePtr = file
		conduit.isDisabled = statusIconDisabled
		conduit.fifoPath = fifoPath
	}

	var pyFile, ico, icoChange string
	var err error
	if icoChange, err = writeFile(2); err != nil {
		return
	}
	if pyFile, err = writeFile(1); err != nil {
		return
	}
	if ico, err = writeFile(0); err != nil {
		return
	}

	// Which python3 binary to use?
	// We must use /usr/bin/python3 because the gi repository is not available via pip.
	// The gi repository, part of PyGObject (Python GObject Introspection), is not
	// directly available via pip because it requires certain libraries to be
	// installed at the system level. PyGObject is used to interface with libraries
	// written in C, such as GTK, GLib, and many others, through GObject Introspection.

	//#nosec G204
	cmd := exec.Command("/usr/bin/python3", pyFile, ico, fmt.Sprintf("%d", os.Getpid()), icoChange, conduit.fifoPath)
	// Create buffers to capture standard output and standard error
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	//fmt.Println(cmd)
	//return

	err = cmd.Run() // blocks
	_ = conduit.filePtr.Close()
	removeThese := []string{conduit.fifoPath, ico, pyFile, icoChange}
	for _, j := range removeThese {
		_ = os.Remove(j)
	}

	switch err {
	case nil:
		// user wants to exit by left click.
		os.Exit(0)
	default:
		// err is not nil
		if err.Error() == "signal: killed" {
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "panel.py err: %s stderr: %s stdout: %s\n", err, cmd.Stderr, cmd.Stdout)
		}
	}

}

func (c *conduitStruct) notifyFifo() bool {
	c.RLock()
	defer c.RUnlock()

	if c.isDisabled {
		return false
	}

	if _, err := io.WriteString(c.filePtr, "evt\n"); err == nil {
		return true
	} else {
		fmt.Fprintf(os.Stderr, "notifyFifo err: %s\n", err)
	}

	go func() {
		fmt.Fprintf(os.Stderr, "Error 3.1 - disabling further notifications\n")
		c.Lock()
		defer c.Unlock()
		c.isDisabled = true
	}()

	return false
}
