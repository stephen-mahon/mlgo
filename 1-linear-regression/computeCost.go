package main

//  Compute cost for linear regression

import "gonum.org/v1/plot/plotter"

func ComputeCost(xys plotter.XYs, theta []float64) (cost, dm, dc float64) {
	// cost = 1/2N * sum((y-(m*x+c)^2)

	for _, xy := range xys {
		d := xy.Y - (xy.X*theta[0] + theta[1])
		cost += d * d
		dm += -xy.X * d
		dc += -d
	}
	n := float64(len(xys))
	return cost / (2 * n), 2 / n * dm, 2 / n * dc
}
