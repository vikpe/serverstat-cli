package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vikpe/serverstat"
	"github.com/vikpe/serverstat/qserver"
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
			genericServer, err := serverstat.GetInfo(serverAddress)

			if err != nil {
				return err
			}

			fmt.Println(genericServerToJson(genericServer))
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

func genericServerToJson(genericServer qserver.GenericServer) string {
	serverToJson := func(v any) string {
		prefix := ""
		indent := "  "
		jsonBytes, _ := json.MarshalIndent(v, prefix, indent)
		return string(jsonBytes)
	}

	if genericServer.Version.IsMvdsv() {
		return serverToJson(convert.ToMvdsvExport(genericServer))
	} else if genericServer.Version.IsQtv() {
		return serverToJson(convert.ToQtvExport(genericServer))
	} else if genericServer.Version.IsQwfwd() {
		return serverToJson(convert.ToQwfwdExport(genericServer))
	} else {
		return serverToJson(genericServer)
	}
}

func main() {
	os.Exit(run(os.Args))
}
