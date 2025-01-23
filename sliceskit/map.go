package sliceskit

func Map[Slice ~[]E, T any, E any](s Slice, mapFunc func(E) T) []T {
	if s == nil {
		return nil
	}

	if mapFunc == nil {
		return nil
	}

	r := make([]T, len(s))
	for i, e := range s {
		r[i] = mapFunc(e)
	}
	return r
}

func MapWithIndex[Slice ~[]E, T any, E any](s Slice, mapFunc func(E, int) T) []T {
	if s == nil {
		return nil
	}

	if mapFunc == nil {
		return nil
	}

	r := make([]T, len(s))
	for i, e := range s {
		r[i] = mapFunc(e, i)
	}
	return r
}
