package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test29(t *testing.T)  {
	tests := []struct{
		dividend int
		divisor int
		output int
	}{
		{
			dividend: 10,
			divisor: 3,
			output: 3,
		},
		{
			dividend: 1,
			divisor: 2,
			output: 0,
		},
		{
			dividend: 7,
			divisor: -3,
			output: -2,
		},
		{
			dividend: 1,
			divisor: 1,
			output: 1,
		},
		{
			dividend: 0,
			divisor: 1,
			output: 0,
		},
	}

	for _, tt := range tests {
		t.Run("tests", func(t *testing.T) {
			res := divide(tt.dividend, tt.divisor)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test39(t *testing.T)  {
	tests := []struct{
		candidates []int
		target int
		output [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target: 7,
			output: [][]int{{2,2,3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target: 8,
			output: [][]int{{2,2,2,2}, {2,3,3}, {3,5}},
		},
	}

	for _, tt := range tests {
		t.Run("tests", func(t *testing.T) {
			res := combinationSum(tt.candidates, tt.target)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test50(t *testing.T) {
	tests := []struct{
		x float64
		n int
		output float64
	}{
		{
			x: 2.00000,
			n: 10,
			output: 1024.00000,
		},
		{
			x: 2.10000,
			n: 3,
			output: 9.26100,
		},
		{
			x: 2.00000,
			n: -2,
			output: 0.25000,
		},
	}

	for _, tt :=range tests {
		t.Run("tests", func(t *testing.T) {
			res := myPow(tt.x, tt.n)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test60(t *testing.T) {
	tests := []struct{
		n int
		k int
		output string
	}{
		{
			n: 3,
			k: 3,
			output: "213",
		},
		{
			n: 4,
			k: 9,
			output: "2314",
		},
	}

	for _, tt :=range tests {
		t.Run("tests", func(t *testing.T) {
			res := getPermutation(tt.n, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test65(t *testing.T) {
	tests := []struct{
		input string
		output bool
	}{
		{
			input: " 6e-1",
			output: true,
		},
		{
			input: "99e2.5 ",
			output: false,
		},
		{
			input: "99e2.5 ",
			output: false,
		},
	}

	for _, tt :=range tests {
		t.Run("tests", func(t *testing.T) {
			res := isNumber(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test69(t *testing.T) {
	tests := []struct{
		input int
		output int
	}{
		{
			input: 4,
			output: 2,
		},
		{
			input: 9,
			output: 3,
		},
		{
			input: 8,
			output: 2,
		},
	}

	for _, tt :=range tests {
		t.Run("tests", func(t *testing.T) {
			res := mySqrt(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test268(t *testing.T) {
	tests := []struct{
		input []int
		output int
	}{
		{
			input: []int{3,0,1},
			output: 2,
		},
		{
			input: []int{9,6,4,2,3,5,7,0,1},
			output: 8,
		},
	}

	for _, tt :=range tests {
		t.Run("tests", func(t *testing.T) {
			res := missingNumber(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

