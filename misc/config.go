package misc

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Addr  string `json:"addr"`
	Redis struct {
		Addr           string        `json:"addr"`
		Passwd         string        `json:"passwd"`
		MaxActive      int           `json:"maxactive"`
		MaxIdle        int           `json:"maxidle"`
		IdleTimeout    time.Duration `json:"idletimeout"`
		ConnectTimeout time.Duration `json:"connect_imeout"`
		ReadTimeout    time.Duration `json:"read_timeout"`
		WriteTimeout   time.Duration `json:"write_timeout"`
		Expire         int           `json:"expire"`
	} `json:"redis"`
}

var (
	Conf *Config
)

func LoadConf(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	decoder := json.NewDecoder(r)
	Conf = &Config{}
	err = decoder.Decode(Conf)
	if err != nil {
		return err
	}
	return nil
}
