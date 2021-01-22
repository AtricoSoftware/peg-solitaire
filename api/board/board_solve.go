package board

import (
	"errors"
	"fmt"
	"os"

	"github.com/atrico-go/core"
)

func (b Board) Solve() ([]MoveList, error) {
	// Seed queue with initial position
	initial := newSolutionNode(b, make(MoveList, 0))
	tree := make(solutionTree, 1)
	tree[b.Id()] = initial
	return solve(tree, newPendingQueue(b.Id()))
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

func solve(tree solutionTree, queue pendingQueue) ([]MoveList, error) {
	solutions := make([]BoardId, 0)
	// TODO (this is for debugging)
	// solved := core.MaxInt
	// solved := 7 // 1m39 => 43s
	solved := 8 // 3m47
	for !queue.isEmpty() {
		nodeId := queue.pop()
		node := tree[nodeId]
		// Have we passed the quickest moves
		if node.moveCount > solved {
			break
		}
		// Check for already solved
		if node.remainingPegs == 1 {
			solutions = append(solutions, nodeId)
			solved = node.moveCount
		}
		// All moves
		board := NewBoardFromId(nodeId)
		for _, move := range board.GetMoves() {
			b2, err := board.MakeMove(move)
			if err != nil {
				core.FdisplayMultiline(os.Stderr, board)
				fmt.Fprintln(os.Stderr, move)
				panic("Invalid move")
			}
			newId := b2.Id()
			if _, exist := tree[newId]; !exist {
				tree[newId] = newSolutionNode(b2, append(node.moves, move))
				queue.push(newId)
			}
		}
	}
	if len(solutions) == 0 {
		return nil, errors.New("cannot be solved")
	}
	solutionMoves := make([]MoveList, len(solutions))
	for i, sln := range solutions {
		solutionMoves[i] = tree[sln].moves
	}
	return solutionMoves, nil
}

// ----------------------------------------------------------------------------------------------------------------------------
// Solution tree
// "Game" tree for calculating solution
// ----------------------------------------------------------------------------------------------------------------------------

type solutionTree map[BoardId]solutionNode

type solutionNode struct {
	moves         MoveList
	moveCount     int
	remainingPegs int
}

func newSolutionNode(board Board, moves MoveList) solutionNode {
	return solutionNode{moves: moves, moveCount: len(moves), remainingPegs: board.PegsRemaining()}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Pending queue
// Nodes pending expansion
// ----------------------------------------------------------------------------------------------------------------------------

type pendingQueue struct {
	queue []BoardId
	len   int
}

func newPendingQueue(ids ...BoardId) pendingQueue {
	return pendingQueue{queue: ids, len: len(ids)}
}

func (q pendingQueue) isEmpty() bool {
	return q.len == 0
}

func (q *pendingQueue) push(id BoardId) {
	q.queue = append(q.queue, id)
	q.len++
}

func (q *pendingQueue) pop() BoardId {
	if q.isEmpty() {
		panic("Queue is isEmpty")
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return val
}
