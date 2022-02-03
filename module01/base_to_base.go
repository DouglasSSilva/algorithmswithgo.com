package module01

import "fmt"

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//
func BaseToBase(value string, base, newBase int) string {
	val := BaseToDec(value, base)
	if newBase == 10 {
		valString := fmt.Sprintf("%d", val)
		return valString

	}

	return DecToBase(val, newBase)
}
