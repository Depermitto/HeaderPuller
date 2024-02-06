package action

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/pkg"
	"errors"
	"github.com/urfave/cli/v2"
	"os"
)

var Wipe cli.ActionFunc = func(cCtx *cli.Context) error {
	if cCtx != nil && cCtx.Args().Present() {
		return hp.ErrArg
	}

	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	if len(pkgs.Packages) == 0 {
		err := os.Remove("hp.yaml")
		if errors.Is(err, os.ErrNotExist) {
			return hp.ErrNotInWorkspace
		}
		return err
	}

	for range pkgs.Packages {
		if err := Remove("0", IdMode); err != nil {
			return err
		}
	}
	removeEmptyDirs("include")
	return nil
}
