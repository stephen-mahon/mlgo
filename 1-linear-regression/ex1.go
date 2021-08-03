package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
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

	file, err := os.Open("ex1data1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var X, y []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		x_n, err := strconv.ParseFloat(s[0], 64)
		if err != nil {
			panic(err)
		}
		y_n, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			panic(err)
		}
		X = append(X, x_n)
		y = append(y, y_n)
	}
	xy := xyPoints(X, y)
	graph(xy)
}

func warmUpExercise() {
	a := mat.NewDiagDense(5, []float64{1, 1, 1, 1, 1})
	fmt.Println(a)
}

func graph(xy plotter.XYs) {
	p := plot.New()

	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "y"

	s, err := plotter.NewScatter(xy)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(s)
	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func xyPoints(x, y []float64) plotter.XYs {
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return pts
}
