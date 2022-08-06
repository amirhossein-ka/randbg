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
	cfg config.Config
)

func parseAll(cfg *config.Config) error {
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

func setFlags(cfg *config.Config) {
	if interval != 0 && interval != config.DefaultInterval {
		cfg.DaemonConfig.Interval = config.Duration(interval)
	}
	if backgroundPath != "" && backgroundPath != config.DefaultBackgroundPath {
		cfg.DaemonConfig.ImgDirectory = backgroundPath
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

	pics, err := lib.DirContent(cfg.DaemonConfig.ImgDirectory)
	if err != nil {
		log.Fatalln(err, cfg.DaemonConfig.ImgDirectory)
	}

	if err = lib.ChangeWall(ctx, pics); err != nil {
		panic(err)
	}
	// _ = pics

	for {
		select {
		case <-time.Tick(time.Duration(cfg.DaemonConfig.Interval)):
			if err := lib.ChangeWall(ctx, pics); err != nil {
				panic(err)
			}
		case <-ctx.Done():
			log.Println("context done ??? stopping...")
			cancel()
		}
	}
}
