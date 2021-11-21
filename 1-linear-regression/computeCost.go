package main

//  Compute cost for linear regression

func ComputeCost(xs, ys, theta []float64) (cost, dm, dc float64) {
	// cost = 1/2N * sum((y-(m*x+c)^2)

	for i := range xs {
		d := ys[i] - (xs[i]*theta[0] + theta[1])
		cost += d * d
		dm += -xs[i] * d
		dc += -d
	}
	n := float64(len(xs))
	return cost / (2 * n), 2 / n * dm, 2 / n * dc
}
