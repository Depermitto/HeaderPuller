package hp

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var UninstallCmd = &cli.Command{
	Name:    "uninstall",
	Aliases: []string{"u"},
	Usage:   "Removes the hp tool",
	Action: func(cCtx *cli.Context) error {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Are you sure? [Y/N] ")
		scanner.Scan()
		answer := scanner.Text()
		if strings.ToUpper(answer) != "Y" {
			return nil
		}

		path := fmt.Sprintf("%v/bin/hp", os.Getenv("GOPATH"))
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("%v\nhp executable must have been moved or already removed", err)
		}
		fmt.Println("Uninstalled successfully")
		return nil
	},
}
