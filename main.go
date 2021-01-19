// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package main

import (
	"fmt"
	"os"

	"github.com/atrico-go/container"

	"github.com/AtricoSoftware/peg-solitaire/api"
	"github.com/AtricoSoftware/peg-solitaire/cmd"
	"github.com/AtricoSoftware/peg-solitaire/settings"
)

func main() {
	c := register()
	cmd := cmd.CreateCommands(c)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func register() container.Container {
	c := container.NewContainer()
	settings.RegisterSettings(c)
	api.RegisterApi(c)
	return c
}