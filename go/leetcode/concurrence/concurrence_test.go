package concurrence

import (
	"fmt"
	"sync"
	"testing"
)

func Test1114(t *testing.T)  {
	foo := NewFoo()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		foo.First()
	}()

	go func() {
		defer wg.Done()
		foo.Two()
	}()

	go func() {
		defer wg.Done()
		foo.Three()
	}()

	wg.Wait()
}

func Test1115(t *testing.T) {
	fb := NewForBar(5)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fb.Foo()
	}()
	go func() {
		defer wg.Done()
		fb.Bar()
	}()
	wg.Wait()
}

func Test1116(t *testing.T) {
	n := 10
	zeo := NewZeroEvenOdd(n)
	done := make(chan bool, 1)
	printNunmber := func (n int) {
		fmt.Printf("%d", n)
	}

	go zeo.zero(printNunmber)
	go func() {
		zeo.even(printNunmber)
		if n % 2 == 0 {
			done <- true
		}
	}()
	go func() {
		zeo.odd(printNunmber)
		if n % 2 != 0 {
			done <- true
		}
	}()

	<- done
}

func Test1117(t *testing.T) {
	input := "OOOOHHHHHHHH"
	h2o := NewH2O()
	output := ""
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	putch := func(char int) {
		mutex.Lock()
		output = output + string(char)
		mutex.Unlock()
		wg.Done()
	}
	wg.Add(len(input))
	for _, ch := range input {
		if ch == 'H' {
			go h2o.Hydrogen(func() { putch('H')})
		} else {
			go h2o.Oxygen(func() { putch('O')})
		}
	}
	wg.Wait()

	fmt.Println("input: ", input)
	fmt.Println("output: ", output)
}

func Test1195(t *testing.T) {
	n := 15
	fb := NewFizzBuzz(n)
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go func() {
		defer wg.Done()
		fb.Fizz(func() { fmt.Print("fizz,")})
	}()
	go func() {
		defer wg.Done()
		fb.Buzz(func() { fmt.Print("buzz,")})
	}()
	go func() {
		defer wg.Done()
		fb.FizzBuzz(func() { fmt.Print("fizzbuzz,")})
	}()
	go func() {
		defer wg.Done()
		fb.Number(func(n int) { fmt.Printf("%d,", n)})
	}()
	wg.Wait()
}

