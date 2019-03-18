# gl

Generate a LICENSE file for your project.

> NOTE: This is a practice project. Please use it at your own risk.

## Install

```
$ go get github.com/mika-f/gl/cmd/gl
```

## Description

Author and Year are automatically obtained from the your computer.

| Key    | Default Value                        |
| ------ | ------------------------------------ |
| Author | gitconfig (local, global and system) |
| Year   | current datetime                     |
| Output | `./LICENSE`                          |


## Information

```
NAME:
   gl - Generate a LICENSE file for your project

USAGE:
   gl <license> [--author author] [--year year] [--output path]

VERSION:
   0.1.0

AUTHOR:
   Fuyuno Mikazuki <https://github.com/mika-f>

COMMANDS:
     agpl       Generate a LICENSE file as GNU AGPLv3
     apache     Generate a LICENSE file as Apache License 2.0
     bsd2       Generate a LICENSE file as BSD 2-Clause "Simplified" License
     bsd3       Generate a LICENSE file as BSD 3-Clause "New" or "Revised" License
     eclipse    Generate a LICENSE file as Eclipse Public License 2.0
     gpl        Generate a LICENSE file as GNU GPLv3
     lgpl       Generate a LICENSE file as GNU LGPLv3
     lgpl2      Generate a LICENSE file as GNU LGPLv2.1
     mit        Generate a LICENSE file as MIT License
     mpl        Generate a LICENSE file as Mozilla Public License 2.0
     unlicense  Generate a LICENSE file as The Unlicense
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
