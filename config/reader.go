package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var ErrInvalidFileExtension = errors.New("invalid config file extension. supported extensions are: .yml|.yaml and .json")

func Parse(path string, cfg *Config) error {
	switch filepath.Ext(path) {
	case ".json":
		return parseJson(path, cfg)
	case ".yml", ".yaml":
		return parseYaml(path, cfg)
	default:
		return ErrInvalidFileExtension
	}
}

func parseYaml(path string, cfg *Config) error {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			setDefaults(cfg)
			log.Println("config file does not exists, using default values. (flags will overwrite configs)")
			return nil
		}
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	return yaml.NewDecoder(file).Decode(cfg)
}

func parseJson(path string, cfg *Config) error {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			setDefaults(cfg)
			log.Println("config file does not exists, using default values. (flags will overwrite configs)")
			return nil
		}
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	return json.NewDecoder(file).Decode(cfg)
}
