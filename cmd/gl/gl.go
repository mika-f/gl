package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/urfave/cli"
)

type license struct {
	Alias string
	Name  string
}

const usage = "Generate a LICENSE file as %s"

func generate(license string, name string, year int, output string) {
	println(license)
	println(name)
	println(year)
	println(output)
}

func gitconfig(option string) string {
	buf, err := exec.Command("git", "config", "--"+option, "user.name").Output()
	if err != nil {
		return ""
	}
	return string(buf)
}

// get author name from gitconfig
// local -> global -> system
func defaultAuthor() string {
	name := gitconfig("local")
	if name != "" {
		return strings.TrimSpace(name)
	}

	name = gitconfig("global")
	if name != "" {
		return strings.TrimSpace(name)
	}

	name = gitconfig("system")
	if name != "" {
		return strings.TrimSpace(name)
	}

	panic(errors.New("could not detect username"))
}

func main() {
	var author string
	var year int
	var output string

	licenses := []license{
		license{
			Alias: "agpl",
			Name:  "GNU AGPLv3",
		},
		license{
			Alias: "apache",
			Name:  "Apache License 2.0",
		},
		license{
			Alias: "bsd2",
			Name:  "BSD 2-Clause \"Simplified\" License",
		},
		license{
			Alias: "bsd3",
			Name:  "BSD 3-Clause \"New\" or \"Revised\" License",
		},
		license{
			Alias: "eclipse",
			Name:  "Eclipse Public License 2.0",
		},
		license{
			Alias: "gpl",
			Name:  "GNU GPLv3",
		},
		license{
			Alias: "lgpl",
			Name:  "GNU LGPLv3",
		},
		license{
			Alias: "lgpl2",
			Name:  "GNU LGPLv2.1",
		},
		license{
			Alias: "mit",
			Name:  "MIT License",
		},
		license{
			Alias: "mpl",
			Name:  "Mozilla Public License 2.0",
		},
		license{
			Alias: "unlicense",
			Name:  "The Unlicense",
		},
	}

	action := func(c *cli.Context) error {
		generate(c.Command.Name, author, year, output)
		return nil
	}

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "author",
			Usage:       "author name",
			Value:       defaultAuthor(),
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

	commands := []cli.Command{}
	for _, license := range licenses {
		command := cli.Command{
			Name:   license.Alias,
			Usage:  fmt.Sprintf(usage, license.Name),
			Flags:  flags,
			Action: action,
		}
		commands = append(commands, command)
	}

	app := cli.NewApp()
	app.Name = "gl"
	app.Usage = "Generate a LICENSE file for your project"
	app.UsageText = "gl <license> [--author author] [--year year]"
	app.Commands = commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
