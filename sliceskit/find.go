package sliceskit

func Find[Slice ~[]E, E any](s Slice, findFunc func(E) bool) (E, bool) {
	for _, e := range s {
		if findFunc(e) {
			return e, true
		}
	}
	var zero E
	return zero, false
}

// FindPtr returns the first pointer in the slice that satisfies the predicate, or nil if not found
func FindPtr[Slice ~[]*E, E any](s Slice, findFunc func(*E) bool) *E {
	for _, ptr := range s {
		if findFunc(ptr) {
			return ptr
		}
	}
	return nil
}
