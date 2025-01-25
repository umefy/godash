package sliceskit

func Reduce[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E) U, initial U) U {
	r, _ := ReduceWithFuncErr(s, func(prev U, current E) (U, error) {
		return reduceFunc(prev, current), nil
	}, initial)
	return r
}

func ReduceWithIndex[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E, index int) U, initial U) U {
	r, _ := ReduceWithIndexAndFuncErr(s, func(prev U, current E, index int) (U, error) {
		return reduceFunc(prev, current, index), nil
	}, initial)

	return r
}

func ReduceWithFuncErr[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E) (U, error), initial U) (U, error) {
	if s == nil {
		return initial, nil
	}

	r := initial
	var err error
	for _, e := range s {
		r, err = reduceFunc(r, e)
		if err != nil {
			return initial, err
		}
	}

	return r, nil
}

func ReduceWithIndexAndFuncErr[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E, index int) (U, error), initial U) (U, error) {
	if s == nil {
		return initial, nil
	}

	r := initial
	var err error
	for i, e := range s {
		r, err = reduceFunc(r, e, i)
		if err != nil {
			return initial, err
		}
	}

	return r, nil
}
