package config

import "time"

const (
	DefaultInterval       = 15 * time.Minute
	DefaultBackgroundPath = "$HOME/Pictures/Wallpapers/"
)

func setDefaults(cfg *Config) {
	cfg.DaemonConfig.Interval = Duration(DefaultInterval)
	cfg.DaemonConfig.ImgDirectory = DefaultBackgroundPath
}
