package yamlcfg

import (
	"flag"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// based on https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
func NewConfig(configPath string, config interface{}) error {
	err := readFile(configPath, config)
	if err != nil {
		return err
	}
	return readEnv(configPath, config)
}

func readFile(configPath string, config interface{}) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(config); err != nil {
		return err
	}
	return nil
}

func readEnv(configPath string, config interface{}) error {
	return envconfig.Process("", config)
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}
