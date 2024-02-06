package action

import (
	"HeaderPuller/hp/pkg"
	"github.com/urfave/cli/v2"
)

// Sync only checks files in the default folder ./include/!!
var Sync cli.ActionFunc = func(cCtx *cli.Context) error {
	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	for _, p := range pkgs.Packages {
		if err := Pull(p.Link, p.Remote); err != nil {
			return err
		}
	}
	return nil
}
