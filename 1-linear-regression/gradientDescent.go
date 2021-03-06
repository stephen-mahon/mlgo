package main

/*
GradientDescent erforms gradient descent to learn theta
	theta = GradientDescent(xys, theta, alpha) updates theta by taking gradient steps with learning rate alpha
*/

func GradientDescent(xs, ys []float64, alpha float64) (theta []float64, cost float64) {
	var dm, dc float64
	theta = []float64{0, 0}
	for i := 0; i < iterations; i++ {
		cost, dm, dc = ComputeCost(xs, ys, theta)
		theta[0] = -dm * alpha
		theta[1] = -dc * alpha
	}

	return theta, cost
}
