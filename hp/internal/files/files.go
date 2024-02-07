package files

import (
	"HeaderPuller/hp"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"io"
	"os"
	"strings"
)

var ValidExtensions = []string{
	".c", ".h",
	".cpp", ".c++", ".cxx", ".cc",
	".hpp", ".h++", ".hxx",
	".ii", ".iml",
	".rs",
}

// IsValid checks if filename ends with one of ValidExtensions. If content
// at filename is directory, IsValid will return false.
func IsValid(fs billy.Filesystem, filename string) bool {
	if _, err := fs.Open(filename); err != nil {
		return false
	}

	for _, ext := range ValidExtensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

// CreateCopy creates a file with contents identical to reader at filepath
func CreateCopy(reader io.Reader, filepath string) error {
	dirname, filename := FilepathSplit(filepath)
	fmt.Println(FileFmt(dirname, filename))

	os.MkdirAll(dirname, hp.Perm)
	file, err := os.Create(FileFmt(dirname, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

// ReadDirRecur recursively searches fs starting at dirname and returns a slice
// of found files. This function will only consider invalid files (judged by IsValid).
// If dirname is file, it will return a [1]billy.File with that one file.
func ReadDirRecur(fs billy.Filesystem, dirname string) (files []billy.File) {
	if IsValid(fs, dirname) {
		path := FileFmt(hp.IncludeDir, dirname)
		// Rename is used to save os.FileInfo.Name() to hp.IncludeDir/filename
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
			f := ReadDirRecur(fs, FileFmt(dirname, info.Name()))
			files = append(files, f...)
		} else {
			file, _ := fs.Open(FileFmt(dirname, info.Name()))
			files = append(files, file)
		}
	}
	return files
}

func FilepathSplit(filepath string) (dirname string, filename string) {
	i := strings.LastIndex(filepath, hp.PathSep)
	if i == -1 {
		dirname, filename = hp.IncludeDir, filepath
	} else {
		dirname, filename = filepath[:i], filepath[i+1:]
	}
	return dirname, filename
}

func FileFmt(pathParts ...string) (filename string) {
	for i, s := range pathParts {
		filename += s
		if i < len(pathParts)-1 {
			filename += hp.PathSep
		}
	}
	return filename
}

func RemoveEmptyDirs(dirname string) {
	dirs, _ := os.ReadDir(dirname)
	for _, dir := range dirs {
		if dir.IsDir() {
			RemoveEmptyDirs(FileFmt(dirname, dir.Name()))
		}
	}

	dirs, _ = os.ReadDir(dirname)
	if dirs == nil || len(dirs) == 0 {
		os.Remove(dirname)
	}
}
