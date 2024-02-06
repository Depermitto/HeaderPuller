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

	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	if slices.ContainsFunc(pkgs.Packages, func(e pkg.ConfigPkg) bool {
		return e.Link == repoLink
	}) {
		return hp.NoErrAlreadyDownloaded
	}

	fs := memfs.New()
	storer := memory.NewStorage()
	_, err = git.Clone(storer, fs, &git.CloneOptions{
		URL:   repoLink,
		Depth: 1,
	})
	if err != nil {
		return err
	}

	// Check if repo has a valid headers directory
	var files []string
	if hp.Valid(headerDir) {
		files = append(files, hp.FileFmt(hp.IncludeDir, headerDir))
	}

	for _, file := range filesFromBilly(fs, headerDir) {
		if !hp.Valid(file.Name()) {
			continue
		}

		if err = createFileFromReader(file, file.Name()); err != nil {
			return err
		}
		file.Close()
		files = append(files, file.Name())
	}

	_, name := hp.FilepathSplit(repoLink)
	configPkg := pkg.ConfigPkg{
		Name:   name,
		Link:   repoLink,
		Remote: headerDir,
		Local:  files,
	}
	pkgs.Packages = append(pkgs.Packages, configPkg)
	pkgs.Update()

	return pkg.Marshall(pkgs)
}
