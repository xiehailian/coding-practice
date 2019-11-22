package basic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/algo/basic"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T)  {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			s := GetLazySingletonInstance()
			fmt.Printf("instance: %p\n", s)
			wg.Done()
		}()
	}
	wg.Wait()
}


func TestDuplicationInArray(t *testing.T) {
	tests := []struct{
		input []int
		output int
	}{
		{
			input:	[]int{1, 2, 3, 5, 4, 6, 2},
			output: 2,
		},
		{
			input:	[]int{},
			output:	-1,
		},
		{
			input: 	nil,
			output:	-1,
		},
	}

	for _, tt := range tests{
		t.Run("test", func(t *testing.T) {
			res := DuplicationInArray(tt.input)
			assert.Equal(t, tt.output, res)
			res = DuplicationInArrayNoEdit(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestFindInPartiallySortedMatrix(t *testing.T) {
	tests := []struct{
		matrix 	[][]int
		target	int
		output	bool
	}{
		{
			matrix: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
			target: 6,
			output: true,
		},
		{
			matrix: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
			target: 7,
			output: true,
		},
		{
			matrix: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
			target: 0,
			output: false,
		},
		{
			matrix: nil,
			target: 16,
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			res := FindInPartiallySortedMatrix(tt.matrix, tt.target)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestReplaceSpaces(t *testing.T) {
	tests := []struct{
		input string
		output string
	}{
		{
			input: "hello world",
			output: "hello%20world",
		},
		{
			input: "",
			output: "",
		},
		{
			input: " ",
			output: "%20",
		},
	}

	for _, tt := range(tests) {
		t.Run("test", func(t *testing.T) {
			res := ReplaceSpaces(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestPrintListInReversedOrder(t *testing.T) {
	l := basic.NewLinkedList()
	e1 := l.PushFront(1)
	l.InsertAfter(2, e1)
	l.InsertAfter(3, e1.Next())
	l.InsertAfter(4, e1.Next().Next())
	l.InsertAfter(5, e1.Next().Next().Next())
	fmt.Println("顺序打印：")
	l.Print()
	fmt.Println("逆序打印：")
	PrintListInReversedOrder(e1)
	fmt.Println()
}
