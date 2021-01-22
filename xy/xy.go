package xy

import (
	"fmt"
)

type XyPair struct {
	X, Y int
}

func (lhs XyPair) String() string {
	return lhs.StringRound()
}
func (lhs XyPair) StringRound() string  {return lhs.StringOpenClose("(",")")}
func (lhs XyPair) StringSquare() string {return lhs.StringOpenClose("[","]")}
func (lhs XyPair) StringCurly() string {return lhs.StringOpenClose("{","}")}
func (lhs XyPair) StringAngle() string {return lhs.StringOpenClose("<",">")}

func (lhs XyPair) StringOpenClose(open, close string) string {return lhs.StringOpenCloseSep(open,close,",")}
func (lhs XyPair) StringOpenCloseSep(open, close, sep string) string {
	return fmt.Sprintf("%s%d%s%d%s", open,lhs.X,sep, lhs.Y,close)
}

func (lhs XyPair) Add(rhs XyPair) XyPair {
	return XyPair{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

func (lhs XyPair) Sub(rhs XyPair) XyPair {
	return XyPair{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

func (lhs XyPair) Mult(rhs int) XyPair {
	return XyPair{lhs.X * rhs, lhs.Y * rhs}
}
