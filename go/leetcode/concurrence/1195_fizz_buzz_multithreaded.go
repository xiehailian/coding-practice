package concurrence

//编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，但是：
//
//如果这个数字可以被 3 整除，输出 "fizz"。
//如果这个数字可以被 5 整除，输出 "buzz"。
//如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。
//例如，当 n = 15，输出： 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/fizz-buzz-multithreaded

type FizzBuzz struct {
	n int
	fizzChan chan int
	buzzChan chan int
	fizzBuzzChan chan int
	numberChan chan int
}

func NewFizzBuzz(n int) *FizzBuzz {
	fb := new(FizzBuzz)
	fb.n = n
	fb.fizzChan = make(chan int, 1)
	fb.buzzChan = make(chan int, 1)
	fb.fizzBuzzChan = make(chan int, 1)
	fb.numberChan = make(chan int, 1)
	fb.numberChan <- 1
	return fb
}

func (fb *FizzBuzz) dispatch(n int) {
	if n % 15 == 0 {
		fb.fizzBuzzChan <- n
	} else if n % 3 == 0 {
		fb.fizzChan <- n
	} else if n % 5 == 0 {
		fb.buzzChan <- n
	} else {
		fb.numberChan <- n
	}
}

func (fb *FizzBuzz) Fizz(printFizz func()) {
	for x := range fb.fizzChan {
		if x > fb.n {
			close(fb.buzzChan)
			close(fb.fizzChan)
			close(fb.numberChan)
			break
		}
		printFizz()
		fb.dispatch(x + 1)
	}
}

func (fb *FizzBuzz) Buzz(printBuzz func()) {
	for x := range fb.buzzChan {
		if x > fb.n {
			close(fb.fizzChan)
			close(fb.fizzChan)
			close(fb.numberChan)
			break
		}

		printBuzz()
		fb.dispatch(x + 1)
	}
}

func (fb *FizzBuzz) FizzBuzz(printFizzBuzz func())  {
	for x := range fb.fizzBuzzChan {
		if x > fb.n {
			close(fb.fizzChan)
			close(fb.buzzChan)
			close(fb.fizzBuzzChan)
			break
		}
		printFizzBuzz()
		fb.dispatch(x + 1)
	}
}

func (fb *FizzBuzz) Number(printNumber func(int)) {
	for x := range fb.numberChan {
		if x > fb.n {
			close(fb.fizzChan)
			close(fb.buzzChan)
			close(fb.fizzBuzzChan)
			break
		}

		printNumber(x)
		fb.dispatch(x + 1)
	}
}