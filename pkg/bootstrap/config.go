package bootstrap

import (
	"io/ioutil"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type Config interface {
	GetNetwork() string
	GetAddress() string
	IsLoggerEnabled() bool
	GetSecret() string
	GetSecretWhitelist() []string
}

var _ Config = &ConfigStruct{}

type ConfigStruct struct {
	Network         string   `yaml:"network"`
	Address         string   `yaml:"address"`
	LoggerEnabled   bool     `yaml:"logger-enabled"`
	Secret          string   `yaml:"secret"`
	SecretWhitelist []string `yaml:"secret-whitelist"`
}

func (cfg *ConfigStruct) GetNetwork() string {
	return cfg.Network
}

func (cfg *ConfigStruct) GetAddress() string {
	return cfg.Address
}

func (cfg *ConfigStruct) IsLoggerEnabled() bool {
	return cfg.LoggerEnabled
}

func (cfg *ConfigStruct) GetSecret() string {
	return cfg.Secret
}

func (cfg *ConfigStruct) GetSecretWhitelist() []string {
	return cfg.SecretWhitelist
}

// DecodeConfig loads config from yaml bytes
func DecodeConfig(data []byte, cfg interface{}) error {
	var tmp map[string]interface{}
	err := yaml.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "yaml",
		Result:      cfg,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(tmp)
}

// DecodeConfigFile loads config from yaml file
func DecodeConfigFile(path string, cfg interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return DecodeConfig(data, cfg)
}
