package sliceskit

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilterSuite struct {
	suite.Suite
}

func (s *FilterSuite) TestFilter() {
	s.Run("should return nil when input slice is nil", func() {
		result := Filter[[]int](nil, func(e int) bool { return e%2 == 0 })
		s.Nil(result)
	})

	s.Run("should return nil when filterFunc is nil", func() {
		result := Filter([]int{1, 2, 3}, nil)
		s.Equal([]int{1, 2, 3}, result)
	})

	s.Run("should return filtered slice", func() {
		slice := []int{1, 2, 3, 4, 5}
		result := Filter(slice, func(e int) bool { return e%2 == 0 })
		s.Equal([]int{2, 4}, result)
	})
}

func (s *FilterSuite) TestFilterWithIndex() {
	s.Run("should return nil when input slice is nil", func() {
		result := FilterWithIndex[[]int](nil, func(_ int, i int) bool { return (i)%2 == 0 })
		s.Nil(result)
	})

	s.Run("should return nil when filterFunc is nil", func() {
		result := FilterWithIndex([]int{1, 2, 3}, nil)
		s.Equal([]int{1, 2, 3}, result)
	})

	s.Run("should return filtered slice", func() {
		slice := []int{1, 2, 3, 4, 5}
		result := FilterWithIndex(slice, func(_ int, i int) bool { return i%2 == 0 })
		s.Equal([]int{1, 3, 5}, result)
	})
}

func TestFilterSuite(t *testing.T) {
	suite.Run(t, new(FilterSuite))
}
