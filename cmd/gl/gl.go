package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gl"
	app.Usage = "Generate a LICENSE file from command line"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
