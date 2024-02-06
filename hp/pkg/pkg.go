package pkg

import (
	"HeaderPuller/hp"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

const ConfigFilepath = "hp.yaml"

type ConfigPkgs struct {
	Packages []ConfigPkg `yaml:"packages"`
}

type ConfigPkg struct {
	Name   string   `yaml:"name"`
	Link   string   `yaml:"link"`
	Remote string   `yaml:"remote"`
	Local  []string `yaml:"local"`
}

func Unmarshalled() (pkgs ConfigPkgs, err error) {
	buffer, err := os.ReadFile(ConfigFilepath)
	if errors.Is(err, os.ErrNotExist) {
		os.Create(ConfigFilepath)
		buffer, err = os.ReadFile(ConfigFilepath)
	}
	if err != nil {
		return ConfigPkgs{}, err
	}

	if err = yaml.Unmarshal(buffer, &pkgs); err != nil {
		return ConfigPkgs{}, err
	}
	return pkgs, nil
}

func Marshall(pkgs ConfigPkgs) error {
	marshal, err := yaml.Marshal(pkgs)
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigFilepath, marshal, hp.Perm)
}
