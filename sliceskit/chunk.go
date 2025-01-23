package sliceskit

func Chunk[Slice ~[]E, E any](s Slice, size int) [][]E {
	if size <= 0 {
		return [][]E{}
	}

	chunks := make([][]E, 0, len(s)/size+1)
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}
