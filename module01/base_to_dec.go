package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {

	valueSize := len(value)

	total := float64(0)
	for i := 0; i < valueSize; i++ {
		numberAsString := string(value[valueSize-(i+1)])
		position := findPosition(numberAsString)

		total += position * math.Pow(float64(base), float64(i))

	}
	return int(total)
}

func findPosition(numberAsString string) float64 {
	modString := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	position := 0
	for i, _ := range modString {
		if modString[i] == numberAsString {
			position = i
		}
	}

	return float64(position)
}
