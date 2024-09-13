package calculator

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestStringTest(t *testing.T) {
	convey.Convey("Passing String to check its exitance", t, func ()  {
		str := "Sagar"
		result := StringTest(str)

		convey.Convey("Expecteing Hello in the result string", func ()  {
			convey.So(result, convey.ShouldStartWith, "Hello")
		})
	})
}