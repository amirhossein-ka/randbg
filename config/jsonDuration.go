package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"time"
)

type Duration time.Duration

func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(*d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		var err error
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (d *Duration) UnmarshalYAML(value *yaml.Node) error {

	dr, err := time.ParseDuration(value.Value)
	if err != nil {
		return err
	}
	*d = Duration(dr)
	return nil
}
