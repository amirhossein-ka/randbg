package args

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"time"
	"unsafe"

	"github.com/amirhossein-ka/randbg/config"
	"github.com/amirhossein-ka/randbg/daemon"
)

var ErrNotEnoughArgs = errors.New("not enough arguments provided")

type (
	Args struct {
		DaemonArgs Daemon
		daemon     bool
	}

	Daemon struct {
		ConfigPath string
		ImgPath    string
		Interval   time.Duration
		Verbose    bool
	}
)

func (a *Args) Init(args []string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	defConfigPath = path.Join(configDir, "randbg", "randbg.yml")

	d := flag.NewFlagSet("daemon", flag.ExitOnError)
	d.DurationVar(&a.DaemonArgs.Interval, "interval", time.Minute*10, "time interval of changing images")
	d.StringVar(&a.DaemonArgs.ImgPath, "image", "$HOME/Pictures/", "path of images")
	d.StringVar(&a.DaemonArgs.ConfigPath, "config", defConfigPath, "path to configuration")

	if len(args) < 2 {
		return ErrNotEnoughArgs
	}

	switch args[1] {
	case "daemon":
		if err = d.Parse(args[2:]); err != nil {
			return err
		}
		a.daemon = true
	default:
		fmt.Println("some help message")
		return nil
	}
	return nil
}

func (a *Args) RunApp(cfg *config.Config) error {
	if a.daemon {
		a.replaceArgs(cfg)
		dmn, err := daemon.New(cfg)
        dmn.CfgPath = a.DaemonArgs.ConfigPath
		if err != nil {
			return err
		}
        fmt.Println(unsafe.Sizeof(dmn))
        fmt.Println(unsafe.Alignof(dmn))
		dmn.Run()
	}

	return nil
}

func (a *Args) replaceArgs(cfg *config.Config) {
	if a.DaemonArgs.Interval != 0 && a.DaemonArgs.Interval != defInterval {
		cfg.DaemonConfig.Interval = config.Duration(a.DaemonArgs.Interval)

	}
	if a.DaemonArgs.ImgPath != "" && a.DaemonArgs.ImgPath != defImgPath {
		cfg.DaemonConfig.ImgDirectory = a.DaemonArgs.ImgPath
	}
}
