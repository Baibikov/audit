package config

import (
	"os"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Grpc 		Grpc 		`yaml:"grpc"`
	Clickhouse 	Clickhouse 	`yaml:"clickhouse"`
}

type Grpc struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

type Clickhouse struct {
	Addr []string `yaml:"addr"`

	Auth struct {
		DB 			string `yaml:"db"`
		User 		string `yaml:"user"`
		Password 	string `yaml:"password"`
	} `yaml:"auth"`

	Settings struct {
		MaxExecutionTime int `yaml:"maxExecutionTime"`
	}

	DialSecondsTimeout int `yaml:"dialSecondsTimeout"`
	Debug bool `yaml:"debug"`
}

func New(path string) (config Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return config, errors.Wrapf(err, "open config by path %s", path)
	}
	defer func(err error) {
		multierr.AppendInto(&err, file.Close())
	}(err)

	d := yaml.NewDecoder(file)

	err = d.Decode(&config)
	if err != nil {
		return config, errors.Wrap(err, "decode config information")
	}

	return config, nil
}