package stochrates

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Linspace(from float64, to float64, by float64) *mat.VecDense {
	if to < from {
		panic("Linspace: invalid interval")
	}
	size := int((to-from)/by) + 1
	pts := make([]float64, size)
	for i := range pts {
		pts[i] = from + float64(i)*by
	}

	return mat.NewVecDense(len(pts), pts)
}

func Sgn(x float64) float64 {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func VecApply(X *mat.VecDense, f func(float64) float64) *mat.VecDense {
	length := X.Len()
	result := mat.NewVecDense(length, nil)
	for i := 0; i < length; i++ {
		result.SetVec(i, f(X.AtVec(i)))
	}
	return result
}

func Min(v *mat.VecDense) float64 {
	data := v.RawVector().Data
	curr := math.MaxFloat64
	for _, d := range data {
		curr = math.Min(curr, d)
	}
	return curr
}

func Max(v *mat.VecDense) float64 {
	data := v.RawVector().Data
	curr := math.Inf(-1)
	for _, d := range data {
		curr = math.Max(curr, d)
	}
	return curr
}
