package concurrence

import "fmt"

//我们提供一个类：
//
//class FooBar {
//  public void foo() {
//    for (int i = 0; i < n; i++) {
//      print("foo");
//    }
//  }
//
//  public void bar() {
//    for (int i = 0; i < n; i++) {
//      print("bar");
//    }
//  }
//}
//两个不同的线程将会共用一个 FooBar 实例。其中一个线程将会调用 foo() 方法，另一个线程将会调用 bar() 方法。
//
//请设计修改程序，以确保 "foobar" 被输出 n 次。
//
//示例 1:
//
//输入: n = 1
//输出: "foobar"
//解释: 这里有两个线程被异步启动。其中一个调用 foo() 方法, 另一个调用 bar() 方法，"foobar" 将被输出一次。
//示例 2:
//
//输入: n = 2
//输出: "foobarfoobar"
//解释: "foobar" 将被输出两次。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/print-foobar-alternately

type FooBar struct {
	n int

	fooBlock chan bool
	barBlock chan bool
}

func NewForBar(n int) *FooBar {
	return &FooBar{n : n, fooBlock: make(chan bool), barBlock: make(chan bool)}
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.n; i++ {
		fmt.Print("foo")
		fb.barBlock <- true
		<- fb.fooBlock
	}
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.n; i++ {
		<- fb.barBlock
		fmt.Print("bar")
		fb.fooBlock <- true

	}
}