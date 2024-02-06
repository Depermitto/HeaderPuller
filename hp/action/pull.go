package action

import (
	"HeaderPuller/hp"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/urfave/cli/v2"
)

var Pull action = func(cCtx *cli.Context) error {
	repoLink, headerDir, err := pullLinks(cCtx)
	if err != nil {
		return err
	}

	fs := memfs.New()
	storer := memory.NewStorage()
	_, err = git.Clone(storer, fs, &git.CloneOptions{
		URL:   "https://" + repoLink,
		Depth: 1,
	})
	if err != nil {
		return err
	}

	if hp.Valid(headerDir) {
		header, err := fs.Open(headerDir)
		if err != nil {
			return err
		}
		defer header.Close()
		return createFileFromReader(header, header.Name())
	}

	// Check if repo has a valid headers directory
	for _, file := range getFiles(fs, headerDir) {
		if !hp.Valid(file.Name()) {
			continue
		}
		err := createFileFromReader(file, file.Name())
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
