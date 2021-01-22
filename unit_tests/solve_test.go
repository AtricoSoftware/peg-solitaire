package unit_tests

import (
	"testing"

	"github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/AtricoSoftware/peg-solitaire/api/board"
	"github.com/AtricoSoftware/peg-solitaire/api/direction"
)

func Test_Solve_AlreadyFinished(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
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
		AddPeg(-1, -3).
		AddPeg(+1, +3).
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
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(-1, -3, direction.Down)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

func Test_Solve_TwoMoves(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(00, -1).
		AddPeg(+1, -1).
		AddPeg(+3, -1).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(00, -1, direction.Right), board.NewMove(+3, -1, direction.Left)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

func Test_Solve_MultipleSolutions(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(00, -1).
		AddPeg(+1, -1).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(00, -1, direction.Right)},
		{board.NewMove(+1 , -1, direction.Left)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

func Test_Solve_GetShortestOnly(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		AddPeg(-1, 00).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
		{board.NewMove(-1, -3, direction.Down, direction.Down)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}

// Multiple paths to same position
func Test_Solve_MergeMovesList(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(+1, -3).
		AddPeg(-1, -2).
		AddPeg(+1, -2).
		AddPeg(-2, -1).
		Build()
	core.DisplayMultiline(b)

	// Act
	moves, err := b.Solve()

	// Assert
	Assert(t).That(err, is.Nil, "Error")
	expected := []board.MoveList{
//		{board.NewMove(-1, -3, direction.Down, direction.Down)},
	}
	Assert(t).That(moves, is.EquivalentTo(expected), "moves")
}
