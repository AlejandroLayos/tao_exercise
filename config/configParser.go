package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

//Data object to manage the .yaml file
type configurationSchema struct {
	Logs struct {
		Level      string `yaml:"level"`
		Path       string `yaml:"path"`
		MaxSize    string `yaml:"max_size"`
		MaxBackups string `yaml:"max_backups"`
		MaxAge     string `yaml:"max_age"`
	} `yaml:"logs"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Data struct {
		DataType string `yaml:"data_type"`
	} `yaml:"data"`
}

type config struct {
	cfg        *configurationSchema
	configPath string
}

type Config interface {
	ConfigYaml() error
}

func NewConfig() Config {
	return &config{&configurationSchema{}, "config/config.yaml"}
}

func (c config) ConfigYaml() error {
	f, err := os.Open(c.configPath)
	if err != nil {
		fmt.Print("Error retrieving configuration file")
		return err
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c.cfg)
	if err != nil {
		fmt.Print("Error decoding configuration")
		return err
	}
	err = f.Close()

	if err != nil {
		fmt.Print("Error Closing configuration file")
		return err
	}

	err = os.Setenv("LOGLEVEL", c.cfg.Logs.Level)
	os.Setenv("LOG_PATH", c.cfg.Logs.Path)
	os.Setenv("MAX_SIZE", c.cfg.Logs.MaxSize)
	os.Setenv("MAX_BACKUPS", c.cfg.Logs.MaxBackups)
	os.Setenv("MAX_AGE", c.cfg.Logs.MaxAge)
	os.Setenv("DATATYPE", c.cfg.Data.DataType)
	os.Setenv("REST_PORT", c.cfg.Server.Port)

	return nil
}
