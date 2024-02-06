package hp

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"os"
	"strings"
)

const (
	IncludeDir = "include"
	PathSep    = string(os.PathSeparator)
	Perm       = 0755
)

var (
	ErrNoArg               = errors.New("requires one argument")
	ErrNoFilesFound        = errors.New("no files found")
	ErrArg                 = errors.New("this configuration doesn't take any arguments")
	ErrNotInWorkspace      = errors.New("not in a HeaderPuller workspace")
	NoErrAlreadyDownloaded = errors.New("already downloaded this package")

	ValidExtensions = []string{
		".c", ".h",
		".cpp", ".c++", ".cxx", ".cc",
		".hpp", ".h++", ".hxx",
		".ii", ".iml",
		".rs",
	}
)

// Valid checks if fileFmt ends with one of ValidExtensions
func Valid(filename string) bool {
	for _, ext := range ValidExtensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func DefaultOptions(repoLink string) *git.CloneOptions {
	return &git.CloneOptions{
		URL:   repoLink,
		Depth: 1,
	}
}
