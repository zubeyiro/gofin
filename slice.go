package gofin

func Map[T any](s []T, f func(T) T) {
	for i, e := range s {
		s[i] = f(e)
	}
}

func ForEach[T any](s []T, f func(int, T)) {
	for i, e := range s {
		f(i, e)
	}
}

func Filter[T comparable](s []T, f func(T) bool) []T {
	var ret []T

	for _, e := range s {
		if f(e) {
			ret = append(ret, e)
		}
	}

	return ret
}

func RemoveIndex[T comparable](s []T, i int) []T {
	var empty T

	copy(s[i:], s[i+1:])
	s[len(s)-1] = empty

	return s[:len(s)-1]
}

func RemoveMatching[T comparable](s []T, f func(T) bool) []T {
	var ret []T

	for _, e := range s {
		if !f(e) {
			ret = append(ret, e)
		}
	}

	return ret
}

func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}

	return false
}

func IndexOf[T comparable](s []T, v T) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}

	return -1
}

func Chunk[T any](s []T, size int) (c [][]T) {
	for size < len(s) {
		s, c = s[size:], append(c, s[0:size:size])
	}

	return append(c, s)
}
