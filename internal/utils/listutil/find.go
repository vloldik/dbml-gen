package listutil

// Searches for a value in list that matches function
func SearchFunc[T any](array []T, accept func(T, int) bool) (found T) {
	for i, el := range array {
		if accept(el, i) {
			return el
		}
	}
	return
}
