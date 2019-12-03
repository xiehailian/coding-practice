package tree

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test94(t *testing.T) {
	nodeG := TreeNode{Val: 1, Left: nil, Right: nil}
	nodeF := TreeNode{Val: 2, Left: &nodeG, Right: nil}
	nodeE := TreeNode{Val: 3, Left: nil, Right: nil}
	nodeD := TreeNode{Val: 4, Left: &nodeE, Right: nil}
	nodeC := TreeNode{Val: 5, Left: nil, Right: nil}
	nodeB := TreeNode{Val: 6, Left: &nodeD, Right: &nodeF}
	nodeA := TreeNode{Val: 7, Left: &nodeB, Right: &nodeC}

	result := inorderTraversal(&nodeA)
	result2 := inorderTraversal2(&nodeA)
	t.Log(result)
	t.Log(result2)
	expected := []int{3, 4, 6, 1, 2, 7, 5}
	if !reflect.DeepEqual(result, expected) && !reflect.DeepEqual(result2, expected) {
		t.Fatalf("expected: %v, but got: %v",
			expected, result)
	}
}

func Test98(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output bool
	}{
		{
			input:  &TreeNode{6, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
			output: true,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := isValidBST(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test100(t *testing.T) {
	tests := []struct {
		p      *TreeNode
		q      *TreeNode
		output bool
	}{
		{
			p:      &TreeNode{6, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
			q:      &TreeNode{6, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
			output: true,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := isSameTree(tt.p, tt.q)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test102(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output [][]int
	}{
		{
			input: &TreeNode{7,
				&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{8, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}},
			output: [][]int{{7}, {4, 8}, {3, 5, 6, 9}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := levelOrder2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test104(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output int
	}{
		{
			input: &TreeNode{7,
				&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{8, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}},
			output: 3,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := maxDepth(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test112(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		sum    int
		output bool
	}{
		{
			input: &TreeNode{7,
				&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{8, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}},
			sum:    14,
			output: true,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := hasPathSum(tt.input, tt.sum)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test226(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output *TreeNode
	}{
		{
			input: &TreeNode{4,
				&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
				&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}},
			output: &TreeNode{4,
				&TreeNode{7, &TreeNode{9, nil, nil}, &TreeNode{6, nil, nil}},
				&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := invertTree(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test285(t *testing.T) {
	tests := []struct {
		p      *TreeNode
		q      *TreeNode
		output *TreeNode
	}{
		{
			p: &TreeNode{6,
				&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}},
			q:      &TreeNode{3, nil, nil},
			output: &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
		},
		{
			p: &TreeNode{7,
				&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{8, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}},
			q:      &TreeNode{0, nil, nil},
			output: nil,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := inorderSuccessor2(tt.p, tt.q)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test297(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "1,2,3,null,null,4,null,null,5,null,null",
			output: "1,2,3,null,null,4,null,null,5,null,null",
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			tree := deserialize(tt.input)
			res := serialize(tree)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test333(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output int
	}{
		{
			input: &TreeNode{10,
				&TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{8, nil, nil}},
				&TreeNode{7, &TreeNode{16, nil, nil}, &TreeNode{9, nil, nil}}},
			output: 3,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := largestBSTSubtree(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func Test426(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output int
	}{
		{
			input: &TreeNode{6,
				&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}},
			output: 3,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := largestBSTSubtree(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
