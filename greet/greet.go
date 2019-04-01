package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Name = "greet"
	app.Usage = "fight the loneliness"

	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	// app.Action = func(c *cli.Context) error {
	// 	name := "Nefertiti"
	// 	if c.NArg() > 0 {
	// 		name = c.Args().Get(0)
	// 	}

	// 	if c.String("lang") == "spanish" {
	// 		fmt.Println("Hola", name)
	// 	} else if c.String("lang") == "chinese" {
	// 		fmt.Println("你好", name)
	// 	} else {
	// 		fmt.Println("hello", name)
	// 	}
	// 	return nil
	// }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
