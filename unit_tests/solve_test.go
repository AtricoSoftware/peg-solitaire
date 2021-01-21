package unit_tests

import (
	"testing"

	"github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/AtricoSoftware/peg-solitaire/api/board"
)

func Test_Solve_AlreadyFinished(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(2, 0).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{make([]board.Move, 0)}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

func Test_Solve_Unsolveable(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(2, 0).
		AddPeg(4, 6).
		Build()
	core.DisplayMultiline(b)

	// Act
	_, err := b.Solve()

	// Assert
	Assert(t).That(err, is.NotNil, "Error")
	Assert(t).That(err.Error(), is.EqualTo("cannot be solved"), "Error Msg")
}

func Test_Solve_OneMove(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(2, 0).
		AddPeg(2, 1).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(2, 0, board.Down)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

func Test_Solve_TwoMoves(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(3, 2).
		AddPeg(4, 2).
		AddPeg(6, 2).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(3, 2, board.Right), board.NewMove(6, 2, board.Left)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

func Test_Solve_MultipleSolutions(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(3, 2).
		AddPeg(4, 2).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(3, 2, board.Right)},
		{board.NewMove(4, 2, board.Left)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

func Test_Solve_GetShortestOnly(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(2, 0).
		AddPeg(2, 1).
		AddPeg(2, 3).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(2, 0, board.Down, board.Down)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}
