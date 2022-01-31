package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	for i := range list {
		if num == list[i] {
			return true
		}
	}
	return false
}
