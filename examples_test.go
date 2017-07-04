package spherand_test

import (
	"fmt"
	"math/rand"

	"github.com/mmcloughlin/spherand"
)

func ExampleGenerator_Geographical() {
	g := spherand.NewGenerator(rand.New(rand.NewSource(1)))
	fmt.Println(g.Geographical())
	// Output: 61.76543033984557 37.67770367266306
}
