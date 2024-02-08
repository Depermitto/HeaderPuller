package hp

import (
	"errors"
	"os"
)

const (
	IncludeDir = "include"
	PathSep    = string(os.PathSeparator)
	Perm       = 0755
	RepoLink   = "https://github.com/Depermitto/HeaderPuller"
)

var (
	ErrNoArg             = errors.New("no arguments provided")
	ErrArg               = errors.New("command doesn't take any arguments")
	ErrNoFilesFound      = errors.New("no files found")
	ErrNotInWorkspace    = errors.New("not in hp workspace")
	ErrAlreadyDownloaded = errors.New("already downloaded this package")
	NoErrPulled          = errors.New("pulled")
	NoErrWiped           = errors.New("wiped")
	NoErrRemoved         = errors.New("removed")
	NoErrSynced          = errors.New("synced")
	NoErrUpdated         = errors.New("updated")
	NoErrUninstalled     = errors.New("uninstalled")
)
