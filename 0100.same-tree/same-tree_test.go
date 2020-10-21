package problem0100

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yakumioto/leetcode/helper"
)

type q struct {
	a []int
	b []int
}

type a struct {
	result bool
}

type qa struct {
	q
	a
}

var (
	qas = []qa{
		{q{a: []int{}, b: []int{}}, a{result: true}},
		{q{a: []int{}, b: []int{1}}, a{result: false}},
		{q{a: []int{1}, b: []int{1}}, a{result: true}},
		{q{a: []int{1, 2, 3}, b: []int{1, 2, 3}}, a{result: true}},
		{q{a: []int{1, 2}, b: []int{1, helper.NULL, 3}}, a{result: false}},
		{q{a: []int{1, 2, 1}, b: []int{1, 1, 3}}, a{result: false}},
	}
)

func TestIsSameTree(t *testing.T) {
	for _, qa := range qas {
		assert.Equal(t,
			qa.a.result,
			IsSameTree(helper.IntArrayToTreeNode(qa.q.a), helper.IntArrayToTreeNode(qa.q.b)),
			"a: %v, b: %v", qa.q.a, qa.q.b)
	}
}
