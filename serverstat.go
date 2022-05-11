package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vikpe/serverstat"
	"github.com/vikpe/serverstat/qserver"
	"github.com/vikpe/serverstat/qserver/mvdsv"
	"github.com/vikpe/serverstat/qserver/proxy"
	"github.com/vikpe/serverstat/qserver/qtv"
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

			fmt.Println(GenericServerToJson(genericServer))
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

func GenericServerToJson(genericServer qserver.GenericServer) string {
	serverToJson := func(v any) string {
		prefix := ""
		indent := "  "
		jsonBytes, _ := json.MarshalIndent(v, prefix, indent)
		return string(jsonBytes)
	}

	if genericServer.Version.IsMvdsv() {
		return serverToJson(mvdsv.Parse(genericServer))
	} else if genericServer.Version.IsQtv() {
		return serverToJson(qtv.Parse(genericServer))
	} else if genericServer.Version.IsProxy() {
		return serverToJson(proxy.Parse(genericServer))
	} else {
		return serverToJson(genericServer)
	}
}

func main() {
	os.Exit(run(os.Args))
}
