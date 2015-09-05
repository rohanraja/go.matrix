package matrix

import (
	"math"
	"testing"

	"github.com/fatih/color"
)

func TestApplyFunc(t *testing.T) {
	A := Normals(3, 3)
	B := A.ApplyFunc(func(f float64) float64 {
		return f * 0
	})

	if ApproxEquals(Zeros(3, 3), B, 0.0001) == false {
		t.Fail()
	}
}
func TestTanhFunc(t *testing.T) {

	A := Normals(3, 3)
	B := A.Tanh()

	if B.Get(1, 1) != math.Tanh(A.Get(1, 1)) {
		t.Fail()
	}
}

func TestSigmoidFunc(t *testing.T) {

	A := Normals(3, 3)
	B := A.Sigmoid()

	color.Green("%v \n%v", A, B)
}
