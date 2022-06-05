package main

import (
	"fmt"
	"github.com/amirhossein-ka/randbg/config"
	"log"
	"os"
	"path/filepath"
)

var cfg config.DaemonConfig

func parseAll(cfg *config.DaemonConfig) error {
	ParseFlags()
	if configPath != "" {
		if err := config.Parse(configPath, cfg); err != nil {
			return err
		}
		setFlags(cfg)
	} else {
		configPath, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		if err := config.Parse(filepath.Join(configPath, "randbg", "randbg.yml"), cfg); err != nil {
			return err
		}
		setFlags(cfg)
	}
	return nil
}

func setFlags(cfg *config.DaemonConfig) {
	if interval != 0 && interval != defaultInterval {
		cfg.Interval = config.Duration(interval)
	}
	if backgroundPath != "" && backgroundPath != defaultWallPath {
		cfg.ImgDirectory = backgroundPath
	}
}

func init() {
	if err := parseAll(&cfg); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println("Hello World !")
	fmt.Printf("%#v\n", cfg)
}
