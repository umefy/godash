package sliceskit_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/sliceskit"
)

type ReduceSuite struct {
	suite.Suite
}

// Reduce should return initial value when input slice is nil
func (s *ReduceSuite) TestReduce_NilSlice() {
	result := sliceskit.Reduce[[]int](nil, func(prev int, current int) int { return prev + current }, 0)
	s.Equal(0, result)
}

// Reduce should return reduced value
func (s *ReduceSuite) TestReduce_ReduceValue() {
	slice := []int{1, 2, 3}
	result := sliceskit.Reduce(slice, func(prev int, current int) int { return prev + current }, 0)
	s.Equal(6, result)
}

// ReduceWithIndex should return initial value when input slice is nil
func (s *ReduceSuite) TestReduceWithIndex_NilSlice() {
	result := sliceskit.ReduceWithIndex[[]int](nil, func(prev int, current int, _ int) int { return prev + current }, 0)
	s.Equal(0, result)
}

// ReduceWithIndex should return reduced value
func (s *ReduceSuite) TestReduceWithIndex_ReduceValue() {
	slice := []int{1, 2, 3}
	result := sliceskit.ReduceWithIndex(slice, func(prev int, current int, i int) int { return prev + current + i }, 0)
	s.Equal(9, result)
}

func (s *ReduceSuite) TestReduceWithFuncErr_ReduceFuncErr() {
	slice := []int{1, 2, 3}
	result, err := sliceskit.ReduceWithFuncErr(slice, func(prev int, current int) (int, error) {
		if current == 2 {
			return 0, errors.New("error")
		}
		return prev + current, nil
	}, 0)
	s.Equal(0, result)
	s.NotNil(err)
}

func (s *ReduceSuite) TestReduceWithIndexAndFuncErr_ReduceFuncErr() {
	slice := []int{1, 2, 3}
	result, err := sliceskit.ReduceWithIndexAndFuncErr(slice, func(prev int, current int, i int) (int, error) {
		if current == 2 {
			return 0, errors.New("error")
		}
		return prev + current + i, nil
	}, 0)
	s.Equal(0, result)
	s.NotNil(err)
}

func TestReduceSuite(t *testing.T) {
	suite.Run(t, new(ReduceSuite))
}
