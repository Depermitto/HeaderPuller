package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/internal/files"
	"HeaderPuller/hp/internal/pkg"
	"github.com/urfave/cli/v2"
	"os"
)

var WipeCmd = &cli.Command{
	Name:  "wipe",
	Usage: "Removes all pulled packages and the the *hp.yaml* file itself.",
	Action: func(cCtx *cli.Context) error {
		if cCtx.Args().Present() {
			return hp.ErrArg
		}

		return Wipe()
	},
}

func Wipe() error {
	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	if len(pkgs.Packages) == 0 {
		os.Remove(pkg.ConfigFilepath) // This is to remove empty config file
		return hp.ErrNotInWorkspace
	}

	for range pkgs.Packages {
		if err := Remove("0", IdMode); err != nil {
			return err
		}
	}
	files.RemoveEmptyDirs("include")
	return hp.NoErrWiped
}
