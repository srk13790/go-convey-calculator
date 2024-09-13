package calculator

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {

	convey.Convey("Given Two Integers", t, func() {
		num1 := 0
		num2 := 5

		convey.Convey("Adding Two integers", func(){
			result := Add(num1, num2)

			convey.Convey("Result should be sum of both integers", func ()  {
				convey.So(result, convey.ShouldEqual, 5)
			})
		})
	})
}

func TestSub(t *testing.T) {
	convey.Convey("Given two intergers", t, func ()  {
		num1 := 4
		num2 := 5
		convey.Convey("Substracting two numbers", func ()  {
			result := Sub(num1, num2)

			convey.Convey("Result should be less than zero", func ()  {
				convey.So(result, convey.ShouldBeLessThan, 0)
			})
		})
	})
}

