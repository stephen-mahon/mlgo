package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Machine Learning Online Class - Exercise 1: Linear Regression

// x refers to the population size in 10,000s
// y refers to the profit in $10,000s

func main() {
	// Part 1: Basic Function
	fmt.Printf("Running warmUpExercise ... \n")
	fmt.Printf("5x5 Identity Matrix: \n")
	warmUpExercise()

	fmt.Printf("Program paused. Press enter to continue.\n")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Part 2: Plotting
	fmt.Printf("Plotting Data ...\n")
	graph()
}

func warmUpExercise() {
	a := mat.NewDiagDense(5, []float64{1, 1, 1, 1, 1})
	fmt.Println(a)
}

func graph() {
	rand.Seed(int64(0))

	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
