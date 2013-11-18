package math

func PowUint64(num uint64, p uint64) uint64 {
	r := uint64(1)
	for i := uint64(0); i < p; i++ {
		r *= num
	}
	return r
}

func PowInt(num int, p int) int {
	r := 1
	for i := 0; i < p; i++ {
		r *= num
	}
	return r
}
