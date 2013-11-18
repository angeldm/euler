package factors

import (
	"github.com/angeldm/euler.go/util/primegen"
)

func FactorsMap(n uint64) map[uint64]uint64 {
	factors := make(map[uint64]uint64, 100) //Hope 100 is enough
	pg := primegen.New()
	for i := pg.Next(); n != 1; i = pg.Next() {

		if n == 1 {
			break
		}
		for n%i == 0 {
			n = n / i
			factors[i]++
		}

	}
	return factors
}

func Factors(n uint64) []uint64 {
	factors := make([]uint64, 0, 100) //Hope 100 is enough
	pg := primegen.New()
	for i := pg.Next(); n != 1; i = pg.Next() {

		if n == 1 {
			break
		}
		for n%i == 0 {
			n = n / i
			factors = factors[0 : len(factors)+1]
			factors[len(factors)-1] = i
		}

	}
	return factors
}
