package main

import (
	"flag"
	"github.com/amirhossein-ka/randbg/config"

	// "log"
	// "os"
	// "path/filepath"
	"time"
)

const (
	defaultWallPath string = "$HOME/Pictures"
	defaultInterval        = 15 * time.Minute
)

var (
	configPath     string
	backgroundPath string
	interval       time.Duration
)

func ParseFlags() {
	// short
	flag.StringVar(&configPath, "c", "", "path to config file, default to USER_CONF_DIR/randbg/config.yml (shorthand)")
	flag.StringVar(&backgroundPath, "p", config.DefaultBackgroundPath, "path to wallpaper directory (shorthand)")
	flag.DurationVar(&interval, "i", config.DefaultInterval, "interval of changing of wallpapers in minutes.")

	// long
	flag.StringVar(&configPath, "config", "", "path to config file, default to USER_CONF_DIR/randbg/config.yml")
	flag.StringVar(&backgroundPath, "path", config.DefaultBackgroundPath, "path to wallpaper directory ")
	flag.DurationVar(&interval, "interval", config.DefaultInterval, "interval of changing of wallpapers in minutes.")
	flag.Parse()

}
