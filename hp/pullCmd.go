package hp

import (
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strings"
)

func fileFromReader(dst io.Reader, filepath string) error {
	i := strings.LastIndex(filepath, "/")
	var dirname, filename string
	if i == -1 {
		dirname, filename = includeDir, filepath
	} else {
		dirname, filename = filepath[:i], filepath[i+1:]
	}
	fmt.Println(fileFmt(dirname, filename))

	os.MkdirAll(dirname, 0755)
	file, err := os.Create(fileFmt(dirname, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, dst)
	return err
}

func getFiles(fs billy.Filesystem, dirname string) (files []billy.File) {
	infos, err := fs.ReadDir(dirname)
	if err != nil {
		return nil
	}

	for _, info := range infos {
		if info.IsDir() {
			f := getFiles(fs, fileFmt(dirname, info.Name()))
			files = append(files, f...)
		} else {
			file, _ := fs.Open(fileFmt(dirname, info.Name()))
			files = append(files, file)
		}
	}
	return files
}

var PullCmd = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Usage:   "pull headers in specified folder",
	Description: `Usage: pull <repo-link> [optional arguments...]
There are 3 variations of this command:
	- pull <repo-link> - providing the repo link will copy every valid fileFmt from <repo-link>/include/ to ./include/
	- pull <repo-link> <fileFmt> - will copy that exact fileFmt if valid from <repo-link>/<fileFmt> to ./include/
	- pull <repo-link> <from> - will copy every valid fileFmt from <repo-link/<from>/ to ./<from>, which is by default ./include/
`,
	Action: func(cCtx *cli.Context) error {
		repoLink, headerDir, err := PullLinks(cCtx)
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

		if Valid(headerDir) {
			header, err := fs.Open(headerDir)
			if err != nil {
				return err
			}
			defer header.Close()
			return fileFromReader(header, header.Name())
		}

		// Check if repo has a valid headers directory
		for _, file := range getFiles(fs, headerDir) {
			if !Valid(file.Name()) {
				continue
			}
			err := fileFromReader(file, file.Name())
			if err != nil {
				return err
			}
			file.Close()
		}
		return nil
	},
}
