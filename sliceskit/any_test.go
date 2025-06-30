package sliceskit_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/sliceskit"
)

type AnySuite struct {
	suite.Suite
}

// Any should return false when input slice is nil
func (s *AnySuite) TestAny_NilSlice() {
	result := sliceskit.Any[[]int](nil, func(e int) bool { return e > 0 })
	s.False(result)
}

// Any should return false when no element matches
func (s *AnySuite) TestAny_NoMatch() {
	slice := []int{1, 2, 3, 4, 5}
	result := sliceskit.Any(slice, func(e int) bool { return e > 10 })
	s.False(result)
}

// Any should return true when at least one element matches
func (s *AnySuite) TestAny_OneMatch() {
	slice := []int{1, 2, 3, 4, 5}
	result := sliceskit.Any(slice, func(e int) bool { return e == 3 })
	s.True(result)
}

// Any should return true when multiple elements match
func (s *AnySuite) TestAny_MultipleMatches() {
	slice := []int{1, 2, 3, 4, 5}
	result := sliceskit.Any(slice, func(e int) bool { return e%2 == 0 })
	s.True(result)
}

// Any should return true when all elements match
func (s *AnySuite) TestAny_AllMatch() {
	slice := []int{2, 4, 6, 8, 10}
	result := sliceskit.Any(slice, func(e int) bool { return e%2 == 0 })
	s.True(result)
}

// Any should work with string slices
func (s *AnySuite) TestAny_StringSlice() {
	slice := []string{"apple", "banana", "cherry"}
	result := sliceskit.Any(slice, func(e string) bool { return len(e) > 5 })
	s.True(result)
}

// Any should work with empty slice
func (s *AnySuite) TestAny_EmptySlice() {
	slice := []int{}
	result := sliceskit.Any(slice, func(e int) bool { return e > 0 })
	s.False(result)
}

// Any should work with custom types
func (s *AnySuite) TestAny_CustomType() {
	type Person struct {
		Name string
		Age  int
	}

	slice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	result := sliceskit.Any(slice, func(p Person) bool { return p.Age > 30 })
	s.True(result)
}

// Any should return false for zero values when predicate checks for non-zero
func (s *AnySuite) TestAny_ZeroValues() {
	slice := []int{0, 0, 0}
	result := sliceskit.Any(slice, func(e int) bool { return e > 0 })
	s.False(result)
}

func TestAnySuite(t *testing.T) {
	suite.Run(t, new(AnySuite))
}
