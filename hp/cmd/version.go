package cmd

import (
	"HeaderPuller/hp"
	"fmt"
	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:      "version",
	Aliases:   []string{"v"},
	Usage:     "Get program version",
	UsageText: "hp version/v",
	Action: func(cCtx *cli.Context) error {
		if cCtx.Args().Present() {
			return hp.ErrArg
		}

		fmt.Printf("hp version %v\n", cCtx.App.Version)
		return nil
	},
}
