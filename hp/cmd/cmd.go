package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/action"
	"HeaderPuller/hp/pkg"
	"github.com/urfave/cli/v2"
	"strconv"
)

var PullCmd = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Usage:   "pull headers in specified folder",
	Description: `Usage: pull <repo-link> [optional arguments...]
There are 3 variations of this command:
	- pull <repo-link> - providing the repo link will copy every valid file from <repo-link>/include/ to ./include/
	- pull <repo-link> <file> - will copy that exact file if valid from <repo-link>/ to ./include/
	- pull <repo-link> <from> - will copy every valid file from <repo-link/<from>/ to ./<from>, which is by default ./include/
`,
	Action: func(cCtx *cli.Context) error {
		if !cCtx.Args().Present() {
			return hp.ErrRequiresArg
		}

		repoLink, headerDir := cCtx.Args().First(), cCtx.Args().Get(1)
		if len(cCtx.Args().Get(1)) == 0 {
			headerDir = hp.IncludeDir
		}
		return action.Pull(repoLink, headerDir)
	},
}

var UninstallCmd = &cli.Command{
	Name:   "uninstall",
	Usage:  "Removes the hp tool",
	Action: action.Uninstall,
}

var SyncCmd = &cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Usage:   "Syncs local files to the latest remote version",
	Action:  action.Sync,
}

var RemoveCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"r", "rm"},
	Usage:   "Removes a package and updates the config file",
	Action: func(cCtx *cli.Context) error {
		if !cCtx.Args().Present() {
			return hp.ErrRequiresArg
		}

		arg := cCtx.Args().First()
		if id, err := strconv.ParseInt(arg, 10, 32); err == nil {
			return action.Remove(pkg.ConfigPkg{Id: int(id)}, action.IdMode)
		}

		if hp.IsRepoLink(arg) {
			return action.Remove(pkg.ConfigPkg{Link: arg}, action.LinkMode)
		}

		_, arg = hp.FilepathSplit(arg)
		return action.Remove(pkg.ConfigPkg{Name: arg}, action.NameMode)
	},
}
