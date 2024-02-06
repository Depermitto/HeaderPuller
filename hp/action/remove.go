package action

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/pkg"
	"os"
	"strconv"
)

type rmMode int

const (
	IdMode rmMode = iota
	NameMode
	LinkMode
)

func Remove(arg string, mode rmMode) error {
	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	var filtered pkg.ConfigPkgs
	for i, p := range pkgs.Packages {
		if (mode == LinkMode && p.Link == arg) ||
			(mode == IdMode && strconv.FormatInt(int64(i), 10) == arg) ||
			(mode == NameMode && p.Name == arg) {

			for _, filepath := range p.Local {
				os.Remove(filepath)
			}
			removeEmptyDirs(p.Remote)
		} else {
			filtered.Packages = append(filtered.Packages, p)
		}
	}

	if len(filtered.Packages) == 0 {
		return os.Remove("hp.yaml")
	}
	return pkg.Marshall(filtered)
}

func removeEmptyDirs(dirname string) {
	dirs, _ := os.ReadDir(dirname)
	for _, dir := range dirs {
		if dir.IsDir() {
			removeEmptyDirs(hp.FileFmt(dirname, dir.Name()))
		}
	}

	dirs, _ = os.ReadDir(dirname)
	if dirs == nil || len(dirs) == 0 {
		os.RemoveAll(dirname)
	}
}
