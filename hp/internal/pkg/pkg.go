package pkg

import (
	"HeaderPuller/hp"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"slices"
)

const ConfigFilepath = "hp.yaml"

type Pkgs struct {
	Packages []Pkg `yaml:"packages"`
}

func (pkgs *Pkgs) Contains(link string, remote string) bool {
	return slices.ContainsFunc(pkgs.Packages, func(e Pkg) bool {
		return e.Link == link && e.Remote == remote
	})
}

func (pkgs *Pkgs) AppendUnique(pkg Pkg) {
	if !pkgs.Contains(pkg.Link, pkg.Remote) {
		pkgs.Packages = append(pkgs.Packages, pkg)
	}
}

type Pkg struct {
	Name   string   `yaml:"name"`
	Link   string   `yaml:"link"`
	Remote string   `yaml:"remote"`
	Local  []string `yaml:"local"`
}

func Unmarshalled() (pkgs Pkgs, err error) {
	buffer, err := os.ReadFile(ConfigFilepath)
	if errors.Is(err, os.ErrNotExist) {
		os.Create(ConfigFilepath)
		buffer, err = os.ReadFile(ConfigFilepath)
	}
	if err != nil {
		return Pkgs{}, err
	}

	if err = yaml.Unmarshal(buffer, &pkgs); err != nil {
		return Pkgs{}, err
	}
	return pkgs, nil
}

func Marshall(pkgs Pkgs) error {
	marshal, err := yaml.Marshal(pkgs)
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigFilepath, marshal, hp.Perm)
}
