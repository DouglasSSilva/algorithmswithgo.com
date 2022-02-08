package module01

func GCD(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}

	for b != 0 {

		a, b = b, a%b

	}

	return a
}
