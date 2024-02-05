package hp

import (
	"errors"
	"github.com/urfave/cli/v2"
	"strings"
)

const includeDir = "./include/"

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

// Valid checks if filename ends with one of ValidExtensions
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
func PullLinks(cCtx *cli.Context) (repoLink string, fromDir string, intoDir string, err error) {
	if !cCtx.Args().Present() {
		return "", "", "", ErrRequiresArg
	}

	repoLink = cCtx.Args().Get(0)
	fromDir = defaultIfEmpty(cCtx.Args().Get(1))
	intoDir = defaultIfEmpty(cCtx.Args().Get(2))
	return repoLink, fromDir, intoDir, nil
}

func defaultIfEmpty(filename string) string {
	if len(filename) == 0 {
		return includeDir
	}
	return filename
}
