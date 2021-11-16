package main

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

/*
	[X] Plot the data
	[X] Set the y-axis label to "Profit in $10,000s"
	[X] Set the x-axis label to "Population of City in 10,000s"
*/

func PlotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()
	p.Y.Label.Text = "Profit in $10,000s"
	p.X.Label.Text = "Population of City in 10,000s"

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter : %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

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
