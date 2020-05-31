package density_estimation

import "math"

type function interface {
	At(float64) float64
}

// At returns the value of the PDF at x
func (p *Pdf) At(x float64) float64 {
	return p.density(x, p.sufficientStats)
}

// Represents a general probability density function
type Pdf struct {
	density         func(float64, []float64) float64
	sufficientStats []float64
}

// The Gaussian/Normal pdf
func Normal(mu float64, sigmaSquared float64) *Pdf {
	sufficientStats := []float64{mu, sigmaSquared}
	normalPdf := func(x float64, sufficientStats []float64) float64 {
		mu := sufficientStats[0]
		sigmaSq := sufficientStats[1]
		return (1 / math.Sqrt(2*math.Pi*sigmaSq)) * math.Exp(-0.5*(math.Pow(x-mu, 2)/sigmaSq))
	}
	return &Pdf{density: normalPdf,
		sufficientStats: sufficientStats}

}

// Triangular pdf, where width is the total width of the distribution,
// and the mode is always at 0.
func Triangular(width float64) *Pdf {
	sufficientStats := []float64{width}
	density := func(x float64, sufficientStats []float64) float64 {
		if math.Abs(x) > sufficientStats[0]-1 {
			return 0
		}
		height := 1 / (0.5 * sufficientStats[0])
		return height - (height/(0.5*width))*math.Abs(x)
	}
	return &Pdf{density: density,
		sufficientStats: sufficientStats}
}

// The Uniform distribution, starting at a, ending at b
func Uniform(a float64, b float64) *Pdf {
	density := func(x float64, ss []float64) float64 {
		if x <= b && x >= a {
			return float64(1 / (b - a))
		}
		return float64(0)
	}
	return &Pdf{density: density,
		sufficientStats: []float64{a, b}}
}

// The Epanechnikov kernel/PDF. Primarily used for kernel density
// estimation, so no arguments are passed (at the moment).
func Epanechnikov() *Pdf {
	return &Pdf{density: func(x float64, ss []float64) float64 {
		if math.Abs(x) > 1 {
			return 0
		}
		return (float64(3) / 4) * (1 - math.Pow(x, 2))
	}}
}
