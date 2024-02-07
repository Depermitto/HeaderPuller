package cmd

import (
	"HeaderPuller/hp"
	"errors"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var UpdateCmd = &cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Usage:   "Update to the latest git commit",
	Action: func(cCtx *cli.Context) error {
		executablePath, err := exec.LookPath("hp")
		if err != nil {
			return errors.New("hp executable not found, update unable to finish")
		}

		tmp := strconv.Itoa(time.Now().Nanosecond())
		if err := exec.Command("git", "clone", "https://github.com/Depermitto/HeaderPuller", tmp).Run(); err != nil {
			return errors.New("cannot clone repo, update unable to finish")
		}
		defer os.RemoveAll(tmp)

		if err := exec.Command("go", "build", "-C", tmp, "-o", executablePath).Run(); err != nil {
			return errors.New("errors found in building the executable, update unable to finish")
		}
		return hp.NoErrUpdated
	},
}
