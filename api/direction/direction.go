package direction

import (
	"fmt"

	. "github.com/AtricoSoftware/peg-solitaire/xy"
)

type Direction Vector

var Up = Direction{Y: -1}
var Down = Direction{Y: +1}
var Left = Direction{X: -1}
var Right = Direction{X: +1}

var All = []Direction{Up, Down, Left, Right}

func (d Direction) String() string {
	return d.getInfo().string
}
func (d Direction) Opposite() Direction {
	return d.getInfo().opposite
}

func ForEachDirection(f func(Direction)) {
	for _, d := range All {
		f(d)
	}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
type directionInfo struct {
	string
	opposite Direction
}

var infoMap = map[Direction]directionInfo{
	Up:    {"Up", Down},
	Down:  {"Down", Up},
	Left:  {"Left", Right},
	Right: {"Right", Left},
}

func (d Direction) getInfo() directionInfo {
	if info, ok := infoMap[d]; ok {
		return info
	}
	panic(fmt.Sprintf(`Invalid direction: "%v"`, Vector(d)))
}
