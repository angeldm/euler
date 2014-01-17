// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
package problem

import ()

func P009() int {

	for a := 3; a < 1000; a++ {
		for b := a; b+a < 1000; b++ {
			for c := b; c < 1000; c++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					return (a * b * c)
				}
			}
		}
	}

	return 0
}
