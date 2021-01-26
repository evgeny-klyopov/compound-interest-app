package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

//go run *.go 2020-12-01 1 10 30000 30000
func main() {
	app := &cli.App{
		Name:  "Compound Interest App",
		Usage: "Compound Interest App",
		Action: func(c *cli.Context) error {
			calc := newCalculate()

			var err error
			if c.Args().First() != "" {
				err = calc.setInArguments(c.Args())
			} else {
				err = calc.dialog()
			}
			if err != nil {
				return err
			}

			calc.run()

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
