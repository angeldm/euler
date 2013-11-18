package strings

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	Convey("Strings", t, func() {
		Convey("Reverse", func() {
			s := "0123456789"
			So(Reverse(s), ShouldEqual, "9876543210")
		})

		Convey("IsPalindromo", func() {
			s := "125521"
			So(IsPalindromo(s), ShouldEqual, true)
		})

		Convey("MakePalindromo", func() {
			s := 125
			So(MakePalindromo(s), ShouldEqual, 125521)
		})
	})
}
