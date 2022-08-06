package config

import "time"

const (
	DefaultInterval       = 15 * time.Minute
	DefaultBackgroundPath = "$HOME/Pictures/Wallpapers/"
	DefaultPidFileAddr    = "/var/run/randbg.pid"
	DefaultSocketFile     = "/tmp/randbg.socket"
)

func setDefaults(cfg *Config) {
	cfg.DaemonConfig.Interval = Duration(DefaultInterval)
	cfg.DaemonConfig.ImgDirectory = DefaultBackgroundPath
	cfg.DaemonConfig.PidFile = DefaultPidFileAddr
	cfg.DaemonConfig.SocketFile = DefaultSocketFile

	cfg.CmdConfig.PidFile = DefaultPidFileAddr
	cfg.CmdConfig.SocketFile = DefaultSocketFile
}
