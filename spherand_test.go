package spherand

import (
	"math"
	"math/rand"
	"testing"
)

func TestUniformSamplerInterface(t *testing.T) {
	var _ UniformSampler = new(rand.Rand)
}

func AssertRange(t *testing.T, x, min, max float64) {
	if x < min || x > max {
		t.Errorf("value outside allowed range: got %f expected [%f, %f]", x, min, max)
	}
}

func TestSphericalRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		azimuth, polar := Spherical()
		AssertRange(t, azimuth, 0, 2*math.Pi)
		AssertRange(t, polar, 0, math.Pi)
	}
}

func TestGeographicalRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		lat, lng := Geographical()
		AssertRange(t, lat, -90, 90)
		AssertRange(t, lng, -180, 180)
	}
}
