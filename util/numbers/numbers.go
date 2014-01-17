package numbers

func NewTriangleGen() chan int {
	ch := make(chan int, 1)
	n := 1
	go func() {
		for i := 2; ; i++ {
			ch <- n
			n += i
		}
	}()
	return ch
}
