package test

import (
	. "stochrates/density_estimation"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKDE(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("TEST", func() {
	var samples = []float64{-2, 0.008, 0.12, 0.84}
	var pts = []float64{-1.1, 0.003, 0.5}
	var kernelNames = []string{"rectangular", "triangular",
		"gaussian", "epanechnikov"}
	var setOfCorrectVals = [][]float64{
		{0, 0.5, 0.75},
		{0, 0.878, 0.28800000000000003},
		{0.06686823369616922, 0.44274362535259104, 0.43065436231920373},
		{0, 0.729429, 0.371904},
	}
	binwidth := 0.5

	for i, kernelName := range kernelNames {
		kernel := NewKernelByName(kernelName)
		kde := DensityEstimation(samples, kernel, binwidth)
		correctVals := setOfCorrectVals[i]
		assertKDEReturnedCorrectValue(kde, pts, correctVals, kernelName)
	}
})

func assertKDEReturnedCorrectValue(kde *Pdf, pts []float64, correct []float64,
	kernelName string) {

	Describe("Evaluating KDEs", func() {
		Context(kernelName+"KDE with 4 data points", func() {
			It("should have these values", func() {
				Expect(kde.At(pts[0])).To(Equal(correct[0]))
				Expect(kde.At(pts[1])).To(Equal(correct[1]))
				Expect(kde.At(pts[2])).To(Equal(correct[2]))
			})
		})
	})
}
