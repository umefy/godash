package sliceskit

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReduceSuite struct {
	suite.Suite
}

// Reduce should return initial value when input slice is nil
func (s *ReduceSuite) TestReduce_NilSlice() {
	result := Reduce[[]int](nil, func(prev int, current int) int { return prev + current }, 0)
	s.Equal(0, result)
}

// Reduce should return initial value when reduceFunc is nil
func (s *ReduceSuite) TestReduce_NilReduceFunc() {
	result := Reduce([]int{1, 2, 3}, nil, 0)
	s.Equal(0, result)
}

// Reduce should return reduced value
func (s *ReduceSuite) TestReduce_ReduceValue() {
	slice := []int{1, 2, 3}
	result := Reduce(slice, func(prev int, current int) int { return prev + current }, 0)
	s.Equal(6, result)
}

// ReduceWithIndex should return initial value when input slice is nil
func (s *ReduceSuite) TestReduceWithIndex_NilSlice() {
	result := ReduceWithIndex[[]int](nil, func(prev int, current int, _ int) int { return prev + current }, 0)
	s.Equal(0, result)
}

// ReduceWithIndex should return initial value when reduceFunc is nil
func (s *ReduceSuite) TestReduceWithIndex_NilReduceFunc() {
	result := ReduceWithIndex([]int{1, 2, 3}, nil, 0)
	s.Equal(0, result)
}

// ReduceWithIndex should return reduced value
func (s *ReduceSuite) TestReduceWithIndex_ReduceValue() {
	slice := []int{1, 2, 3}
	result := ReduceWithIndex(slice, func(prev int, current int, i int) int { return prev + current + i }, 0)
	s.Equal(9, result)
}

func TestReduceSuite(t *testing.T) {
	suite.Run(t, new(ReduceSuite))
}
