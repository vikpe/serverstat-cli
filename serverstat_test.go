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

	t.Run("MVDSV", func(t *testing.T) {
		go func() {
			responseHeader := string([]byte{0xff, 0xff, 0xff, 0xff, 'n', '\\'})
			responseBody := `\maxfps\77\pm_ktjump\1\*version\MVDSV 0.35-dev\maxclients\8\maxspectators\4
65 -9999 16 -666 "\s\[ServeMe]" "" 12 11 "lqwc" ""`
			udphelper.New(":8000").Respond([]byte((responseHeader + responseBody)))
		}()
		time.Sleep(10 * time.Millisecond)

		outputAsJson := app.run(":8000")
		expectedOutput := jsonFileToString("./test_files/mvdsv.json")

		assert.Equal(t, expectedOutput, outputAsJson)
	})

	t.Run("QTV", func(t *testing.T) {
		go func() {
			responseHeader := string([]byte{0xff, 0xff, 0xff, 0xff, 'n', '\\'})
			responseBody := `\*version\QTV 1.12-rc1\hostname\qw.foppa.dk - qtv\maxclients\100`
			udphelper.New(":8001").Respond([]byte((responseHeader + responseBody)))
		}()
		time.Sleep(10 * time.Millisecond)

		outputAsJson := app.run(":8001")
		expectedOutput := jsonFileToString("./test_files/qtv.json")

		assert.Equal(t, expectedOutput, outputAsJson)
	})

	t.Run("QWFWD", func(t *testing.T) {
		go func() {
			responseHeader := string([]byte{0xff, 0xff, 0xff, 0xff, 'n', '\\'})
			responseBody := `\*version\qwfwd 1.2\maxclients\128\hostname\qw.foppa.dk - qwfwd`
			udphelper.New(":8002").Respond([]byte((responseHeader + responseBody)))
		}()
		time.Sleep(10 * time.Millisecond)

		outputAsJson := app.run(":8002")
		expectedOutput := jsonFileToString("./test_files/qwfwd.json")

		assert.Equal(t, expectedOutput, outputAsJson)
	})

	t.Run("Unknown", func(t *testing.T) {
		go func() {
			responseHeader := string([]byte{0xff, 0xff, 0xff, 0xff, 'n', '\\'})
			responseBody := `\*version\foo v1.2\hostname\foo.bar`
			udphelper.New(":8003").Respond([]byte((responseHeader + responseBody)))
		}()
		time.Sleep(10 * time.Millisecond)

		outputAsJson := app.run(":8003")
		expectedOutput := jsonFileToString("./test_files/unknown.json")

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
