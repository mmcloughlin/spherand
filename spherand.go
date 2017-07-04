// Package spherand generates random points uniformly on a sphere.
package spherand

import (
	"math"
	"math/rand"
)

// UniformSampler generates samples from the uniform distribution on [0,1).
type UniformSampler interface {
	Float64() float64
}

// UniformSamplerFunc can satisfy the UniformSampler interface with a plain
// function.
type UniformSamplerFunc func() float64

// Float64 calls f.
func (f UniformSamplerFunc) Float64() float64 {
	return f()
}

// Generator can generate spherical points based on a configured source of
// uniform random values.
type Generator struct {
	sampler UniformSampler
}

// NewGenerator builds a Generator backed by the given uniform sampler. Note
// that the standard library rand.Rand can be used as a UniformSampler.
func NewGenerator(sampler UniformSampler) Generator {
	return Generator{
		sampler: sampler,
	}
}

// global is the internal generator used by package-level functions. It is
// backed by the global Float64 generator.
var global = NewGenerator(UniformSamplerFunc(rand.Float64))

// Spherical generates a random point on the unit sphere in spherical
// coordinates. This returns an azimuthal angle in radians between 0 and 2pi and
// a polar angle between 0 and pi.
func (g Generator) Spherical() (azimuth, polar float64) {
	azimuth = 2 * math.Pi * g.sampler.Float64()
	polar = math.Acos(2*g.sampler.Float64() - 1)
	return
}

// Spherical generates a random point on the unit sphere in spherical
// coordinates. This returns an azimuthal angle in radians between 0 and 2pi and
// a polar angle between 0 and pi.
// This package level function uses the default random source.
func Spherical() (azimuth, polar float64) {
	return global.Spherical()
}

// Geographical returns a random geographical point as latitude in degrees
// between -90 and 90, and longitude between -180 and 180.
func (g Generator) Geographical() (lat, lng float64) {
	azimuth, polar := g.Spherical()
	lat = 90 - 180.0*polar/math.Pi
	lng = 180.0*azimuth/math.Pi - 180
	return
}

// Geographical returns a random geographical point as latitude in degrees
// between -90 and 90, and longitude between -180 and 180.
// This package level function uses the default random source.
func Geographical() (lat, lng float64) {
	return global.Geographical()
}
