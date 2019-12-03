package arraylist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := New()
	assert.Equal(t, true, list.Empty())

	list.Add("a", "b", "c")
	assert.Equal(t, 3, list.Size())

	list.Remove(1)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, false, list.Contains("b"))

	list.Remove(0)
	list.Remove(0)
	list.Remove(0)
	fmt.Println(list.String())
}

func benchmarkGet(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Get(n)
		}
	}
}

func benchmarkAdd(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Remove(n)
		}
	}
}

func BenchmarkArrayListGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkArrayListRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkArrayListRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkArrayListRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}
