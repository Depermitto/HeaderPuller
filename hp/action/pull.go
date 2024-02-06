package action

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/pkg"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/urfave/cli/v2"
	"slices"
)

var Pull = func(repoLink string, headerDir string, cCtx *cli.Context) error {
	if !hp.IsRepoLink(repoLink) {
		repoLink = "https://" + repoLink
	}

	localPkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	if !cCtx.Bool("force") && slices.ContainsFunc(localPkgs.Packages, func(e pkg.ConfigPkg) bool {
		return e.Link == repoLink && e.Remote == headerDir
	}) {
		return hp.NoErrAlreadyDownloaded
	}

	fs := memfs.New()
	storer := memory.NewStorage()
	if _, err = git.Clone(storer, fs, hp.DefaultOptions(repoLink)); err != nil {
		return err
	}

	var billyFiles = filesFromBilly(fs, headerDir, cCtx)
	if len(billyFiles) == 0 {
		_ = Wipe(nil) // This is to remove empty hp.yaml
		return hp.ErrNoFilesFound
	}

	var filepaths []string
	for _, file := range billyFiles {
		if !hp.ValidFile(fs, file.Name(), cCtx) {
			continue
		}

		if err = createFrom(file, file.Name()); err != nil {
			return err
		}
		file.Close()
		filepaths = append(filepaths, file.Name())
	}

	_, name := hp.FilepathSplit(repoLink)
	configPkg := pkg.ConfigPkg{
		Name:   name,
		Link:   repoLink,
		Remote: headerDir,
		Local:  filepaths,
	}
	localPkgs.Append(configPkg)
	return pkg.Marshall(localPkgs)
}
