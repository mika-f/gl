package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func generate(license string, name string, year int, output string) {
	println(license)
	println(name)
	println(year)
	println(output)
}

func main() {
	var author string
	var year int
	var output string

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "author",
			Usage:       "author name",
			Destination: &author,
		},
		cli.IntFlag{
			Name:        "year",
			Usage:       "copyright year",
			Value:       time.Now().Year(),
			Destination: &year,
		},
		cli.StringFlag{
			Name:        "output",
			Usage:       "output path",
			Value:       "LICENSE",
			Destination: &output,
		},
	}

	app := cli.NewApp()
	app.Name = "gl"
	app.Usage = "Generate a LICENSE file for your project"
	app.UsageText = "gl <license> [--author author] [--year year]"
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "mit",
			Usage: "Generate LICENSE file as MIT",
			Flags: flags,
			Action: func(c *cli.Context) error {
				generate("mit", author, year, output)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
