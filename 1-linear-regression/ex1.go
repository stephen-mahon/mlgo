package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

// Machine Learning Online Class - Exercise 1: Linear Regression

// x refers to the population size in 10,000s
// y refers to the profit in $10,000s
var fileName = "ex1data1.txt"

func main() {
	xys, err := readData(fileName)
	if err != nil {
		log.Fatalf("could not read %v: %v", fileName, err)
	}
	_ = xys

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

type xy struct{ x, y float64 }

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys []xy
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
		}
		xys = append(xys, xy{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}

func plotData(path string, xys []xy) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()
	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}

	s, err := plotter.NewScatter(plotter.XYs{
		{0, 0}, {0.5, 0.5}, {1, 1},
	})
	if err != nil {
		return fmt.Errorf("could not create scatter : %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(s)

	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}
