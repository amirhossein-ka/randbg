package config

type (
	Config struct {
		DaemonConfig DaemonConfig `json:"daemon_config" yaml:"daemon_config"`
	}

	DaemonConfig struct {
		ImgDirectory string   `json:"image_directory,omitempty" yaml:"img_directory"`
        Interval     Duration `json:"interval,omitempty" yaml:"interval"`
    }
)
