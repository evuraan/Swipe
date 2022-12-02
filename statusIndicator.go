package main

import (
	"embed"
	"fmt"
	"os"
	"time"
)

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
