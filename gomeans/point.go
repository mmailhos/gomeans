package gomeans

import "math"

//Point struct is a simple coordinate
type Point struct {
	X float64
	Y float64
}

//Distance function calculates distance between two points in the cartesian plan
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}
