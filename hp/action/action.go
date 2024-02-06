package action

import (
	"HeaderPuller/hp"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func createFrom(reader io.Reader, filepath string) error {
	dirname, filename := hp.FilepathSplit(filepath)
	fmt.Println(hp.FileFmt(dirname, filename))

	os.MkdirAll(dirname, hp.Perm)
	file, err := os.Create(hp.FileFmt(dirname, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func filesFromBilly(fs billy.Filesystem, dirname string, cCtx *cli.Context) (files []billy.File) {
	if hp.ValidFile(fs, dirname, cCtx) {
		path := hp.FileFmt(hp.IncludeDir, dirname)
		_ = fs.Rename(dirname, path)
		file, _ := fs.Open(path)
		return []billy.File{file}
	}

	infos, err := fs.ReadDir(dirname)
	if err != nil {
		return []billy.File{}
	}

	for _, info := range infos {
		if info.IsDir() {
			f := filesFromBilly(fs, hp.FileFmt(dirname, info.Name()), cCtx)
			files = append(files, f...)
		} else {
			file, _ := fs.Open(hp.FileFmt(dirname, info.Name()))
			files = append(files, file)
		}
	}
	return files
}
