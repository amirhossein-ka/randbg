package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var ErrInvalidFileExtension = errors.New("invalid config file extension. supported extensions are: .yml|.yaml and .json")

func Parse(path string, cfg *DaemonConfig) error {
	switch filepath.Ext(path) {
	case ".json":
		return parseJson(path, cfg)
	case ".yml", ".yaml":
		return parseYaml(path, cfg)
	default:
		return ErrInvalidFileExtension
	}
}

func parseYaml(path string, cfg *DaemonConfig) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return yaml.NewDecoder(file).Decode(cfg)
}

func parseJson(path string, cfg *DaemonConfig) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(cfg)
}
