package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/internal/ops"
	"HeaderPuller/hp/internal/pkg"
	"errors"
	"github.com/urfave/cli/v2"
)

var SyncCmd = &cli.Command{
	Name:      "sync",
	Aliases:   []string{"s"},
	Usage:     "Updates every package to the latest version",
	UsageText: "hp sync/s",
	Action: func(cCtx *cli.Context) error {
		if cCtx.Args().Present() {
			return hp.ErrArg
		}

		if !pkg.Initialized() {
			return hp.ErrNotInWorkspace
		}

		return Sync()
	},
}

// Sync only checks files in the default folder ./include/!!
func Sync() error {
	pkgs := pkg.Unmarshalled()
	for _, p := range pkgs.Packages {
		err := Pull(p.Link, p.Remote, ops.New(ops.WithForce))
		if !errors.Is(err, hp.NoErrPulled) {
			return err
		}
	}
	return hp.NoErrSynced
}
