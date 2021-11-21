package main

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

type xyer struct {
	xs, ys []float64
}

func (x xyer) Len() int {
	return len(x.xs)
}

func (x xyer) XY(i int) (float64, float64) {
	return x.xs[i], x.ys[i]
}

func PlotData(path string, xs, ys []float64) error {
	p := plot.New()
	p.Y.Label.Text = "Profit in $10,000s"
	p.X.Label.Text = "Population of City in 10,000s"

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	s, err := plotter.NewScatter(xyer{xs, ys})
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
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
