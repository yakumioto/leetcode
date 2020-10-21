package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type q struct {
	s string
}

type a struct {
	result int
}

type qa struct {
	q
	a
}

var (
	qas = []qa{
		{q{s: "abcabcbb"}, a{result: 3}},
		{q{s: "bbbbb"}, a{result: 1}},
		{q{s: "pwwkew"}, a{result: 3}},
		{q{s: " "}, a{result: 1}},
		{q{s: "dvdf"}, a{result: 3}},
	}
)

func TestLengthOfLongestSubstringWithExhaustive(t *testing.T) {
	for _, qa := range qas {
		assert.Equal(t,
			qa.a.result,
			LengthOfLongestSubstringWithExhaustive(qa.q.s),
			"input question: s: %v", qa.q.s)
	}
}

func TestLengthOfLongestSubstringWithSlidingWindow(t *testing.T) {
	for _, qa := range qas {
		assert.Equal(t,
			qa.a.result,
			LengthOfLongestSubstringWithSlidingWindow(qa.q.s),
			"input question: s: %v", qa.q.s)
	}
}
