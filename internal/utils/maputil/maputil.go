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

func MapFunc[K1, K2 comparable, T1, T2 any](source map[K1]T1, mapFunc func(map[K2]T2, K1, T1) (K2, T2)) map[K2]T2 {
	secondMap := make(map[K2]T2, len(source))
	for k, v := range source {
		k1, k2 := mapFunc(secondMap, k, v)
		secondMap[k1] = k2
	}
	return secondMap
}
