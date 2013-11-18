package fibonacci

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	Convey("Fibonnacci", t, func() {
		fib := Fibonacci()
		sum := 0

		for i := <-fib; i < 4000000; i = <-fib {
			if i%2 == 0 {
				sum += i
			}
		}
		So(sum, ShouldEqual, 4613732)

	})
}
