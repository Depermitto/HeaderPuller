package cmd

import (
	"HeaderPuller/hp/internal/pkg"
	"fmt"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:        "list",
	Aliases:     []string{"l"},
	Usage:       "Lists currently pulled packages in workspace",
	Description: "List all installed packages along with their identifiers. Ids correspond to order the packages have been added by and names are git repository names stripped of the author.",
	Action: func(cCtx *cli.Context) error {
		localPkgs, err := pkg.Unmarshalled()
		if err != nil {
			return err
		}

		for i, p := range localPkgs.Packages {
			fmt.Printf("%v: %v\n", i, p.Name)
		}
		return nil
	},
}

var VersionCmd = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Get program version",
	Action: func(cCtx *cli.Context) error {
		fmt.Printf("hp version %v\n", cCtx.App.Version)
		return nil
	},
}
