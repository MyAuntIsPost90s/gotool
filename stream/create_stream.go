package stream

type Stream[T any] struct {
	Val []T
}

func CreateStream[T any](val []T) *Stream[T] {
	return &Stream[T]{
		Val: val,
	}
}

func (stream *Stream[T]) Filter(fn func(T) bool) *Stream[T] {
	stream.Val = Filter(stream.Val, fn)
	return stream
}

func (stream *Stream[T]) Map(fn func(T) any) *Stream[any] {
	return CreateStream(Map(stream.Val, fn))
}

func (stream *Stream[T]) Reduce(initial any, fn func(T, any) any) any {
	return Reduce(stream.Val, initial, fn)
}

func (stream *Stream[T]) Exists(fn func(T) bool) bool {
	return Exists(stream.Val, fn)
}
