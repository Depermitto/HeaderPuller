package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.App{}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error running hp, %v", err)
	}
}
