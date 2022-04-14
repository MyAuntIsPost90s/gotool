package stream

func Reduce[T any, K any](arr []T, initial K, fn func(T, K) K) K {
	result := initial
	for _, v := range arr {
		result = fn(v, result)
	}
	return result
}
