package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/internal/files"
	"HeaderPuller/hp/internal/pkg"
	"HeaderPuller/hp/internal/repo"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

type rmMode int

const (
	IdMode rmMode = iota
	NameMode
	LinkMode
)

var RemoveCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm", "r"},
	Usage:   "Removes a package and updates the config file",
	Description: `Removes files and folders of all header files encompassing a package. There are 3 variations of this command:
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
			return Remove(arg, IdMode)
		}

		if repo.IsRepoLink(arg) {
			return Remove(arg, LinkMode)
		}

		_, arg = files.FilepathSplit(arg)
		return Remove(arg, NameMode)
	},
}

func Remove(arg string, mode rmMode) error {
	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	var filtered pkg.Pkgs
	for i, p := range pkgs.Packages {
		if (mode == LinkMode && p.Link == arg) ||
			(mode == IdMode && strconv.FormatInt(int64(i), 10) == arg) ||
			(mode == NameMode && p.Name == arg) {

			for _, filepath := range p.Local {
				os.Remove(filepath)
			}
			files.RemoveEmptyDirs(p.Remote)
		} else {
			filtered.Packages = append(filtered.Packages, p)
		}
	}

	if len(filtered.Packages) == 0 {
		return os.Remove(pkg.ConfigFilepath)
	}
	return pkg.Marshall(filtered)
}