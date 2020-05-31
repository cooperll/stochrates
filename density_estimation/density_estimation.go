package density_estimation

// DensityEstimation returns an estimate of the PDF of data. At the moment,
// only Kernel Density Estimation is used. Future versions will
// potentially include parametric methods as well.
// data is the samples of the density, kernel is the desired kernel for
// density estimation (see "density_estimation/kernels.go"), and h is
// the bandwidth.
func DensityEstimation(data []float64, kernel *Pdf, h float64) *Pdf {
	return NewKDE(data, kernel, h)
}

// Explicitly estimates the PDF of data using the nonparametric, kernel density
// estimation method.
// data is the samples of the density, kernel is the desired kernel for
// density estimation (see "density_estimation/kernels.go"), and h is
// the bandwidth.
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
