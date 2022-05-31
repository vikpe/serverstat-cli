package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vikpe/serverstat"
	"github.com/vikpe/serverstat/qserver/convert"
)

func run(args []string) int {
	cli.AppHelpTemplate = `{{.Name}} [{{.Version}}]
{{.Description}}

  Usage:   {{.UsageText}}
Example:   {{.Name}} qw.foppa.dk:27501
`

	app := &cli.App{
		Name:        "serverstat",
		Description: "Get info from QuakeWorld servers.",
		UsageText:   "serverstat <address>",
		Version:     "__VERSION__", // updated during build workflow
		Action: func(c *cli.Context) error {
			serverAddress := c.Args().First()
			server, err := serverstat.GetInfo(serverAddress)

			if err != nil {
				return err
			}

			fmt.Println(convert.ToJson(server))
			return nil
		},
	}

	if 1 == len(args) {
		args = append(args, "--help")
	}

	err := app.Run(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args))
}
