package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type q struct {
	l1 *ListNode
	l2 *ListNode
}

type a struct {
	result *ListNode
}

type qa struct {
	q
	a
}

var (
	qas = []qa{
		{
			q{
				l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
				l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			},
			a{
				result: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}},
		},
		{
			q{
				l1: &ListNode{Val: 5},
				l2: &ListNode{Val: 5},
			},
			a{
				result: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
			},
		},
		{
			q{
				l1: &ListNode{Val: 1},
				l2: &ListNode{Val: 9, Next: &ListNode{Val: 9}},
			},
			a{
				result: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
			},
		},
	}
)

func TestAddTwoNumbers(t *testing.T) {
	for _, qa := range qas {
		assert.Equal(t, qa.a.result, AddTwoNumbers(qa.q.l1, qa.q.l2), "input question: %v", qa.q)
	}
}
