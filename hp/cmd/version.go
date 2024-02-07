package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Get program version",
	Action: func(cCtx *cli.Context) error {
		fmt.Printf("hp version %v\n", cCtx.App.Version)
		return nil
	},
}
