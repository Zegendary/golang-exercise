package triangle

import "math"

type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
	} else if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		k = NaT
	} else if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		k = NaT
	} else if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a+b < c || a+c < b || b+c < a {
		k = NaT
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
