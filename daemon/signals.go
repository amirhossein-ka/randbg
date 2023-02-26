package daemon

import (
	"fmt"
	"github.com/amirhossein-ka/randbg/config"
	"github.com/amirhossein-ka/randbg/daemon/directory"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (d *Daemon) handleSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGINT)
	for {
		select {
		case <-d.stop:
			return
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				d.sigint()
			}
		}
	}
}

func (d *Daemon) sigint() {
	log.Println("received SIGINT/SIGTERM, exiting gracefully...")
	d.stop <- struct{}{}
}

func (d *Daemon) sighup() {
	log.Println("received SIGHUP, reloading config and image directory content...")
	// Todo add a args struct for saving this stuff to avoid hard coding
	if err := config.Parse("", d.cfg); err != nil {
		panic(fmt.Sprintf("an error occurred while reloading configuration: %v\n", err.Error()))
	}

	if _, err := directory.DirContent(d.cfg.DaemonConfig.ImgDirectory); err != nil {
		panic(fmt.Sprintf("an error occurred while reloading image directory content: %v\n", err.Error()))
	}
}
