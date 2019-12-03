package queues

type Queue interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Top() (value interface{}, ok bool)

	Size() int
	Empty() bool
}

//https://leetcode-cn.com/problems/design-circular-deque/solution/she-ji-xun-huan-shuang-duan-dui-lie-tui-dao-yu-shi/
// 顺序队列（基于数组，每次取数据需要进行数据搬迁，优化是等位置不够了再一次搬迁）、
// 链式队列（容易实现，无限扩展，但不太常用）、
// 循环队列（基于数组，head/tail循环移动，多占用一个数据空间）、
// 阻塞队列（取空队列头部阻塞至有值入队）、
// 并发队列(最简单的就是入队出队操作加锁)

// 队列的应用
// 线程池等资源池，通常有最大数量限制，不适合链式队列
// 银行等业务排队系统
