package design

import "math/rand"

const MAX_LEVEL = 16

type SkiplistNode struct {
	value    int
	level    int
	forwards []*SkiplistNode
}

func NewSkiplistNode(value, level int) *SkiplistNode {
	return &SkiplistNode{value, level, make([]*SkiplistNode, level, level)}
}

type Skiplist struct {
	head   *SkiplistNode
	level  int
	length int
}

func NewSkiplist() Skiplist {
	head := NewSkiplistNode(0, MAX_LEVEL)
	return Skiplist{head, 1, 0}
}

func (this *Skiplist) Search(target int) bool {
	if this.length == 0 {
		return false
	}

	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].value == target {
				return true
			} else if cur.forwards[i].value > target {
				break
			}
			cur = cur.forwards[i]
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {

	cur := this.head
	update := [MAX_LEVEL]*SkiplistNode{}
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].value == num {
				return
			}
			if cur.forwards[i].value > num {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	new := NewSkiplistNode(num, level)
	for i := 0; i < level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = new
		new.forwards[i] = next
	}

	if level > this.level {
		this.level = level
	}

	this.length++
}

func (this *Skiplist) Erase(num int) bool {

	if !this.Search(num) {
		return false
	}

	cur := this.head
	update := [MAX_LEVEL]*SkiplistNode{}
	for i := this.level - 1; i >= 0; i-- {
		update[i] = this.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].value == num {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == this.head && cur.forwards[i] == nil {
			this.level = i
		}

		if update[i].forwards[i] == nil {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}
	this.length--
	return true
}
