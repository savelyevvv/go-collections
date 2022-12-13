package list

import "errors"

func Append[T any](slice *[]T, elem T) {
	*slice = append(*slice, elem)
}

func Count[T any](slice *[]T, elem T, equal func(a, b T) bool) int {
	var count int
	for _, v := range *slice {
		if equal(v, elem) {
			count++
		}
	}
	return count
}

func Insert[T any](slice *[]T, index int, elem T) {
	s := *slice
	i, _ := alignIndex(index, len(s))
	if len(s) == cap(s) {
		t := make([]T, len(s), (cap(s)+1)*2)
		copy(t, s)
		s = t
	}
	s = s[:len(s)+1]
	copy(s[i+1:], s[i:])
	s[i] = elem
	*slice = s
}

func Reverse[T any](slice *[]T) {
	s := *slice
	var (
		i = 0
		j = len(s) - 1
	)
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func Clear[T any](slice *[]T) {
	*slice = nil
}

func Extend[T any](slice *[]T, elems ...T) {
	*slice = append(*slice, elems...)
}

func Pop[T any](slice *[]T, index int) T {
	i, err := alignIndex(index, len(*slice))
	if err != nil {
		panic(err.Error())
	}
	v := (*slice)[i]
	delete(slice, i)
	return v
}

func Copy[T any](slice *[]T) []T {
	s := *slice
	if s == nil {
		return nil
	}
	r := make([]T, len(s))
	copy(r, s)
	return r
}

func Index[T any](slice *[]T, elem T, equal func(a, b T) bool) int {
	var (
		i        int
		v        T
		contains bool
	)
	for i, v = range *slice {
		if equal(v, elem) {
			contains = true
			break
		}
	}
	var index int
	if contains {
		index = i
	} else {
		index = -1
	}
	return index
}

func Remove[T any](slice *[]T, elem T, equal func(a, b T) bool) {
	i := Index(slice, elem, equal)
	if i != -1 {
		delete(slice, i)
	}
}

func alignIndex(i int, n int) (int, error) {
	var err error
	switch {
	case i < 0:
		i += n
		if i < 0 {
			i = 0
			err = errors.New("index out of bounds")
		}
	default:
		if i >= n {
			i = n
			err = errors.New("index out of bounds")
		}
	}
	return i, err
}

func delete[T any](slice *[]T, i int) {
	s := *slice
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]
	*slice = s
}

func test() {
}
