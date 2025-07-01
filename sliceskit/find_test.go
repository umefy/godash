package sliceskit_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/sliceskit"
)

type FindSuite struct {
	suite.Suite
}

// Find should return zero value and false when input slice is nil
func (s *FindSuite) TestFind_NilSlice() {
	result, found := sliceskit.Find[[]int](nil, func(e int) bool { return e > 0 })
	s.Equal(0, result)
	s.False(found)
}

// Find should return zero value and false when no element matches
func (s *FindSuite) TestFind_NoMatch() {
	slice := []int{1, 2, 3, 4, 5}
	result, found := sliceskit.Find(slice, func(e int) bool { return e > 10 })
	s.Equal(0, result)
	s.False(found)
}

// Find should return first matching element and true
func (s *FindSuite) TestFind_FirstMatch() {
	slice := []int{1, 2, 3, 4, 5}
	result, found := sliceskit.Find(slice, func(e int) bool { return e%2 == 0 })
	s.Equal(2, result)
	s.True(found)
}

// Find should return first matching element even if multiple match
func (s *FindSuite) TestFind_MultipleMatches() {
	slice := []int{1, 2, 3, 4, 5}
	result, found := sliceskit.Find(slice, func(e int) bool { return e%2 == 0 })
	s.Equal(2, result) // Should return first match (2), not 4
	s.True(found)
}

// Find should work with string slices
func (s *FindSuite) TestFind_StringSlice() {
	slice := []string{"apple", "banana", "cherry"}
	result, found := sliceskit.Find(slice, func(e string) bool { return len(e) > 5 })
	s.Equal("banana", result)
	s.True(found)
}

// Find should work with empty slice
func (s *FindSuite) TestFind_EmptySlice() {
	slice := []int{}
	result, found := sliceskit.Find(slice, func(e int) bool { return e > 0 })
	s.Equal(0, result)
	s.False(found)
}

// Find should work with custom types
func (s *FindSuite) TestFind_CustomType() {
	type Person struct {
		Name string
		Age  int
	}

	slice := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	result, found := sliceskit.Find(slice, func(p Person) bool { return p.Age >= 30 })
	s.Equal(Person{Name: "Bob", Age: 30}, result)
	s.True(found)
}

// FindPtr should return nil when input slice is nil
func (s *FindSuite) TestFindPtr_NilSlice() {
	result := sliceskit.FindPtr[[]*int](nil, func(e *int) bool { return e != nil && *e > 0 })
	s.Nil(result)
}

// FindPtr should return nil when no pointer matches
func (s *FindSuite) TestFindPtr_NoMatch() {
	slice := []*int{ptr(1), ptr(2), ptr(3)}
	result := sliceskit.FindPtr(slice, func(e *int) bool { return *e > 10 })
	s.Nil(result)
}

// FindPtr should return the first pointer that matches
func (s *FindSuite) TestFindPtr_FirstMatch() {
	slice := []*int{ptr(1), ptr(2), ptr(3)}
	result := sliceskit.FindPtr(slice, func(e *int) bool { return *e%2 == 0 })
	s.NotNil(result)
	s.Equal(slice[1], result)
}

// FindPtr should work with nil pointers in the slice
func (s *FindSuite) TestFindPtr_WithNilPointers() {
	slice := []*int{nil, ptr(2), nil}
	result := sliceskit.FindPtr(slice, func(e *int) bool { return e != nil && *e == 2 })
	s.NotNil(result)
	s.Equal(slice[1], result)
}

// FindPtr should work with custom types
func (s *FindSuite) TestFindPtr_CustomType() {
	type Person struct {
		Name string
		Age  int
	}
	alice := &Person{"Alice", 25}
	bob := &Person{"Bob", 30}
	charlie := &Person{"Charlie", 35}
	slice := []*Person{alice, bob, charlie}
	result := sliceskit.FindPtr(slice, func(p *Person) bool { return p.Age >= 30 })
	s.Equal(bob, result)
}

// helper for pointer values
func ptr[T any](v T) *T { return &v }

func TestFindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}
