package hp

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

type File struct {
	billy        billy.File
	strippedName string
}

func createWithContent(dst io.Reader, filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	_, err := io.Copy(file, dst)
	return err
}

func getFiles(fs billy.Filesystem, dirname string) (files []File) {
	infos, err := fs.ReadDir(dirname)
	if err != nil {
		return nil
	}

	for _, info := range infos {
		file, _ := fs.Open(dirname + info.Name())
		files = append(files, File{
			billy:        file,
			strippedName: info.Name(),
		})
	}
	return files
}

var PullCmd = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Usage:   "pull headers in <repo-link>/include folder",
	Action: func(cCtx *cli.Context) error {
		repoLink, headerDir, intoDir, err := PullLinks(cCtx)
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

		// If intoDir doesn't exist, create one
		_ = os.Mkdir(intoDir, 0755)

		// Check if repo has a valid headers directory
		for _, file := range getFiles(fs, headerDir) {
			if !Valid(file.strippedName) {
				continue
			}

			err := createWithContent(file.billy, intoDir+file.strippedName)
			if err != nil {
				return err
			}
			file.billy.Close()
		}
		return nil
	},
}
