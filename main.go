package main

import (
	"HeaderPuller/hp"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.App{Name: "HeaderPuller", Commands: []*cli.Command{echoCmd, hp.PullCmd, hp.UninstallCmd}}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var echoCmd = &cli.Command{
	Name:    "echo",
	Aliases: []string{"e"},
	Usage:   "print argument to the stdout",
	Action: func(cCtx *cli.Context) error {
		if !cCtx.Args().Present() {
			return hp.ErrRequiresArg
		}
		fmt.Println(cCtx.Args().First())
		return nil
	},
}
