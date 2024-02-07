package cmd

import (
	"HeaderPuller/hp/internal/config"
	"HeaderPuller/hp/internal/pkg"
	"github.com/urfave/cli/v2"
)

var SyncCmd = &cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Usage:   "Updates every package to the latest version.",
	Action:  Sync,
}

// Sync only checks files in the default folder ./include/!!
func Sync(*cli.Context) error {
	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	for _, p := range pkgs.Packages {
		if err := Pull(p.Link, p.Remote, config.New(config.WithForce)); err != nil {
			return err
		}
	}
	return nil
}
