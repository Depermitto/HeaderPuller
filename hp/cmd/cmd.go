package cmd

import (
	"HeaderPuller/hp/action"
	"github.com/urfave/cli/v2"
)

var PullCmd = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Usage:   "pull headers in specified folder",
	Description: `Usage: pull <repo-link> [optional arguments...]
There are 3 variations of this command:
	- pull <repo-link> - providing the repo link will copy every valid fileFmt from <repo-link>/include/ to ./include/
	- pull <repo-link> <fileFmt> - will copy that exact fileFmt if valid from <repo-link>/<fileFmt> to ./include/
	- pull <repo-link> <from> - will copy every valid fileFmt from <repo-link/<from>/ to ./<from>, which is by default ./include/
`,
	Action: cli.ActionFunc(action.Pull),
}

var UninstallCmd = &cli.Command{
	Name:   "uninstall",
	Usage:  "Removes the hp tool",
	Action: cli.ActionFunc(action.Uninstall),
}
