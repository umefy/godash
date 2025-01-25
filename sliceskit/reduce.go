package sliceskit

func Reduce[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E) U, initial U) U {
	if s == nil || reduceFunc == nil {
		return initial
	}

	r := initial
	for _, e := range s {
		r = reduceFunc(r, e)
	}

	return r
}

func ReduceWithIndex[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E, index int) U, initial U) U {
	if s == nil || reduceFunc == nil {
		return initial
	}

	r := initial
	for i, e := range s {
		r = reduceFunc(r, e, i)
	}

	return r
}
