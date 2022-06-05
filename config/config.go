package config

type (
	DaemonConfig struct {
		ImgDirectory string   `json:"img_directory,omitempty" yaml:"img_directory"`
		Interval     Duration `json:"interval,omitempty" yaml:"interval"`
	}
)
