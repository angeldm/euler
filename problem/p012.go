// What is the value of the first triangle number to have over five hundred divisors?
package problem

import (
	"github.com/angeldm/euler/util/factor"
	"github.com/angeldm/euler/util/numbers"
)

func P012() int {
	ch := numbers.NewTriangleGen()
	for i := 0; ; i++ {
		n := <-ch
		if factor.NumDivisors(uint64(n)) > 500 {
			return n
		}
	}

	return 0
}
