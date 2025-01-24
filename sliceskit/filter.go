package sliceskit

// Filter a slice of type E based on filterFunc
// Will generate a new slice, won't change original slice
func Filter[Slice ~[]E, E any](s Slice, filterFunc func(E) bool) Slice {
	if s == nil {
		return nil
	}

	if filterFunc == nil {
		filterFunc = func(E) bool {
			return true
		}
	}

	var result Slice
	for _, e := range s {
		if filterFunc(e) {
			result = append(result, e)
		}
	}

	return result
}

// FilterWithIndex is same with Filter, but allow filter function to have index
func FilterWithIndex[Slice ~[]E, E any](s Slice, filterFunc func(E, int) bool) Slice {
	if s == nil {
		return nil
	}

	if filterFunc == nil {
		filterFunc = func(E, int) bool {
			return true
		}
	}

	var result Slice
	for i, e := range s {
		if filterFunc(e, i) {
			result = append(result, e)
		}
	}

	return result
}
