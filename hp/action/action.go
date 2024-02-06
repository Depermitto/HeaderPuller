package action

import (
	"HeaderPuller/hp"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"io"
	"os"
)

func createFileFromReader(reader io.Reader, filepath string) error {
	dirname, filename := hp.FilepathSplit(filepath)
	fmt.Println(hp.FileFmt(dirname, filename))

	os.MkdirAll(dirname, 0755)
	file, err := os.Create(hp.FileFmt(dirname, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func filesFromBilly(fs billy.Filesystem, dirname string) (files []billy.File) {
	infos, err := fs.ReadDir(dirname)
	if err != nil {
		return []billy.File{}
	}

	for _, info := range infos {
		if info.IsDir() {
			f := filesFromBilly(fs, hp.FileFmt(dirname, info.Name()))
			files = append(files, f...)
		} else {
			file, _ := fs.Open(hp.FileFmt(dirname, info.Name()))
			files = append(files, file)
		}
	}
	return files
}
