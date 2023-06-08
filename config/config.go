package config

import (
	"errors"
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	LoggingConfig  LoggingConfig  `json:"logging" yaml:"logging"`
	DataBaseConfig DataBaseConfig `json:"database" yaml:"database"`
}

type LoggingConfig struct {
	Development bool `json:"development" yaml:"development"`
}
type DataBaseConfig struct {
	UserName     string `json:"userName" yaml:"userName"`
	Password     string `json:"password" yaml:"password"`
	Host         string `json:"host" yaml:"host"`
	Schema       string `json:"schema" yaml:"schema"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
}

func NewConfig(configFile string) (*Config, error) {
	if configFile == "" {
		return nil, errors.New("you must provide configFile")
	}
	k := koanf.New(".")

	if err := k.Load(file.Provider(configFile), json.Parser()); err != nil {
		fmt.Println("Failed to load configration from json file ", configFile)
		return nil, err
	}
	var config Config
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		fmt.Printf("failed to unmarshal with conf. err: %v", err)
	}

	return &config, nil
}
