// +build diagrams

package spherand

import (
	"image/color"
	"math/rand"
	"testing"

	"github.com/mmcloughlin/globe"
)

type GeographicalGenerator func() (lat, lng float64)

func Wrong() (lat, lng float64) {
	lat = rand.Float64()*180 - 90
	lng = rand.Float64()*360 - 180
	return
}

func Diagram(gen GeographicalGenerator, n int) *globe.Globe {
	radius := 0.04
	c := color.NRGBA{0, 0, 0, 48}

	g := globe.New()
	g.DrawGraticule(10.0)
	for i := 0; i < n; i++ {
		lat, lng := gen()
		g.DrawDot(lat, lng, radius, globe.Color(c))
	}
	g.CenterOn(75, 0)
	return g
}

func TestDiagrams(t *testing.T) {
	n := 100000
	side := 400
	diagrams := []struct {
		Generator GeographicalGenerator
		Filename  string
	}{
		{Wrong, "wrong.png"},
		{Geographical, "right.png"},
	}
	for _, diagram := range diagrams {
		g := Diagram(diagram.Generator, n)
		g.SavePNG(diagram.Filename, side)
	}
}
