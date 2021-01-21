// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package api

import (
	"fmt"
	"os"
	"time"

	"github.com/atrico-go/container"
	"github.com/atrico-go/core"

	"github.com/AtricoSoftware/peg-solitaire/api/board"
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
	b := board.NewBoard()
	core.DisplayMultiline(b)
	start := time.Now()
	moves,err := b.Solve()
	end := time.Now()
	if err == nil {
		for iSet,moveSet := range moves {
			fmt.Printf("Solution: %d (%d moves)\n", iSet, len(moveSet))
			for i,move:= range moveSet {
				fmt.Printf("%d:%v\n", i+1, move)
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Printf("time: %v\n", end.Sub(start))
	return nil
}
