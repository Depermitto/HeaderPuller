package pkg

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

const ConfigFilepath = "hp.yaml"

type ConfigPkgs struct {
	Packages []ConfigPkg `yaml:"packages"`
}

type ConfigPkg struct {
	Id     int      `yaml:"id"`
	Name   string   `yaml:"name"`
	Link   string   `yaml:"link"`
	Remote string   `yaml:"remote"`
	Local  []string `yaml:"local"`
}

func (pkgs ConfigPkgs) Update() {
	for i := range pkgs.Packages {
		pkgs.Packages[i].Id = i + 1
	}
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
	return pkgs, err
}

func Marshall(pkgs ConfigPkgs) error {
	marshal, err := yaml.Marshal(pkgs)
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFilepath, marshal, 0755)
	return err
}
