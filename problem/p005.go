package problem

import (
	"github.com/angeldm/euler/util/factor"
	"github.com/angeldm/euler/util/math"
)

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func P005() int {
	pows := make(map[uint64]uint64)
	for i := 2; i <= 20; i++ {
		factors := factor.FactorMap(uint64(i))
		for k, v := range factors {
			if val, ok := pows[k]; ok {
				if val < v {
					pows[k] = v
				}
			} else {
				pows[k] = v
			}
		}
	}

	result := uint64(1)
	for k, v := range pows {
		result *= math.PowUint64(k, v)
	}

	return int(result)
}
