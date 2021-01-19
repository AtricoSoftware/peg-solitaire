// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package api

import (
	"github.com/atrico-go/container"
)

// Api command to run
type ApiCommand interface {
	Run() error
}


// Solve board
type SolveApi ApiCommand

// Register Api services
func RegisterApi(c container.Container) {
	RegisterSolve(c)
}