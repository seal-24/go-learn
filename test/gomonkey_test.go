package test

import (
	"fmt"
	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func add(i int32, j int32)int32{
	return i+j
}

// 函数太短会被内联，导致gomokey失效
// go test -gcflags='-N -l' gomonkey_test.go
func TestMonkeySum(t *testing.T) {

	// fun mock 例子
	patch1 := gomonkey.ApplyFunc(add,func(i int32, j int32)int32{
		// 和+2
		return i+j+2
	})
	defer  patch1.Reset()
	// 一定记得取消函数指针替换
	Convey("Test", t, func() {
		sum := add(1, 2) //sum=5
		fmt.Println(sum)
		So(sum == 5, ShouldBeTrue)
	})


}

func TestMonkeyFake(t *testing.T) {
	outputExpect := "xxx-vethName100-yyy"
	Convey("one func for succ", t, func() {
		patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
			return outputExpect, nil
		})
		defer patches.Reset()
		output, err := fake.Exec("", "")
		So(err, ShouldEqual, nil)
		So(output, ShouldEqual, outputExpect)
	})

}
