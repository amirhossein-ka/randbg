package main

import (
	"os"
	"strconv"

	"github.com/amirhossein-ka/randbg/config"
)

func createPid(cfg *config.DaemonConfig) error {
	pidFile, err := os.CreateTemp("", cfg.PidFile)
	if err != nil {
		return err
	}
	defer pidFile.Close()

	appPid := os.Getpid()

  _, err = pidFile.WriteString(strconv.Itoa(appPid))
  if err != nil {
    return err
  }
  
	return nil
}
