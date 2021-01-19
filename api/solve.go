// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package api

import (
	"github.com/atrico-go/container"

  	"github.com/AtricoSoftware/peg-solitaire/settings"
)

func RegisterSolve(c container.Container) {
	c.Singleton(func(config settings.Settings) SolveApi {return solveApi{config: config}})
}

type solveApi struct {
config settings.Settings
}

// Solve board
func (svc solveApi) Run() error {
	// Implementation here!
	return nil
}
