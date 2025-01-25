package sliceskit

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilterSuite struct {
	suite.Suite
}

// Filter should return nil when input slice is nil
func (s *FilterSuite) TestFilter_NilSlice() {
	result := Filter[[]int](nil, func(e int) bool { return e%2 == 0 })
	s.Nil(result)
}

// Filter should return nil when filterFunc is nil
func (s *FilterSuite) TestFilter_NilFilterFunc() {
	result := Filter([]int{1, 2, 3}, nil)
	s.Equal([]int{1, 2, 3}, result)
}

// Filter should return filtered slice
func (s *FilterSuite) TestFilter_FilteredSlice() {
	slice := []int{1, 2, 3, 4, 5}
	result := Filter(slice, func(e int) bool { return e%2 == 0 })
	s.Equal([]int{2, 4}, result)
}

// FilterWithIndex should return nil when input slice is nil
func (s *FilterSuite) TestFilterWithIndex_NilSlice() {
	result := FilterWithIndex[[]int](nil, func(_ int, i int) bool { return i%2 == 0 })
	s.Nil(result)
}

// FilterWithIndex should return nil when filterFunc is nil
func (s *FilterSuite) TestFilterWithIndex_NilFilterFunc() {
	result := FilterWithIndex([]int{1, 2, 3}, nil)
	s.Equal([]int{1, 2, 3}, result)
}

// FilterWithIndex should return filtered slice
func (s *FilterSuite) TestFilterWithIndex_FilteredSlice() {
	slice := []int{0, 1, 2}
	result := FilterWithIndex(slice, func(e int, i int) bool { return (e*i)%2 == 0 })
	s.Equal([]int{0, 2}, result)
}

func TestFilterSuite(t *testing.T) {
	suite.Run(t, new(FilterSuite))
}
