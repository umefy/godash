package sliceskit_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/sliceskit"
)

type EverySuite struct {
	suite.Suite
}

// Every should return true when input slice is nil (vacuous truth)
func (s *EverySuite) TestEvery_NilSlice() {
	result := sliceskit.Every[[]int](nil, func(e int) bool { return e > 0 })
	s.True(result)
}

// Every should return true when all elements match
func (s *EverySuite) TestEvery_AllMatch() {
	slice := []int{2, 4, 6, 8, 10}
	result := sliceskit.Every(slice, func(e int) bool { return e%2 == 0 })
	s.True(result)
}

// Every should return false when at least one element doesn't match
func (s *EverySuite) TestEvery_OneDoesNotMatch() {
	slice := []int{2, 3, 4, 6, 8}
	result := sliceskit.Every(slice, func(e int) bool { return e%2 == 0 })
	s.False(result)
}

// Every should return false when no elements match
func (s *EverySuite) TestEvery_NoMatch() {
	slice := []int{1, 3, 5, 7, 9}
	result := sliceskit.Every(slice, func(e int) bool { return e%2 == 0 })
	s.False(result)
}

// Every should work with string slices
func (s *EverySuite) TestEvery_StringSlice() {
	slice := []string{"apple", "banana", "cherry"}
	result := sliceskit.Every(slice, func(e string) bool { return len(e) > 3 })
	s.True(result)
}

// Every should work with empty slice (vacuous truth)
func (s *EverySuite) TestEvery_EmptySlice() {
	slice := []int{}
	result := sliceskit.Every(slice, func(e int) bool { return e > 0 })
	s.True(result)
}

// Every should work with custom types
func (s *EverySuite) TestEvery_CustomType() {
	type Person struct {
		Name string
		Age  int
	}

	slice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	result := sliceskit.Every(slice, func(p Person) bool { return p.Age >= 25 })
	s.True(result)
}

// Every should return true for zero values when predicate is always true
func (s *EverySuite) TestEvery_ZeroValues() {
	slice := []int{0, 0, 0}
	result := sliceskit.Every(slice, func(e int) bool { return e == 0 })
	s.True(result)
}

// Every should return false for zero values when predicate checks for non-zero
func (s *EverySuite) TestEvery_ZeroValuesFalse() {
	slice := []int{0, 0, 0}
	result := sliceskit.Every(slice, func(e int) bool { return e > 0 })
	s.False(result)
}

// Every should work with single element slice
func (s *EverySuite) TestEvery_SingleElement() {
	slice := []int{5}
	result := sliceskit.Every(slice, func(e int) bool { return e > 0 })
	s.True(result)
}

// Every should work with single element slice that doesn't match
func (s *EverySuite) TestEvery_SingleElementNoMatch() {
	slice := []int{5}
	result := sliceskit.Every(slice, func(e int) bool { return e < 0 })
	s.False(result)
}

func TestEverySuite(t *testing.T) {
	suite.Run(t, new(EverySuite))
}
