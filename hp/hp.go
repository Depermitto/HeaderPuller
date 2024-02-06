package hp

import (
	"errors"
	"os"
	"strings"
)

const (
	IncludeDir = "include"
	PathSep    = string(os.PathSeparator)
)

var (
	ErrRequiresArg         = errors.New("requires one argument")
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
