package sliceskit

// Filter a slice of type E based on filterFunc
// Will generate a new slice, won't change original slice
func Filter[Slice ~[]E, E any](s Slice, filterFunc func(E) bool) Slice {
	r, _ := FilterWithFuncErr(s, func(e E) (bool, error) {
		return filterFunc(e), nil
	})

	return r
}

// FilterWithIndex is same with Filter, but allow filter function to have index
func FilterWithIndex[Slice ~[]E, E any](s Slice, filterFunc func(E, int) bool) Slice {
	r, _ := FilterWithIndexAndFuncErr(s, func(e E, i int) (bool, error) {
		return filterFunc(e, i), nil
	})

	return r
}

func FilterWithFuncErr[Slice ~[]E, E any](s Slice, filterFunc func(E) (bool, error)) (Slice, error) {

	var result Slice
	for _, e := range s {
		ok, err := filterFunc(e)
		if err != nil {
			return nil, err
		}

		if ok {
			result = append(result, e)
		}
	}

	return result, nil
}

func FilterWithIndexAndFuncErr[Slice ~[]E, E any](s Slice, filterFunc func(E, int) (bool, error)) (Slice, error) {
	if s == nil {
		return nil, nil
	}

	var result Slice
	for i, e := range s {
		ok, err := filterFunc(e, i)
		if err != nil {
			return nil, err
		}

		if ok {
			result = append(result, e)
		}
	}

	return result, nil
}
