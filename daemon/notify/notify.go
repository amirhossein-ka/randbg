package notify

import (
	"log"
	"os/exec"
	"path/filepath"
)

func Notify(title, message, iconPath string) error {
	iconPath = absPath(iconPath)
	cmd := func() error {
        command, err := exec.LookPath("notify-send")
        if err != nil {
            return err
        }

        c := exec.Command(command, title, message, "-i", iconPath)
        return c.Run()
	}


    err := cmd()
    if err != nil {
        log.Printf("cant send notification, err => %v\n", err)
        return err
    }

    return nil
}

func absPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		p = path
	}

	return p
}
