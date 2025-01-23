package sliceskit

func Any[Slice ~[]E, E any](s Slice, anyFunc func(E) bool) bool {
	for _, e := range s {
		if anyFunc(e) {
			return true
		}
	}
	return false
}
