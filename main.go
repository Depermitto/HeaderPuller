package main

import (
	"HeaderPuller/hp/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.App{Name: "HeaderPuller", Commands: []*cli.Command{cmd.PullCmd, cmd.UninstallCmd}}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
