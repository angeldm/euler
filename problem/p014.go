// Which starting number, under one million, produces the longest chain?
package problem

func sequence(n uint64) int {
	count := 1
	for n != 1 {
		n = collatz(n)
		count++

	}
	return count
}

func collatz(n uint64) uint64 {
	if n%2 == 0 {
		return n / 2
	} else {
		return (3 * n) + 1
	}
}

func P014() uint64 {
	max := 0
	num := uint64(0)
	for i := uint64(1); i < 1000000; i++ {
		n := sequence(i)
		if n > max {
			max = n
			num = i
		}
	}
	return num
}
