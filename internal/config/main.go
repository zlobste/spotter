package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config defines an interface of global service configurations.
type Config interface {
	Logger() *logrus.Logger
}

// Config defines global service configurations.
type config struct {
	Log string `yaml:"log"`
}

// NewConfig returns global service configurations.
func New(path string) Config {
	cfg := config{}

	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("failed to read config: %s", path)))
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("failed to unmarshal config: %s", path)))
	}

	return &cfg
}

// Logger returns new configured logger.
func (c *config) Logger() *logrus.Logger {
	level, err := logrus.ParseLevel(c.Log)
	if err != nil {
		panic(errors.Wrapf(err, "failed to parse logging level %s", c.Log))
	}

	logger := logrus.New()
	logger.SetLevel(level)

	return logger
}
