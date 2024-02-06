package action

import (
	"HeaderPuller/hp"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strings"
)

type action func(cCtx *cli.Context) error

// pullLinks fetches repoLink, fromDir and intoDir based on Args() from *cli.Context. One
// arg is required - for the repo. fromDir and intoDir default to ./include/.
func pullLinks(cCtx *cli.Context) (repoLink string, fromDir string, err error) {
	if !cCtx.Args().Present() {
		return "", "", hp.ErrRequiresArg
	}

	repoLink = cCtx.Args().Get(0)
	fromDir = cCtx.Args().Get(1)
	if len(cCtx.Args().Get(1)) == 0 {
		fromDir = hp.IncludeDir
	}
	return repoLink, fromDir, nil
}

func fileFmt(pathParts ...string) (filename string) {
	for i, s := range pathParts {
		filename += s
		if i < len(pathParts)-1 {
			filename += hp.PathSep
		}
	}
	return filename
}

func createFileFromReader(reader io.Reader, filepath string) error {
	i := strings.LastIndex(filepath, hp.PathSep)
	var dirname, filename string
	if i == -1 {
		dirname, filename = hp.IncludeDir, filepath
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

	_, err = io.Copy(file, reader)
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
