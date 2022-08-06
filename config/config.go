package config

type (
	Config struct {
		DaemonConfig DaemonConfig `json:"daemon_config" yaml:"daemon_config"`
		CmdConfig    CmdConfig    `yaml:"cmd_config" json:"cmd_config"`
	}

	DaemonConfig struct {
		ImgDirectory string   `json:"img_directory,omitempty" yaml:"img_directory"`
		Interval     Duration `json:"interval,omitempty" yaml:"interval"`
		PidFile      string   `json:"pid_file" yaml:"pid_file"`
		SocketFile   string   `json:"socket_file" yaml:"socket_file"`
	}
	CmdConfig struct {
		PidFile    string `json:"pid_file" yaml:"pid_file"`
		SocketFile string `json:"socket_file" yaml:"socket_file"`
	}
)
