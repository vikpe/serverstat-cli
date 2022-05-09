package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/udphelper"
)

type AppRunner struct {
	app string
}

func (ar AppRunner) run(appArgs string) string {
	argsStr := strings.TrimSpace(fmt.Sprintf("%s %s", ar.app, appArgs))
	argsArr := strings.Split(argsStr, " ")

	captureOutput := func(f func()) string {
		rescueStderr := os.Stderr
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		os.Stderr = w

		f()

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = rescueStderr
		os.Stdout = rescueStdout

		return string(out)
	}

	return captureOutput(func() { run(argsArr) })
}

var app = AppRunner{
	app: "masterstat",
}

func TestHelp(t *testing.T) {
	helpText := `serverstat [__VERSION__]
Get info from QuakeWorld servers.

  Usage:   serverstat <address>
Example:   serverstat qw.foppa.dk:27501
`

	t.Run("No args", func(t *testing.T) {
		assert.Equal(t, app.run(""), helpText)
	})

	t.Run("Help", func(t *testing.T) {
		assert.Equal(t, app.run("help"), helpText)
		assert.Equal(t, app.run("--help"), helpText)
		assert.Equal(t, app.run("-h"), helpText)
	})
}

func TestGetInfo(t *testing.T) {
	t.Run("UDP request error", func(t *testing.T) {
		output := app.run("foo:666")
		assert.Contains(t, output, "ERROR:")
	})

	t.Run("Success", func(t *testing.T) {
		go func() {
			responseHeader := string([]byte{0xff, 0xff, 0xff, 0xff, 'n', '\\'})
			responseBody := `\maxfps\77\pm_ktjump\1\*version\MVDSV 0.35-dev
65 -9999 16 -666 "\s\[ServeMe]" "" 12 11 "lqwc" ""`
			udphelper.New(":8000").Respond([]byte((responseHeader + responseBody)))
		}()
		time.Sleep(10 * time.Millisecond)

		outputAsJson := app.run(":8000")
		expectedOutput := jsonFileToString("./test_files/expected_output.json")

		assert.Equal(t, expectedOutput, outputAsJson)
	})
}

func jsonFileToString(filePath string) string {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return string(byteValue)
}
