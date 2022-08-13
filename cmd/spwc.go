package main

import (
	"log"
	"os"

	config "github.com/antixcode6/spwc/pkg"
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
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
