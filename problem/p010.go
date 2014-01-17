// Find the sum of all the primes below two million.
package problem

import (
	"github.com/angeldm/euler/util/primegen"
)

func P010() uint64 {
	primeGen := primegen.New()
	sum := uint64(0)
	for i := primeGen.Next(); i < 2000000; i = primeGen.Next() {
		sum += i
	}
	return sum
}
