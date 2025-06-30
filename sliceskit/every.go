package sliceskit

func Every[Slice ~[]E, E any](s Slice, everyFunc func(E) bool) bool {
	for _, e := range s {
		if !everyFunc(e) {
			return false
		}
	}
	return true
}
