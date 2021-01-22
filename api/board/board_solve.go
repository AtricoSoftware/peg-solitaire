package board

import (
	"errors"
	"fmt"
	"os"

	"github.com/atrico-go/core"
)

func (b Board) Solve() ([]MoveList, error) {
	// Seed queue with initial position
	initial := newRootSolutionNode(b)
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
		// Have we passed the quickest Moves
		if node.MoveCount > solved {
			break
		}
		// Check for already solved
		if node.RemainingPegs == 1 {
			solutions = append(solutions, nodeId)
			solved = node.MoveCount
		} else {
			// All Moves
			board := NewBoardFromId(nodeId)
			moves := board.GetMoves()
			node.Links = make([]moveLink, len(moves))
			for i, move := range moves {
				b2, err := board.MakeMove(move)
				if err != nil {
					core.FdisplayMultiline(os.Stderr, board)
					fmt.Fprintln(os.Stderr, move)
					panic("Invalid move")
				}
				newId := b2.Id()
				node.Links[i] = moveLink{move, newId}
				if _, exist := tree[newId]; !exist {
					tree[newId] = newSolutionNode(b2, nodeId, append(node.Moves, move))
					queue.push(newId)
				}
			}
			// Update node (with moves)
			tree[nodeId] = node
		}
	}
	if len(solutions) == 0 {
		return nil, errors.New("cannot be solved")
	}
	solutionMoves := make([]MoveList, 0)
	for _, sln := range solutions {
		solutionMoves = append(solutionMoves, tree.resolveMoves(sln)...)
	}
	return solutionMoves, nil
}

// ----------------------------------------------------------------------------------------------------------------------------
// Solution tree
// "Game" tree for calculating solution
// ----------------------------------------------------------------------------------------------------------------------------

type solutionTree map[BoardId]solutionNode

type solutionNode struct {
	Parents       []BoardId
	Moves         MoveList
	MoveCount     int
	RemainingPegs int
	Links         []moveLink
}

type moveLink struct {
	Move
	Target BoardId
}

func newRootSolutionNode(board Board) solutionNode {
	return solutionNode{Parents: make([]BoardId, 0), Moves: make(MoveList, 0), MoveCount: 0, RemainingPegs: board.PegsRemaining()}
}

func newSolutionNode(board Board, parent BoardId, moves MoveList) solutionNode {
	return solutionNode{Parents: []BoardId{parent}, Moves: moves, MoveCount: len(moves), RemainingPegs: board.PegsRemaining()}
}

func (t solutionTree) resolveMoves(id BoardId) []MoveList {
	return t.resolveMovesImpl(id, make(MoveList, 0))
}

// TODO - optimise this(?)
func (t solutionTree) resolveMovesImpl(id BoardId, movesSoFar MoveList) []MoveList {
	node := t[id]
	parentCount := len(node.Parents)
	// End recursion
	if parentCount == 0 {
		return []MoveList{movesSoFar}
	}
	moves := make([]MoveList, 0)
	for _, parent := range node.Parents {
		// Find move in parent
		nextMove := t.findMoveInParent(parent, id)
		moves = append(moves, t.resolveMovesImpl(parent, insertIntoMoveList(nextMove, movesSoFar))...)
	}
	return moves
}

func (t solutionTree) findMoveInParent(parentId, nodeId BoardId) Move {
	parent := t[parentId]
	for _, link := range parent.Links {
		if link.Target == nodeId {
			return link.Move
		}
	}
	panic("Invalid parent relationship")
}

func insertIntoMoveList(move Move, list MoveList) MoveList {
	newList := make(MoveList, len(list)+1)
	newList[0] = move
	copy(newList[1:], list)
	return newList
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
