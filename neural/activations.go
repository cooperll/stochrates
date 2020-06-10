package neural

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Activation interface {
	Transform(*mat.VecDense)
}

type ReLu struct{}
type Sigmoid struct{}

func (r ReLu) Tranform(v *mat.VecDense) {
	for i := 0; i < v.Len(); i++ {
		v.SetVec(i, math.Max(float64(0), v.AtVec(i)))
	}
}
