package spherand

import (
	"math/rand"
	"testing"
)

func TestUniformSamplerInterface(t *testing.T) {
	var _ UniformSampler = new(rand.Rand)
}
