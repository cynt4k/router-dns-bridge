package cmd

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func readDefaultConfig(configDir string) error {
	yamlFile, err := ioutil.ReadFile(filepath.Join(configDir, "default.yaml"))

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, cfg)

	if err != nil {
		return err
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(yamlFile))

	if err != nil {
		return err
	}

	return nil
}
