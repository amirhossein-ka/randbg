package daemon

import (
	"github.com/amirhossein-ka/randbg/config"
	"github.com/amirhossein-ka/randbg/daemon/directory"
	"github.com/amirhossein-ka/randbg/daemon/wallpaper"
	"log"
	"time"
)

type Daemon struct {
	cfg      *config.Config
	images   []string
	interval time.Duration
	timer    *time.Timer
	stop     chan struct{}
}

func New(cfg *config.Config) (*Daemon, error) {
	var err error
	stop := make(chan struct{})
	interval := time.Duration(cfg.DaemonConfig.Interval)
	timer := time.NewTimer(interval)
	d := &Daemon{
		cfg:      cfg,
		stop:     stop,
		timer:    timer,
		interval: interval,
	}
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Daemon) Run() {
	if d.locked() {
		log.Fatalln("daemon already running")
	}
	d.lock()
	defer d.unlock()

	go d.handleSignals()
	go d.setWall()

	<-d.stop
}

func (d *Daemon) setWall() {
	pics, err := directory.DirContent(d.cfg.DaemonConfig.ImgDirectory)
	if err != nil {
		log.Printf("an error occurred while loading image directory content: %v\n", err.Error())
		d.stop<- struct{}{}
	}
	d.images = pics
	for {
		select {
		case <-d.stop:
			return
		case <-d.timer.C:
			if err = wallpaper.ChangeWall(d.images); err != nil {
				log.Printf("an error occured while changing wallpaper: %v\n", err.Error())
				d.stop<- struct{}{}
			}
			d.timer.Reset(d.interval)
		}
	}
}
