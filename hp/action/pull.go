package action

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/pkg"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"slices"
)

var Pull = func(repoLink string, headerDir string) error {
	if !hp.IsRepoLink(repoLink) {
		repoLink = "https://" + repoLink
	}

	localPkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	if slices.ContainsFunc(localPkgs.Packages, func(e pkg.ConfigPkg) bool {
		return e.Link == repoLink && e.Remote == headerDir
	}) {
		return hp.NoErrAlreadyDownloaded
	}

	fs := memfs.New()
	storer := memory.NewStorage()
	if _, err = git.Clone(storer, fs, hp.DefaultOptions(repoLink)); err != nil {
		return err
	}

	var billyFiles = filesFromBilly(fs, headerDir)
	if len(billyFiles) == 0 {
		return hp.ErrNoFilesFound
	}

	var filepaths []string
	for _, file := range billyFiles {
		if !hp.Valid(file.Name()) {
			continue
		}

		if err = createFileFromReader(file, file.Name()); err != nil {
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
	localPkgs.Packages = append(localPkgs.Packages, configPkg)

	return pkg.Marshall(localPkgs)
}
