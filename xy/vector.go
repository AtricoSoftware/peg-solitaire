package xy

type Vector XyPair

func (lhs Vector) String() string {
	return XyPair(lhs).StringAngle()
}

func (lhs Vector) Mult(rhs int) Vector {
	return Vector(XyPair(lhs).Mult(rhs))
}

