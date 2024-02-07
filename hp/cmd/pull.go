package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/internal/files"
	"HeaderPuller/hp/internal/ops"
	"HeaderPuller/hp/internal/pkg"
	"HeaderPuller/hp/internal/repo"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/urfave/cli/v2"
)

var PullCmd = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "ignore-extensions",
			Aliases: []string{"i"},
			Usage:   "ignore file extensions, allows pulling any type of file.",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "force pull.",
		}},
	Usage: "Pull headers in specified folder and update/create the ops file",
	Description: `There are 3 variations of this command:
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

		conf := ops.New()
		conf.SetForce(cCtx.Bool("force"))
		conf.SetIgnoreExt(cCtx.Bool("ignore-extensions"))

		return Pull(repoLink, headerDir, conf)
	},
}

func Pull(repoLink string, headerDir string, c *ops.Config) error {
	if !repo.IsRepoLink(repoLink) {
		repoLink = "https://" + repoLink
	}

	localPkgs := pkg.Unmarshalled()
	defer pkg.UninitializeIfEmpty()
	if !c.Force() && localPkgs.Contains(repoLink, headerDir) {
		return hp.ErrAlreadyDownloaded
	}

	fs := memfs.New()
	storer := memory.NewStorage()
	if _, err := git.Clone(storer, fs, repo.DefaultOptions(repoLink)); err != nil {
		return err
	}

	var foundFiles = files.ReadDirRecur(fs, headerDir)
	if len(foundFiles) == 0 {
		return hp.ErrNoFilesFound
	}

	var filepaths []string
	for _, file := range foundFiles {
		if c.IgnoreExt() || files.IsValid(fs, file.Name()) {
			if err := files.CreateCopy(file, file.Name()); err != nil {
				return err
			}
			file.Close()
			filepaths = append(filepaths, file.Name())
		}
	}

	_, name := files.FilepathSplit(repoLink)
	configPkg := pkg.Pkg{
		Name:   name,
		Link:   repoLink,
		Remote: headerDir,
		Local:  filepaths,
	}
	localPkgs.AppendUnique(configPkg)
	pkg.Marshall(localPkgs)

	return hp.NoErrPulled
}
