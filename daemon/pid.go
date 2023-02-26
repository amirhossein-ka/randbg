package daemon

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)


func getPidFileName() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(home + "/.randbg.pid"), nil
}

// CreatePid create a pid file in PidFile path
func CreatePid(pid int) error {
    pidFile, err := getPidFileName()
    if err != nil {
        return err
    }
	p, err := os.OpenFile(pidFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
    defer p.Close()

	_, err = p.WriteString(fmt.Sprintf("%d", pid))

	return err

}

func ReadPid() (int, error) {
    pidFile, err := getPidFileName()
    if err != nil {
        return 0, err
    }
	// pid string
	var pString string
	p, err := os.OpenFile(pidFile, os.O_RDONLY, 0600)
	if err != nil {
		return 0, err
	}
    defer p.Close()

	scanner := bufio.NewScanner(p)
	for scanner.Scan() {
		pString = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(pString)
	if err != nil {
		return 0, err
	}

	return pid, nil
}

