package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/action"
	"HeaderPuller/hp/pkg"
	"fmt"
	"github.com/urfave/cli/v2"
	"strconv"
)

var PullCmd = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Usage:   "pull headers in specified folder and update/create the config file",
	Description: `Usage: pull <repo-link> [optional arguments...]
There are 3 variations of this command:
	- pull <repo-link> - providing the repo link will copy every valid file from <repo-link>/include/ to ./include/
	- pull <repo-link> <file> - will copy that exact file if valid from <repo-link>/ to ./include/
	- pull <repo-link> <from> - will copy every valid file from <repo-link/<from>/ to ./<from>, which is by default ./include/
`,
	Action: func(cCtx *cli.Context) error {
		if !cCtx.Args().Present() {
			return hp.ErrNoArg
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
	Usage:   "Syncs local files to the latest remote version based on config file",
	Action:  action.Sync,
}

var RemoveCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm", "r"},
	Usage:   "Removes a package and updates the config file",
	Action: func(cCtx *cli.Context) error {
		if !cCtx.Args().Present() {
			return hp.ErrNoArg
		}

		arg := cCtx.Args().First()
		if _, err := strconv.ParseInt(arg, 10, 32); err == nil {
			return action.Remove(arg, action.IdMode)
		}

		if hp.IsRepoLink(arg) {
			return action.Remove(arg, action.LinkMode)
		}

		_, arg = hp.FilepathSplit(arg)
		return action.Remove(arg, action.NameMode)
	},
}

var WipeCmd = &cli.Command{
	Name:   "wipe",
	Usage:  "Removes all packages and then HeaderPuller itself from workspace",
	Action: action.Wipe,
}

var ListCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "Lists currently pulled packages in workspace",
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
