package hp

import (
	"errors"
	"github.com/urfave/cli/v2"
	"strings"
)

const includeDir = "./include"

var (
	ErrRequiresArg = errors.New("requires one argument")
	ErrArgAmount   = errors.New("incorrect argument amount")

	ValidExtensions = []string{
		".c", ".h",
		".cpp", ".c++", ".cxx", ".cc",
		".hpp", ".h++", ".hxx",
		".ii", ".iml",
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

// PullLinks fetches repoLink, fromDir and intoDir based on Args() from *cli.Context. One
// arg is required - for the repo. fromDir and intoDir default to ./include/.
func PullLinks(cCtx *cli.Context) (repoLink string, fromDir string, err error) {
	if !cCtx.Args().Present() {
		return "", "", ErrRequiresArg
	}

	repoLink = cCtx.Args().Get(0)
	fromDir = ifEmpty(cCtx.Args().Get(1), includeDir)
	return repoLink, fromDir, nil
}

func ifEmpty(filename string, other string) string {
	if len(filename) == 0 {
		return other
	}
	return filename
}

func fileFmt(pathParts ...string) (filename string) {
	for i, s := range pathParts {
		filename += s
		if i < len(pathParts)-1 {
			filename += "/"
		}
	}
	return filename
}
