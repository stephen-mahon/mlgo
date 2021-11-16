package main

//  Compute cost for linear regression

import "gonum.org/v1/plot/plotter"

func ComputeCost(xys plotter.XYs, theta []float64) (cost float64) {
	// cost = 1/2N * sum((y-(m*x+c)^2)

	for _, xy := range xys {
		d := xy.Y - (xy.X*theta[0] + theta[1])
		cost += d * d
	}

	return cost / (2 * float64(len(xys)))
}
