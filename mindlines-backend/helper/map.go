package helper

func Map[T any, U any](input []T, fn func(T) U) []U {
	res := make([]U, len(input))
	for ix, el := range input {
		res[ix] = fn(el)
	}
	return res
}
