package test

import (
	"math"

	"gonum.org/v1/gonum/mat"

	util "stochrates/util"
)

var xBumps = []float64{0.1, 0.13, 0.15, 0.23, 0.25, 0.4,
	0.44, 0.65, 0.76, 0.78, 0.81}
var hBumps = []float64{4, 5, 3, 4, 5, 4.2,
	2.1, 4.3, 3.1, 2.1, 4.2}
var wBumps = []float64{0.005, 0.005, 0.006, 0.01, 0.01, 0.03,
	0.01, 0.01, 0.005, 0.008, 0.005}

var xBlocks = []float64{0.1, 0.13, 0.15, 0.23, 0.25, 0.4,
	0.44, 0.65, 0.76, 0.78, 0.81}
var hBlocks = []float64{4, -5, 3, -4, 5, -4.2,
	2.1, 4.3, -3.1, 2.1, -4.2}

func Doppler(X *mat.VecDense) *mat.VecDense {
	fun := func(x float64) float64 {
		return 3.4603 * (math.Sqrt(x*(1-x)) *
			math.Sin(2.1*math.Pi/(x+0.05)))
	}
	return util.VecApply(X, fun)
}

func Heavisine(X *mat.VecDense) *mat.VecDense {
	fun := func(x float64) float64 {
		return 0.7961 * (math.Sin(4*math.Pi*x) -
			util.Sgn(x-0.3) - util.Sgn(0.72-x))
	}
	return util.VecApply(X, fun)
}

func Bumps(X *mat.VecDense) *mat.VecDense {
	fun := func(x float64) float64 {
		K := func(i float64) float64 {
			return math.Pow((1 + math.Abs(x)), -4)
		}
		length := len(xBumps)
		var sum float64
		for i := 0; i < length; i++ {
			sum += hBumps[i] * K((x-xBumps[i])/wBumps[i])
		}
		return 1.5802 * (math.Sin(4*math.Pi*x) -
			util.Sgn(x-0.3) - util.Sgn(0.72-x))
	}
	return util.VecApply(X, fun)
}

func Blocks(X *mat.VecDense) *mat.VecDense {
	fun := func(x float64) float64 {
		K := func(i float64) float64 {
			return (1 + util.Sgn(x)) / 2
		}
		length := len(xBlocks)
		var sum float64
		for i := 0; i < length; i++ {
			sum += hBlocks[i] * K(x-xBlocks[i])
		}
		return 1.5802 * (math.Sin(4*math.Pi*x) -
			util.Sgn(x-0.3) - util.Sgn(0.72-x))
	}
	return util.VecApply(X, fun)
}
