package config

import "testing"

const TestData = "./testdata/"

func TestParse(t *testing.T) {
	type args struct {
		Path string
		Cfg  *Config
	}

	tests := []struct {
		Name    string
		Args    args
		WantErr bool
	}{
		{
			Name: "json file",
			Args: args{
				Path: TestData + "test_cfg.json",
				Cfg:  &Config{},
			},
			WantErr: false,
		},
		{
			Name: "yaml file",
			Args: args{
				Path: TestData + "test_cfg.yaml",
				Cfg:  &Config{},
			},
			WantErr: false,
		},
		{
			Name: "yml file",
			Args: args{
				Path: TestData + "test_cfg.yml",
				Cfg:  &Config{},
			},
			WantErr: false,
		},
		{
			Name: "Unknown file extension",
			Args: args{
				Path: TestData + "test_cfg.xxx",
				Cfg:  &Config{},
			},
			WantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(tt *testing.T) {
			err := Parse(test.Args.Path, test.Args.Cfg)
			if (err != nil) != test.WantErr {
				tt.Errorf("Parse(): wantErr: %v, got: %v\n", test.WantErr, err)
			}
		})
	}
}
