package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/internal/files"
	"HeaderPuller/hp/internal/pkg"
	"errors"
	"github.com/urfave/cli/v2"
)

var WipeCmd = &cli.Command{
	Name:  "wipe",
	Usage: "Removes all pulled packages and the the *hp.yaml* file itself.",
	Action: func(cCtx *cli.Context) error {
		if !pkg.Initialized() {
			return hp.ErrNotInWorkspace
		}

		if cCtx.Args().Present() {
			return hp.ErrArg
		}

		return Wipe()
	},
}

func Wipe() error {
	pkgs := pkg.Unmarshalled()
	defer pkg.UninitializeIfEmpty()

	for range pkgs.Packages {
		err := Remove("0", IdMode)
		if !errors.Is(err, hp.NoErrRemoved) {
			return err
		}
	}
	files.RemoveEmptyDirs(hp.IncludeDir)

	return hp.NoErrWiped
}
