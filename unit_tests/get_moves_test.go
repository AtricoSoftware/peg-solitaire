package unit_tests

import (
	"testing"

	"github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/AtricoSoftware/peg-solitaire/api/board"
)

func Test_GetMoves_Initial(t *testing.T) {
	// Arrange
	b := board.NewBoard()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(3, 5, board.Up),
		board.NewMove(3, 1, board.Down),
		board.NewMove(5, 3, board.Left),
		board.NewMove(1, 3, board.Right),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_JustOne(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(2,0).
		AddPeg(2,1).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(2, 0, board.Down),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_Double(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(2,0).
		AddPeg(2,1).
		AddPeg(2,3).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(2, 0, board.Down),
		board.NewMove(2, 0, board.Down, board.Down),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_Double2(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(2,0).
		AddPeg(2,1).
		AddPeg(2,3).
		AddPeg(0,2).
		AddPeg(1,2).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(2, 0, board.Down),
		board.NewMove(2, 0, board.Down, board.Down),
		board.NewMove(0, 2, board.Right),
		board.NewMove(0, 2, board.Right, board.Down),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}

func Test_GetMoves_MultipleFeedInFeedOut(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		// Run of 4
		AddPeg(2,0).
		AddPeg(2,1).
		AddPeg(2,3).
		AddPeg(2,5).
		AddPeg(3,6).
		// Single feed in at 1
		AddPeg(0,2).
		AddPeg(1,2).
		// Double feed in at 2
		AddPeg(0,3).
		AddPeg(1,4).
		// Double lead out from 2
		AddPeg(3,4).
		AddPeg(5,4).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves := b.GetMoves()

	// Assert
	expected := []board.Move{
		board.NewMove(2, 0, board.Down),
		board.NewMove(2, 0, board.Down, board.Down),
		board.NewMove(2, 0, board.Down, board.Down, board.Down),
		board.NewMove(2, 0, board.Down, board.Down, board.Down, board.Right),
		board.NewMove(2, 0, board.Down, board.Down, board.Left),
		board.NewMove(2, 0, board.Down, board.Down, board.Right),
		board.NewMove(2, 0, board.Down, board.Down, board.Right, board.Right),
		board.NewMove(0, 2, board.Down),
		board.NewMove(0, 2, board.Down, board.Right),
		board.NewMove(0, 2, board.Down, board.Right, board.Up),
		board.NewMove(0, 2, board.Down, board.Right, board.Up, board.Left),
		board.NewMove(0, 2, board.Down, board.Right, board.Down),
		board.NewMove(0, 2, board.Down, board.Right, board.Down, board.Right),
		board.NewMove(0, 2, board.Down, board.Right, board.Right),
		board.NewMove(0, 2, board.Down, board.Right, board.Right, board.Right),
		board.NewMove(0, 2, board.Right),
		board.NewMove(0, 2, board.Right, board.Down),
		board.NewMove(0, 2, board.Right, board.Down, board.Down),
		board.NewMove(0, 2, board.Right, board.Down, board.Down, board.Right),
		board.NewMove(0, 2, board.Right, board.Down, board.Left),
		board.NewMove(0, 2, board.Right, board.Down, board.Left, board.Up),
		board.NewMove(0, 2, board.Right, board.Down, board.Right),
		board.NewMove(0, 2, board.Right, board.Down, board.Right, board.Right),
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "Correct moves")
}
