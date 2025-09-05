package geometry

import "math"

func Hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
