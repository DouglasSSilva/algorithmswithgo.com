package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	modString := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	value := ""
	for dec >= base {
		mod := dec % base
		value = fmt.Sprintf("%s%s", modString[mod], value)
		dec = dec / base

	}
	value = fmt.Sprintf("%s%s", modString[dec], value)
	return value
}

// func decToBase2Bitwise(dec int) {
// 	var base2 uint
// 	base2 := dec << 31

// }
