package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/urfave/cli"
)

var wg sync.WaitGroup
var dir = "files/"

// contants for min and max int
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	app := cli.NewApp()
	app.Name = "MyCli"
	app.Usage = "To check csv files"
	app.Version = "1.0.0"

	// flags for option command
	var fileLocation string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file",
			Value:       "files/",
			Usage:       "specify files location",
			Destination: &fileLocation,
		},
	}

	app.Commands = []cli.Command{

		//first command load column data
		{
			Name:    "column",
			Aliases: []string{"c"},
			Usage:   "Show load result",
			Action: func(c *cli.Context) error {
				input := c.Args().Get(0)
				columns := strings.Split(input, ",")

				if len(input) > 1 {
					csvReaders(columns, fileLocation)
				} else if len(input) == 1 {
					csvReader(columns[0], fileLocation)

				} else {
					fmt.Println("Please input one column name")
					return nil

				}

				return nil
			},
		},

		// second command to print current directory
		{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "Print current directory",
			Action: func(c *cli.Context) error {
				// test2
				dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(dir)
				return nil
			},
		},
	}

	app.Run(os.Args)

}
