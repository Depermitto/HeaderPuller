package cmd

import (
	"HeaderPuller/hp"
	"HeaderPuller/hp/internal/ops"
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strings"
)

var UninstallCmd = &cli.Command{
	Name: "uninstall",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:               "no-confirmation",
			Usage:              "doesn't ask for uninstall confirmation.",
			DisableDefaultText: true,
		},
	},
	Usage:     "Upon confirmation, wipes hp from the computer entirely",
	UsageText: "hp uninstall",
	Action: func(cCtx *cli.Context) error {
		if cCtx.Args().Present() {
			return hp.ErrArg
		}

		conf := ops.New()
		conf.SetNoConfirm(cCtx.Bool("no-confirmation"))

		return Uninstall(conf)
	},
}

func Uninstall(c *ops.Config) error {
	if !c.NoConfirm() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Are you sure? [Y/N] ")
		scanner.Scan()
		answer := scanner.Text()
		if strings.ToUpper(answer) != "Y" {
			return nil
		}
	}

	path, err := exec.LookPath("hp")
	if err != nil {
		return fmt.Errorf("%v\nhp executable must have been moved or already removed", err)
	}

	os.Remove(path)
	return hp.NoErrUninstalled
}
