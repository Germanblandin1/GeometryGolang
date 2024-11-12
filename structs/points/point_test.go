package points

import (
	"geometry/structs"
	"math"
	"testing"
)

func equal(f1, f2 float64) bool {
	return math.Abs(f1-f2) < structs.EPS
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestAdd(t *testing.T) {
	P1 := NewPoint2D(1, 2)
	P2 := NewPoint2D(2, 3)

	expected := NewPoint2D(3, 5)
	obtained := P1.Add(P2)

	if !(equal(expected.X, obtained.X) && equal(expected.Y, obtained.Y)) {
		t.Fatalf(`Expected = %v, obtained %v`, expected, obtained)
	}
}
