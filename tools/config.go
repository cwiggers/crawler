package tools

import (
	"encoding/json"
	"os"
)

type Config struct {
	Addr string `json:"addr"`
}

var (
	Conf *Config
)

func LoadConf(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(r)
	Conf = &Config{}
	err = decoder.Decode(Conf)
	if err != nil {
		return err
	}
	return nil
}
