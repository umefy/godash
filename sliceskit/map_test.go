package sliceskit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mapSuite struct {
	suite.Suite
}

func (s *mapSuite) TestMapFunc() {
	s.Run("should return nil when input slice is nil", func() {
		result := Map[[]int](nil, func(e int) int { return e * 2 })
		assert.Nil(s.T(), result)
	})

	s.Run("should return nil when mapFunc is nil", func() {
		result := Map[[]int, int]([]int{1, 2, 3}, nil)
		assert.Nil(s.T(), result)
	})

	s.Run("should return mapped slice", func() {
		slice := []int{1, 2, 3}
		result := Map(slice, func(e int) int { return e * 2 })
		assert.Equal(s.T(), []int{2, 4, 6}, result)
	})
}

func (s *mapSuite) TestMapFuncIndex() {
	s.Run("should return nil when input slice is nil", func() {
		result := MapWithIndex[[]int](nil, func(e int, i int) int { return e * i })
		assert.Nil(s.T(), result)
	})

	s.Run("should return nil when mapFunc is nil", func() {
		result := MapWithIndex[[]int, int]([]int{1, 2, 3}, nil)
		assert.Nil(s.T(), result)
	})

	s.Run("should return mapped slice", func() {
		slice := []int{1, 2, 3}
		result := MapWithIndex(slice, func(e int, i int) int { return e * i })
		assert.Equal(s.T(), []int{0, 2, 6}, result)
	})
}

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(mapSuite))
}
