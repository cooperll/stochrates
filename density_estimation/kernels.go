package density_estimation

func NewKernelByName(name string) *Pdf {
	if name == "rectangular" {
		return Uniform(-1, 1)
	} else if name == "gaussian" {
		return Normal(0, 1)
	} else if name == "triangular" {
		return Triangular(2)
	} else if name == "epanechnikov" {
		return Epanechnikov()
	} else {
		return nil
	}
}

func NewKernelByFunc(density func(float64, []float64) float64) *Pdf {
	return &Pdf{density: density}
}
