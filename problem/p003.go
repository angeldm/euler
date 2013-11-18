// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// What is the largest prime factor of the number 600851475143 ?
package problem

import (
	"github.com/angeldm/euler.go/util/primegen"
	"math"
)

func p003() int {

	pg := primegen.New()
	var max uint64
	for i := pg.Next(); i < uint64(math.Sqrt(float64(600851475143))); i = pg.Next() {
		if 600851475143%i == 0 {
			max = i
		}
	}

	return int(max)
}
