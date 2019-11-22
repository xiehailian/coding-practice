package indexing

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBST(t *testing.T) {
	bst := NewBST()
	var arr = []int{5, 3, 6, 8, 4, 2}
	for _, v :=range(arr) {
		bst.Add(v)
	}

	bst.Print()

	bst.PreOrder()
	fmt.Println()

	bst.InOrder()
	fmt.Println()

	bst.PostOrder()
	fmt.Println()

	bst.LevelOrder()
}

func TestBST_Remove(t *testing.T) {
	bst := NewBST()
	rand.Seed(time.Now().UnixNano())

	var n = 1000
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(n))
	}

	for i := 0; i < n - 10; i++ {
		bst.Remove(i)
	}

	bst.Print()

}

func TestBitMap(t *testing.T)  {
	bitMap := NewBitMap(80)
	for i := 0; i <= 100; i += 10 {
		bitMap.Set(uint(i))
	}

	for i := 0; i <= 100; i+= 10 {
		t.Log(bitMap.Get(uint(i)))
	}

}