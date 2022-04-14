package stream

func Map[T any, K any](arr []T, fn func(T) K) []K {
	result := make([]K, 0)
	for _, v := range arr {
		result = append(result, fn(v))
	}
	return result
}
