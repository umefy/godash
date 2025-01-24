package sliceskit

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapSuite struct {
	suite.Suite
}

func (s *MapSuite) TestMapFunc() {
	s.Run("should return nil when input slice is nil", func() {
		result := Map[[]int](nil, func(e int) int { return e * 2 })
		s.Nil(result)
	})

	s.Run("should return nil when mapFunc is nil", func() {
		result := Map[[]int, int]([]int{1, 2, 3}, nil)
		s.Nil(result)
	})

	s.Run("should return mapped slice", func() {
		slice := []int{1, 2, 3}
		result := Map(slice, func(e int) int { return e * 2 })
		s.Equal([]int{2, 4, 6}, result)
	})
}

func (s *MapSuite) TestMapFuncIndex() {
	s.Run("should return nil when input slice is nil", func() {
		result := MapWithIndex[[]int](nil, func(e int, i int) int { return e * i })
		s.Nil(result)
	})

	s.Run("should return nil when mapFunc is nil", func() {
		result := MapWithIndex[[]int, int]([]int{1, 2, 3}, nil)
		s.Nil(result)
	})

	s.Run("should return mapped slice", func() {
		slice := []int{1, 2, 3}
		result := MapWithIndex(slice, func(e int, i int) int { return e * i })
		s.Equal([]int{1, 2, 6}, result)
	})
}

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(MapSuite))
}
