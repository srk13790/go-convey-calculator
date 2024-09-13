package calculator

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestEvenNumber(t *testing.T) {
    convey.Convey("Given an even number", t, func() {
        number := 3

        convey.Convey("It should be even", func() {
            convey.So(number, ShouldBeEven)
        })
    })

    convey.Convey("Given an odd number", t, func() {
        number := 5

        convey.Convey("It should fail the even check", func() {
            convey.So(number, ShouldBeEven)
        })
    })
}

// func TestNonRepeatingCharacter(t *testing.T) {
//     convey.Convey("Checking For Non Repeating Chanaracter", t, func ()  {
//         str := "swiss"
//         expected := "w"
//         actual := ReturnFirstNonRepeatletter(str)

//         convey.Convey("Expected chatacter here is %s", expected, func ()  {
//             convey.So(actual, ShouldReturnFirstRepeatletter(actual, expected))
//         })

        
//     })
// }