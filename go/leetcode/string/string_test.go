package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test8(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "   -42",
			output: -42,
		},
		{
			input:  "4193 with words",
			output: 4193,
		},
		{
			input:  "words and 987",
			output: 0,
		},
		{
			input:  "-91283472332",
			output: -2147483648,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := myAtoi(tt.input)
			assert.Equal(t, tt.output, res)
		})

	}

}

func Test10(t *testing.T) {
	tests := []struct {
		input  string
		p      string
		output bool
	}{
		{
			input:  "aa",
			p:      "a",
			output: false,
		},
		{
			input:  "aa",
			p:      "a*",
			output: true,
		},
		{
			input:  "ab",
			p:      ".*",
			output: true,
		},
		{
			input:  "aab",
			p:      "c*a*b",
			output: true,
		},
		{
			input:  "mississippi",
			p:      "mis*is*p*",
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := isMatch2(tt.input, tt.p)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test20(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "([[{[[]]}]])",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := isValid(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test32(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "(()",
			output: 2,
		},
		{
			input:  ")()())",
			output: 4,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := longestValidParentheses(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test139(t *testing.T) {
	tests := []struct {
		s      string
		w      []string
		output bool
	}{
		{
			s:      "leetcode",
			w:      []string{"leet", "code"},
			output: true,
		},
		{
			s:      "applepenapple",
			w:      []string{"apple", "pen"},
			output: true,
		},
		{
			s:      "catsandog",
			w:      []string{"cats", "dog", "sand", "and", "cat"},
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := wordBreak(tt.s, tt.w)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test150(t *testing.T) {
	tests := []struct {
		input  []string
		output int
	}{
		{
			input:  []string{"2", "1", "+", "3", "*"},
			output: 9,
		},
		{
			input:  []string{"4", "13", "5", "/", "+"},
			output: 6,
		},
		{
			input:  []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			output: 22,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := evalRPN(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test316(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "bcabc",
			output: "abc",
		},
		{
			input:  "cbacdcbc",
			output: "acdb",
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := removeDuplicateLetters(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test567(t *testing.T) {
	tests := []struct {
		s1     string
		s2     string
		output bool
	}{
		{
			s1:     "ab",
			s2:     "eidbaooo",
			output: true,
		},
		{
			s1:     "ab",
			s2:     "eidboaoo",
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := checkInclusion(tt.s1, tt.s2)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test647(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "abc",
			output: 3,
		},
		{
			input:  "aaa",
			output: 6,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := countSubstrings(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
