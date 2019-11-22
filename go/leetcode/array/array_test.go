package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test15(t *testing.T) {
	tests := []struct{
		input []int
		output [][]int
	}{
		{
			input:	[]int{-1, 0, 1},
			output: [][]int{{-1, 0, 1}},
		},
		{
			input:	[]int{-1, 0, 1, 1},
			output:	[][]int{{-1, 0, 1}},
		},
		{
			input: 	[]int{-1, 0, 1, 2, -2},
			output:	[][]int{{-1, 0, 1}, {0, 2, -2}},
		},
	}

	for _, tt := range tests{
		t.Run("test", func(t *testing.T) {
			res := threeSum(tt.input)
			assert.Equal(t, tt.output, res)
			res = threeSum(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test33(t *testing.T) {
	tests := []struct{
		nums []int
		target int
		output int
	}{
		{
			nums: []int{4, 5, 6, 7, 8, 1, 2, 3},
			target: 7,
			output: 3,
		},
		{
			nums: []int{4, 5, 6, 7, 8, 1, 2, 3},
			target: 9,
			output: -1,
		},
		{
			nums: []int{2, 0, 1},
			target: 0,
			output: 0,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := search(tt.nums, tt.target)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test41(t *testing.T) {
	tests := []struct{
		input []int
		output int
	}{
		{
			input: []int{1, 2, 0},
			output: 3,
		},
		{
			input: []int{7, 8, 9},
			output: 1,
		},
		{
			input: []int{1, 2, 3},
			output: 1,
		},
	}

	for _, tt := range(tests) {
		t.Run("test", func(t *testing.T) {
			res := firstMissingPositive(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test56(t *testing.T)  {
	tests := []struct{
		input [][]int
		output [][]int
	}{
		{
			input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			output: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input: [][]int{{1, 3}},
			output: [][]int{{1, 3}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := merge(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test169(t *testing.T) {
	tests := []struct{
		input []int
		output int
	}{
		{
			input: []int{3, 2, 3},
			output: 3,
		},
		{
			input: []int{2, 2, 1, 1, 1, 2, 2},
			output: 2,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := majorityElement(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test189(t *testing.T)  {
	tests := []struct{
		nums []int
		k int
		output []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			k: 2,
			output: []int{4, 5, 1, 2, 3},
		},
		{
			nums: []int{1, 2, 1, 2, 1},
			k: 10,
			output: []int{1, 2, 1, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := rotate(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test239(t *testing.T)  {
	tests := []struct{
		nums []int
		k int
		output []int
	}{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k: 3,
			output: []int{3, 3, 5, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := maxSlidingWindow2(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test283(t *testing.T)  {
	arr := []int{9,0,4,0,5,0,2,7,2,1}
	fmt.Println("输入：", arr)
	moveZeroes(arr)
	fmt.Println("输出：", arr)
}

func Test442(t *testing.T) {
	tests := []struct{
		input []int
		output []int
	}{
		{
			input: []int{4,3,2,7,8,2,3,1},
			output: []int{2,3},
		},
	}

	for _, tt :=range tests {
		t.Run("tests", func(t *testing.T) {
			res := findDuplicates(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test1198(t *testing.T) {
	tests := []struct{
		input [][]int
		output int
	}{
		{
			input: [][]int{{1, 2, 3, 4, 5}, {1, 4, 5, 8, 10}, {1, 5, 7, 9, 11}, {1, 3, 5, 7, 9}},
			output: 1,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := smallestCommonElement(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}