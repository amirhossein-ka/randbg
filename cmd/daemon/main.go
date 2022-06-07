package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/amirhossein-ka/randbg/config"
	"github.com/amirhossein-ka/randbg/lib"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	cfg config.DaemonConfig
)

func parseAll(cfg *config.DaemonConfig) error {
	if !flag.Parsed() {
		ParseFlags()
	}
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
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	fmt.Printf("%#v\n", cfg)
	go handleSignals()

	pics, err := lib.DirContent(cfg.ImgDirectory)
	if err != nil {
		log.Fatalln(err, cfg.ImgDirectory)
	}

	if err = lib.ChangeWall(ctx, pics); err != nil {
		panic(err)
	}
	for {
		select {
		case <-time.Tick(time.Duration(cfg.Interval)):
			if err := lib.ChangeWall(ctx, pics); err != nil {
				panic(err)
			}
		case <-ctx.Done():
			log.Println("context done ??? stopping...")
			cancel()
		}
	}
}
