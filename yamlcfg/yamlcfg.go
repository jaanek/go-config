package yamlcfg

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func NewConfig(configPath string, config interface{}) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return err
	}
	return nil
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
