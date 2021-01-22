package board

import "sort"

//    - - -   + + +
//    3 2 1 0 1 2 3
// -3     @ @ @
// -2     @ @ @
// -1 @ @ @ @ @ @ @
//  0 @ @ @ o @ @ @
// +1 @ @ @ @ @ @ @
// +2     @ @ @
// +3     @ @ @

func (b Board) forEachPegPosition(f func(Position)) {
	b.forEachPegPositionImpl(f, false)
}
func (b Board) forEachPegPositionReverse(f func(Position)) {
	b.forEachPegPositionImpl(f, true)
}
func (b Board) forEachPegPositionImpl(f func(Position), reverse bool) {
	keys := make([]Position, len(b.holes))
	i := 0
	for pos := range b.holes {
		keys[i] = pos
		i++
	}
	less := func(i, j int) bool { return keys[i].Less(keys[j]) }
	if reverse {
		less = func(i, j int) bool { return !keys[i].Less(keys[j]) }
	}
	sort.Slice(keys, less)
	for _, pos := range keys {
		f(pos)
	}
}
