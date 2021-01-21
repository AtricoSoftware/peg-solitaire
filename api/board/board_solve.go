package board

import (
	"errors"

	"github.com/atrico-go/core"
)

func (b Board) Solve() ([]MoveList, error) {
	// Seed queue with initial position
	return solve(newNodeQueue(newSolutionNode(b, make(MoveList, 0))))
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

func solve(queue nodeQueue) ([]MoveList, error) {
	solutions := make([]MoveList, 0)
	solved := core.MaxInt
	for !queue.isEmpty() {
		node := queue.pop()
		// Have we passed the quickest moves
		if node.moveCount > solved {
			break
		}
		// Check for already solved
		if node.remainingPegs == 1 {
			solutions = append(solutions, node.moves)
			solved = node.moveCount
		}
		// All moves
		for _, move := range node.board.GetMoves() {
			queue.push(node.AddMove(move))
		}
	}
	var err error
	if len(solutions) == 0 {
		err = errors.New("cannot be solved")
	}
	return solutions, err
}

// ----------------------------------------------------------------------------------------------------------------------------
// Solution node
// ----------------------------------------------------------------------------------------------------------------------------

type solutionNode struct {
	board         Board
	moves         MoveList
	moveCount     int
	remainingPegs int
}

func newSolutionNode(board Board, moves MoveList) solutionNode {
	return solutionNode{board: board, moves: moves, moveCount: len(moves), remainingPegs: board.PegsRemaining()}
}

func (node solutionNode) AddMove(move Move) solutionNode {
	b2, err := node.board.MakeMove(move)
	if err != nil {
		panic("Invalid move")
	}
	return newSolutionNode(b2, append(node.moves, move))
}

// ----------------------------------------------------------------------------------------------------------------------------
// Node queue
// ----------------------------------------------------------------------------------------------------------------------------

func newNodeQueue(nodes ...solutionNode) nodeQueue {
	return nodeQueue{queue: nodes, len: len(nodes)}
}

type nodeQueue struct {
	queue []solutionNode
	len   int
}

func (q nodeQueue) isEmpty() bool {
	return q.len == 0
}

func (q *nodeQueue) push(node solutionNode) {
	q.queue = append(q.queue, node)
	q.len++
}

func (q *nodeQueue) pop() solutionNode {
	if q.isEmpty() {
		panic("Queue is isEmpty")
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return val
}
