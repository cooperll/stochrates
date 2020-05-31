# Overview

Stochrates is an ML/DL library for Golang, currently in very early stages of development. The goal is to provide excellent statistical analysis and machine learning capabilities for Golang, choosing to not just be limited to deep learning, but to embrace both the computer scientists' and statisticians' points of view.

To get involved, just open a ticket in "Issues", and I can send you an introductory email and give you more info to help you contribute effectively!

## Current work:

For those interested in contributing, the next set of features stochrates is looking to add is:

[x] nonparametric probability density estimation

[ ] local polynomial regression/Nadaraya-Watson estimator

[ ] artificial neural networks

[ ] PCA

[ ] bootstrapping

[ ] Wavelets/Curvelets/Shearlets and respective regression methods


## Installation

stochrates works like any other Go package. To fetch the package, just run:

 ` go get github.com/cooperll/stochrates `


## Testing

Currently, ginkgo and gomega are the libraries used for testing, on top of the built-in Golang testing functionality. Tests can be run from the stochrates base directory with the following command:

 ` go test -v ./... `
 
 ## Documentation
 
At least for the beginning, stochrates documentation can be found on pkg.go.dev: https://pkg.go.dev/mod/github.com/cooperll/stochrates
