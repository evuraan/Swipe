package main

import (
	"embed"
	"fmt"
	"os"
	"time"
)

//go:embed images/Swipe_300x300.png indicator/panel.py
var embedFs embed.FS

// writes to disk and returns the path. type 0: icon, 1: py file
func writeFile(typeInt int) (string, error) {

	retrieve := "images/Swipe_300x300.png"
	if typeInt == 1 {
		retrieve = "indicator/panel.py"
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
