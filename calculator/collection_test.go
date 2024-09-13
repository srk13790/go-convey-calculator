package calculator

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestTextContains(t *testing.T) {
	convey.Convey("Passing a string", t, func ()  {
		str := "Sagar"

		convey.Convey("String will get converted to map", func ()  {
			result := TextContains(str)

			resultMap := TextContainsMaps(str)

			convey.Convey("Result should contain `a` and `g`", func ()  {
				convey.So(result, convey.ShouldContain, "a")
				convey.So(result, convey.ShouldNotContain, "k")

				convey.So("S", convey.ShouldBeIn, result)

				convey.So(resultMap, convey.ShouldContainKey, "a")

				convey.So(resultMap, convey.ShouldNotContainKey, "k")
			})
		})
	})
}

func TestPanicFunction(t *testing.T) {
	convey.Convey("Checking the input return valid result or not", t, func ()  {
		
	})
}