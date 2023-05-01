package helpers

func Filter[T any](vs []T, f func(T) bool) []T {
	res := []T{}
	for _, v := range vs {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Map[T, U any](vs []T, f func(T) U) []U {
	res := []U{}
	for _, v := range vs {
		res = append(res, f(v))
	}
	return res
}
