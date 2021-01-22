package unit_tests

import (
	"testing"

	"github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/AtricoSoftware/peg-solitaire/api/board"
	"github.com/AtricoSoftware/peg-solitaire/api/direction"
)

func Test_GetMoves_Initial(t *testing.T) {
	// Arrange
	b := board.NewStandardBoard()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(00, +2, direction.Up),
		board.NewMove(00, -2, direction.Down),
		board.NewMove(+2, 00, direction.Left),
		board.NewMove(-2, 00, direction.Right),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_JustOne(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(-1, -3, direction.Down),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_Double(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		AddPeg(-1, 00).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(-1, -3, direction.Down),
		board.NewMove(-1, -3, direction.Down, direction.Down),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_Double2(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		AddPeg(-1, 00).
		AddPeg(-3, -1).
		AddPeg(-2, -1).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(-1, -3, direction.Down),
		board.NewMove(-1, -3, direction.Down, direction.Down),
		board.NewMove(-3, -1, direction.Right),
		board.NewMove(-3, -1, direction.Right, direction.Down),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_MultipleFeedInFeedOut(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		// Run of 4
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		AddPeg(-1, 00).
		AddPeg(-1, +2).
		AddPeg(00, +3).
		// Single feed in at 1
		AddPeg(-3, -1).
		AddPeg(-2, -1).
		// Double feed in at -1
		AddPeg(-3, 00).
		AddPeg(-2, +1).
		// Double lead out from -1
		AddPeg(00, +1).
		AddPeg(+2, +1).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(-1, -3, direction.Down),
		board.NewMove(-1, -3, direction.Down, direction.Down),
		board.NewMove(-1, -3, direction.Down, direction.Down, direction.Down),
		board.NewMove(-1, -3, direction.Down, direction.Down, direction.Down, direction.Right),
		board.NewMove(-1, -3, direction.Down, direction.Down, direction.Left),
		board.NewMove(-1, -3, direction.Down, direction.Down, direction.Right),
		board.NewMove(-1, -3, direction.Down, direction.Down, direction.Right, direction.Right),
		board.NewMove(-3, -1, direction.Down),
		board.NewMove(-3, -1, direction.Down, direction.Right),
		board.NewMove(-3, -1, direction.Down, direction.Right, direction.Up),
		board.NewMove(-3, -1, direction.Down, direction.Right, direction.Up, direction.Left),
		board.NewMove(-3, -1, direction.Down, direction.Right, direction.Down),
		board.NewMove(-3, -1, direction.Down, direction.Right, direction.Down, direction.Right),
		board.NewMove(-3, -1, direction.Down, direction.Right, direction.Right),
		board.NewMove(-3, -1, direction.Down, direction.Right, direction.Right, direction.Right),
		board.NewMove(-3, -1, direction.Right),
		board.NewMove(-3, -1, direction.Right, direction.Down),
		board.NewMove(-3, -1, direction.Right, direction.Down, direction.Down),
		board.NewMove(-3, -1, direction.Right, direction.Down, direction.Down, direction.Right),
		board.NewMove(-3, -1, direction.Right, direction.Down, direction.Left),
		board.NewMove(-3, -1, direction.Right, direction.Down, direction.Left, direction.Up),
		board.NewMove(-3, -1, direction.Right, direction.Down, direction.Right),
		board.NewMove(-3, -1, direction.Right, direction.Down, direction.Right, direction.Right),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}
