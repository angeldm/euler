// What is the 10 001st prime number?
package problem

import (
	"github.com/angeldm/euler/util/primegen"
)

func P007() uint64 {
	pg := primegen.New()
	for i := 0; i < 10000; i++ {
		pg.Next()
	}
	return pg.Next()
}
