package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// Machine Learning Online Class - Exercise 1: Linear Regression

var fileName = "ex1data1.txt"
var iterations int
var alpha float64

func main() {
	flag.IntVar(&iterations, "n", 1500, "number of iterations for linear regression")
	flag.Float64Var(&alpha, "a", 0.01, "the learning rate for gradient descent")
	flag.Parse()

	// ======================= Part 1: Plotting =======================
	_, _, err := readData(fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", fileName, err)
	}

}

func pause() {
	fmt.Print("Program paused. Press enter to continue.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func readData(path string) (xs, ys []float64, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xs = append(xs, x)
		ys = append(ys, y)

	}
	if err := s.Err(); err != nil {
		return nil, nil, fmt.Errorf("could not scan: %v", err)
	}
	return xs, ys, nil
}

/*

func plotem(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter : %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	p.Add(s)

	x, c := linearRegression(xys, alpha)

	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*x + c}, {20, 20*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}

	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}

func linearRegression(xys plotter.XYs, alpha float64) (m, c float64) {

	for i := 0; i < iterations; i++ {
		dm, dc := computeGradient(xys, m, c)
		m += -dm * alpha
		c += -dc * alpha
		fmt.Printf("grad(%.2f, %.2f) = (%.2f, %.2f)\n", m, c, dm, dc)
	}
	fmt.Printf("cost(%.2f, %.2f) = %.2f\n", m, c, computeCost(xys, m, c))
	return m, c
}

func computeGradient(xys plotter.XYs, m, c float64) (dm, dc float64) {
	// cost = 1/N * sum((y-(m*x+c)^2)

	// d cost / dm = d/dt(1/N * sum((y-(m*x+c)^2))
	// d cost / dm = 2/N * sum(-x(y-(m*x+c)))

	// d cost / dc = d/dc(1/N * sum((y-(m*x+c)^2))
	// d cost / dc = 2/N * sum(-(y-(m*x+c)))

	for _, xy := range xys {
		d := xy.Y - (xy.X*m + c)
		dm += -xy.X * d
		dc += -d
	}

	return 2 / float64(len(xys)) * dm, 2 / float64(len(xys)) * dc
}
*/
