package stream

import "strings"

func Join[T any](arr []T, sep string, fn func(T) string) string {
	strList := make([]string, 0)
	for _, v := range arr {
		strList = append(strList, fn(v))
	}
	return strings.Join(strList, sep)
}
