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
	Description: `usage: pull <repo-link> [optional args...]
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
	Usage:  "upon confirmation, wipes hp from the computer entirely",
	Action: action.Uninstall,
}

var SyncCmd = &cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Usage:   "updates every package to the latest version.",
	Action:  action.Sync,
}

var RemoveCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm", "r"},
	Usage:   "removes a package and updates the config file",
	Description: `removes files and folders of all header files encompassing a package. There are 3 variations of this command:
- remove <id> - delete by id
- remove <name> - remove by package name
- remove <repo-link> - remove by repository link

The ids and packages names are provided by the list command.
`,
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
	Usage:  "removes all pulled packages and the the *hp.yaml* file itself.",
	Action: action.Wipe,
}

var ListCmd = &cli.Command{
	Name:        "list",
	Aliases:     []string{"l"},
	Usage:       "lists currently pulled packages in workspace",
	Description: "list all installed packages along with their identifiers. Ids correspond to order the packages have been added by and names are git repository names stripped of the author.",
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
	Usage:   "get program version",
	Action: func(cCtx *cli.Context) error {
		fmt.Printf("hp version %v\n", cCtx.App.Version)
		return nil
	},
}
