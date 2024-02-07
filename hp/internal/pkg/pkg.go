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

func Initialized() bool {
	_, err := os.ReadFile(ConfigFilepath)
	return !errors.Is(err, os.ErrNotExist)
}

func UninitializeIfEmpty() {
	if Initialized() {
		pkgs := Unmarshalled()
		if len(pkgs.Packages) == 0 {
			os.Remove(ConfigFilepath)
		}
	}
}

func Unmarshalled() (pkgs Pkgs) {
	if !Initialized() {
		os.Create(ConfigFilepath)
	}

	buffer, _ := os.ReadFile(ConfigFilepath)
	_ = yaml.Unmarshal(buffer, &pkgs)
	return pkgs
}

func Marshall(pkgs Pkgs) {
	marshal, _ := yaml.Marshal(pkgs)
	_ = os.WriteFile(ConfigFilepath, marshal, hp.Perm)
}
