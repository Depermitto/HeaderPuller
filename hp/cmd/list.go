package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/internal/pkg"
	"fmt"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:        "list",
	Aliases:     []string{"l"},
	Usage:       "Lists currently pulled packages in workspace",
	UsageText:   "hp list/l",
	Description: "List all installed packages along with their identifiers. Ids correspond to order the packages have been added by and names are git repository names stripped of the author.",
	Action: func(cCtx *cli.Context) error {
		if cCtx.Args().Present() {
			return hp.ErrArg
		}

		if !pkg.Initialized() {
			return hp.ErrNotInWorkspace
		}

		pkgs := pkg.Unmarshalled()
		for i, p := range pkgs.Packages {
			fmt.Printf("%v: %v\n", i, p.Name)
		}
		return nil
	},
}
