package hp

import (
	"errors"
	"os"
)

const (
	IncludeDir = "include"
	PathSep    = string(os.PathSeparator)
	Perm       = 0755
)

var (
	ErrNoArg             = errors.New("requires one argument")
	ErrNoFilesFound      = errors.New("no files found")
	ErrArg               = errors.New("this configuration doesn't take any arguments")
	ErrNotInWorkspace    = errors.New("not in hp workspace")
	ErrAlreadyDownloaded = errors.New("already downloaded this package")
	NoErrWiped           = errors.New("wipe successful")
)
