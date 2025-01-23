package sliceskit

// Map a slice of type E to a slice of type T
// Will generate a new slice, won't change original slice
func Map[Slice ~[]E, T any, E any](s Slice, mapFunc func(E) T) []T {
	r, _ := MapWithFuncErr(s, func(e E) (T, error) {
		return mapFunc(e), nil
	})

	return r
}

// MapWithIndex is same with Map, but allow map function to have index
func MapWithIndex[Slice ~[]E, T any, E any](s Slice, mapFunc func(E, int) T) []T {
	r, _ := MapWithIndexAndFuncErr(s, func(e E, i int) (T, error) {
		return mapFunc(e, i), nil
	})

	return r
}

// MapWithFuncErr is same with Map, but allow map function to return error
func MapWithFuncErr[Slice ~[]E, T any, E any](s Slice, mapFunc func(E) (T, error)) ([]T, error) {

	if s == nil {
		return nil, nil
	}

	r := make([]T, len(s))
	for i, e := range s {
		t, err := mapFunc(e)
		if err != nil {
			return nil, err
		}
		r[i] = t
	}
	return r, nil
}

// MapWithIndexAndFuncErr is same with MapWithIndex, but allow map function to return error
func MapWithIndexAndFuncErr[Slice ~[]E, T any, E any](s Slice, mapFunc func(E, int) (T, error)) ([]T, error) {

	if s == nil {
		return nil, nil
	}

	r := make([]T, len(s))
	for i, e := range s {
		t, err := mapFunc(e, i)
		if err != nil {
			return nil, err
		}
		r[i] = t
	}
	return r, nil
}
