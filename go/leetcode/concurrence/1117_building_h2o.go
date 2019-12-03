package concurrence

import (
	"sync"
)

//你必须保证产生一个水分子所需线程的结合必须发生在下一个水分子产生之前。
//
//换句话说:
//
//如果一个氧线程到达屏障时没有氢线程到达，它必须等候直到两个氢线程到达。
//如果一个氢线程到达屏障时没有其它线程到达，它必须等候直到一个氧线程和另一个氢线程到达。
//书写满足这些限制条件的氢、氧线程同步代码。
//
// 
//
//示例 1:
//
//输入: "HOH"
//输出: "HHO"
//解释: "HOH" 和 "OHH" 依然都是有效解。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/building-h2o


type H2O struct {
	H int
	O int
	m *sync.Mutex
	c *sync.Cond
}

func NewH2O() *H2O {
	h2o := new(H2O)
	h2o.H = 0
	h2o.O = 0
	h2o.m = &sync.Mutex{}
	h2o.c = sync.NewCond(h2o.m)
	return h2o
}

func (self *H2O) Hydrogen(releaseHydrogen func()) {
	self.m.Lock()
	for self.H >= 2 {
		self.c.Wait()
	}

	releaseHydrogen()
	self.H += 1
	if self.H == 2 && self.O == 1 {
		self.H = 0
		self.O = 0
		self.c.Broadcast()
	}
	self.m.Unlock()
}

func (self *H2O) Oxygen(releaseOxygen func()) {
	self.m.Lock()
	for self.O >= 1 {
		self.c.Wait()
	}

	releaseOxygen()

	self.O += 1
	if self.H == 2 && self.O == 1 {
		self.H = 0
		self.O = 0
		self.c.Broadcast()
	}
	self.m.Unlock()
}