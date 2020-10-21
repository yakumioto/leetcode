package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type q struct {
	nums   []int
	target int
}

type a struct {
	result []int
}

type qa struct {
	q
	a
}

var (
	qas = []qa{
		{q{nums: []int{2, 7, 11, 15}, target: 9}, a{result: []int{1, 1}}},
		{q{nums: []int{2, 7, 11, 15}, target: 0}, a{result: nil}},
	}
)

func TestTwoSum(t *testing.T) {
	for _, qa := range qas {
		assert.Equalf(t,
			qa.a.result,
			TwoSum(qa.q.nums, qa.q.target),
			"input question: nums: %v, target: %v", qa.q.nums, qa.q.target)
	}
}
