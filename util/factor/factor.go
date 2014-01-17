package factor

import (
	"github.com/angeldm/euler/util/primegen"
	"math"
)

func FactorMap(n uint64) map[uint64]uint64 {
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

func Factor(n uint64) []uint64 {
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

func NumDivisors(n uint64) uint64 {
	//Buggy implementation, if n is a perfect square then the number of divisors will necessary be one too much
	divisors := uint64(1)
	for i := uint64(2); i <= uint64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors++
		}
	}
	divisors *= 2
	return divisors

}
