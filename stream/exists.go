package stream

func Exists[T any](arr []T, fn func(T) bool) bool {
	for _, v := range arr {
		if fn(v) {
			return true
		}
	}
	return false
}
