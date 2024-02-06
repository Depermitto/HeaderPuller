package main

import (
	"HeaderPuller/hp/cmd"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Name:        "hp",
		Usage:       "pull header files from git repositories to current workspace",
		UsageText:   "hp command [args...]",
		Version:     "0.1",
		HideVersion: true,
		Commands: []*cli.Command{
			cmd.PullCmd,
			cmd.ListCmd,
			cmd.SyncCmd,
			cmd.RemoveCmd,
			cmd.WipeCmd,
			cmd.VersionCmd,
			cmd.UninstallCmd,
		}}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
