package maputil

func Values[U comparable, T any](source map[U]T) []T {
	list := make([]T, len(source))
	i := 0
	for _, v := range source {
		list[i] = v
		i++
	}

	return list
}
