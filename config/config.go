package config

import "time"

type (
	Config struct {
		ImgDirectory string        `json:"img_directory,omitempty" yaml:"img_directory"`
		Interval     time.Duration `json:"interval,omitempty" yaml:"interval"`
	}
)
