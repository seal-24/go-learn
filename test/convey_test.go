package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Sum(a,b int) int {
	ret := a+ b
	return ret
}

func TestSum(t *testing.T)  {
	val := Sum(1, 2)
	Convey("Test", t, func() {
		So(val == 3, ShouldBeTrue)
	})

}


