package daemon

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	LockFile = "/tmp/randbg.lck"
)

// locked check that a lockfile exists, return true if exists
func (d *Daemon) locked() bool {
	if _, err := os.Stat(LockFile); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// lock creates a lockfile at LockFile
func (d *Daemon) lock() {
	f, err := os.OpenFile(LockFile, os.O_CREATE|os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("cannot create lockfile at: %s\n", err)
	}
}

// unlock removes lockfile from LockFile
func (d *Daemon) unlock() {
	err := os.Remove(LockFile)
	if err != nil {
		panic(fmt.Sprintf("cannot remove lockfile, err: %s\n", err))
	}
}
