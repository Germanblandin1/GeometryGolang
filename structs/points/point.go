package points

import (
	"geometry/structs"
	"math"
)

type Point2D struct {
	X float64
	Y float64
}

func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{X: x, Y: y}
}

//Sum
func (p1 *Point2D) Add(p2 *Point2D) *Point2D {
	return &Point2D{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

//Vector X Scalar
func (p1 *Point2D) VXS(scalar float64) *Point2D {
	return &Point2D{
		X: scalar * p1.X,
		Y: scalar * p1.Y,
	}
}

func (p1 *Point2D) Sub(p2 *Point2D) *Point2D {
	return p1.Add(p2.VXS(-1))
}

//dotProduct
func (p1 *Point2D) DotProduct(p2 *Point2D) float64 {
	return p1.X*p2.X + p2.Y*p2.Y
}

func (p1 *Point2D) Equal(p2 *Point2D) bool {
	return math.Abs(p1.X-p2.X) < structs.EPS && math.Abs(p1.Y-p2.Y) < structs.EPS
}
