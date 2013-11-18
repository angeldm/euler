package factor

import (
	"github.com/angeldm/euler/util/math"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	Convey("Factor", t, func() {
		pows := make(map[uint64]uint64)
		for i := 2; i <= 20; i++ {
			factors := FactorMap(uint64(i))
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
		So(result, ShouldEqual, 232792560)

	})
}
