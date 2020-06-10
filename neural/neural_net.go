package neural

import "gonum.org/v1/gonum/mat"

type NeuralNet struct {
	layers []func(*mat.VecDense) *mat.VecDense
	input  *mat.VecDense
	output *mat.VecDense
}

func (n *NeuralNet) init(layers []func(*mat.VecDense) *mat.VecDense) {
	n.layers = layers
}

func (n *NeuralNet) run(input *mat.VecDense) {
	n.input = input
	prevOutput := input
	for _, layer := range n.layers {
		layerOutput := layer(prevOutput)
		prevOutput = layerOutput
	}
	n.output = prevOutput
}
