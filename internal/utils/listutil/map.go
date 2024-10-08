package listutil

func Map[A, B any](list []A, mapFn func(A, int) B) []B {
	mappedList := make([]B, len(list))
	for i, el := range list {
		mappedList[i] = mapFn(el, i)
	}
	return mappedList
}
