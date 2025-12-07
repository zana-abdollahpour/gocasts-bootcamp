package triangle

import "math"

// Kind represents the type of triangle
type Kind int

const (
	NaT Kind = iota // Not a Triangle
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// KindFromSides determines the kind of triangle given three sides
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) ||
		a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
