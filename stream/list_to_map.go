package stream

func ListToMap[T any, T2 int | string | bool | *any, T3 any](arr []T, fn1 func(T) T2, fn2 func(T) T3) map[T2]T3 {
	result := make(map[T2]T3)
	for _, v := range arr {
		result[fn1(v)] = fn2(v)
	}
	return result
}
