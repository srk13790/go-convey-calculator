package calculator

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestPanicTests(t *testing.T) {
	convey.Convey("Checking function is panicing or not", t, func ()  {
		convey.Convey("Passing two integers", func ()  {
			convey.So(func() { PanicTests(1, 2) }, convey.ShouldPanic)
		})
	})
}

func TestDivide(t *testing.T) {
    convey.Convey("Given two integers", t, func() {

        convey.Convey("When dividing by zero", func() {
            convey.So(func() { Divide(10, 0) }, convey.ShouldPanic)
        })

        convey.Convey("When dividing by a non-zero integer", func() {
            result := Divide(10, 2)
            convey.So(result, convey.ShouldEqual, 5)
        })
    })
}