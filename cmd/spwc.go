package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/antixcode6/spwc/pkg"
	"github.com/antixcode6/spwc/pkg/crypt"
	"github.com/urfave/cli/v2"
)

// CLI stuff
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Initializes the directory your passwords will be stored",
				Action: func(c *cli.Context) error {
					err := config.SetPath()
					if err != nil {
						return nil // kill it
					}
					return nil
				},
			},
			{
				Name:  "ls",
				Usage: "Lists passwords",
				Action: func(c *cli.Context) error {
					crypt.Ls()
					return nil
				},
			},
			{
				Name:  "insert",
				Usage: "Inserts password to spwc",
				Action: func(c *cli.Context) error {
					args := c.Args().Slice()
					err := crypt.Insert(args[0], args[1])
					if err != nil {
						fmt.Fprintf(os.Stderr, "Unable to insert password: %v", err)
						return nil
					}
					return nil
				},
			},
			{
				Name:  "rm",
				Usage: "Removes a password from spwc",
				Action: func(c *cli.Context) error {
					err := config.SetPath()
					if err != nil {
						return nil // kill it
					}
					return nil
				},
			},
			{
				Name:  "generate",
				Usage: "Generates a new password",
				Action: func(c *cli.Context) error {
					err := config.SetPath()
					if err != nil {
						return nil // kill it
					}
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
