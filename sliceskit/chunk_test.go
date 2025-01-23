package sliceskit_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/sliceskit"
)

type ChunkSuite struct {
	suite.Suite
}

// Chunk should return empty slice when input slice is nil
func (s *ChunkSuite) TestChunk_NilSlice() {
	result := sliceskit.Chunk[[]int](nil, 2)
	s.Equal([][]int{}, result)
}

// Chunk should return empty slice when size is zero
func (s *ChunkSuite) TestChunk_ZeroSize() {
	slice := []int{1, 2, 3, 4, 5}
	result := sliceskit.Chunk(slice, 0)
	s.Equal([][]int{}, result)
}

// Chunk should return empty slice when size is negative
func (s *ChunkSuite) TestChunk_NegativeSize() {
	slice := []int{1, 2, 3, 4, 5}
	result := sliceskit.Chunk(slice, -1)
	s.Equal([][]int{}, result)
}

// Chunk should work with empty slice
func (s *ChunkSuite) TestChunk_EmptySlice() {
	slice := []int{}
	result := sliceskit.Chunk(slice, 2)
	s.Equal([][]int{}, result)
}

// Chunk should create chunks of exact size when slice length is divisible by size
func (s *ChunkSuite) TestChunk_ExactDivisible() {
	slice := []int{1, 2, 3, 4, 5, 6}
	result := sliceskit.Chunk(slice, 2)
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
	s.Equal(expected, result)
}

// Chunk should create chunks with last chunk smaller when slice length is not divisible by size
func (s *ChunkSuite) TestChunk_NotDivisible() {
	slice := []int{1, 2, 3, 4, 5}
	result := sliceskit.Chunk(slice, 2)
	expected := [][]int{{1, 2}, {3, 4}, {5}}
	s.Equal(expected, result)
}

// Chunk should work when size is larger than slice length
func (s *ChunkSuite) TestChunk_SizeLargerThanSlice() {
	slice := []int{1, 2, 3}
	result := sliceskit.Chunk(slice, 5)
	expected := [][]int{{1, 2, 3}}
	s.Equal(expected, result)
}

// Chunk should work with size of 1
func (s *ChunkSuite) TestChunk_SizeOne() {
	slice := []int{1, 2, 3}
	result := sliceskit.Chunk(slice, 1)
	expected := [][]int{{1}, {2}, {3}}
	s.Equal(expected, result)
}

// Chunk should work with string slices
func (s *ChunkSuite) TestChunk_StringSlice() {
	slice := []string{"a", "b", "c", "d"}
	result := sliceskit.Chunk(slice, 2)
	expected := [][]string{{"a", "b"}, {"c", "d"}}
	s.Equal(expected, result)
}

// Chunk should work with custom types
func (s *ChunkSuite) TestChunk_CustomType() {
	type Person struct {
		Name string
		Age  int
	}

	slice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
		{Name: "David", Age: 40},
	}

	result := sliceskit.Chunk(slice, 2)
	expected := [][]Person{
		{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 30}},
		{{Name: "Charlie", Age: 35}, {Name: "David", Age: 40}},
	}
	s.Equal(expected, result)
}

// Chunk should preserve the original slice (no mutation)
func (s *ChunkSuite) TestChunk_NoMutation() {
	original := []int{1, 2, 3, 4, 5}
	slice := make([]int, len(original))
	copy(slice, original)

	result := sliceskit.Chunk(slice, 2)

	// Verify original slice is unchanged
	s.Equal(original, slice)

	// Verify chunks contain the correct elements
	expected := [][]int{{1, 2}, {3, 4}, {5}}
	s.Equal(expected, result)
}

// Chunk should work with single element slice
func (s *ChunkSuite) TestChunk_SingleElement() {
	slice := []int{42}
	result := sliceskit.Chunk(slice, 3)
	expected := [][]int{{42}}
	s.Equal(expected, result)
}

func TestChunkSuite(t *testing.T) {
	suite.Run(t, new(ChunkSuite))
}
