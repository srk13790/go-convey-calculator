package calculator

// func PanicTests(a, b int) int {
// 	if (a < b) {
// 		panic("Should be greater than b")
// 	} else {
// 		return 1
// 	}
// }

func PanicTests(a, b int) int {
    if a < b {
		panic("a should be greater than b")
	}
	return 1
}

func Divide(a, b int) int {
    if b == 0 {
        panic("division by zero")   
    }
    return a / b
}