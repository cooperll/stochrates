package density_estimation

func DensityEstimation(data []float64, kernel *Pdf, h float64) *Pdf {
	return NewKDE(data, kernel, h)
}

func NewKDE(data []float64, kernel *Pdf, h float64) *Pdf {
	kde := func(x float64, suffStats []float64) float64 {
		n := len(data)
		scale := (1 / (float64(n) * float64(h)))
		var val float64
		for _, x_i := range suffStats {
			val += kernel.At((x - x_i) / h)
		}
		return scale * val
	}
	return &Pdf{density: kde, sufficientStats: data}
}
