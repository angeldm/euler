package problem

import (
	"github.com/angeldm/euler/util/strings"
)

// Find the largest palindrome made from the product of two 3-digit numbers..
func P004() int {
	for i := 997; i > 99; i-- {
		pal := strings.MakePalindromo(i)
		for j := 999; j > 99; j-- {
			if pal/j > 999 || j*j < pal {
				break
			}
			if pal%j == 0 {
				return pal
			}
		}
	}
	return 0
}
