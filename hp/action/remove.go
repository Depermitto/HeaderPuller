package action

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/pkg"
	"os"
)

type rmMode int

const (
	IdMode rmMode = iota
	NameMode
	LinkMode
)

func Remove(configPkg pkg.ConfigPkg, mode rmMode) error {
	pkgs, err := pkg.Unmarshalled()
	if err != nil {
		return err
	}

	var filtered pkg.ConfigPkgs
	for _, p := range pkgs.Packages {
		if (mode == LinkMode && p.Link == configPkg.Link) ||
			(mode == IdMode && p.Id == configPkg.Id) ||
			(mode == NameMode && p.Name == configPkg.Name) {

			for _, filepath := range p.Local {
				os.Remove(filepath)
			}
			removeEmptyDirs(p.Remote)
		} else {
			filtered.Packages = append(filtered.Packages, p)
		}
	}
	filtered.Update()

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
	if len(dirs) == 0 {
		os.RemoveAll(dirname)
	}
}
