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
	keys := make([]Position, len(b.holes))
	i := 0
	for pos := range b.holes {
		keys[i] = pos
		i++
	}
	sort.Slice(keys, func(i, j int) bool {return keys[i].Less(keys[j])})
	for _,pos := range keys {
		f(pos)
	}
}
