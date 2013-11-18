package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	Convey("Math", t, func() {
		Convey("PowUint64", func() {
			So(PowUint64(9999, 3), ShouldEqual, 999700029999)
		})

		Convey("PowInt", func() {
			So(PowInt(99, 2), ShouldEqual, 9801)
		})
	})
}
