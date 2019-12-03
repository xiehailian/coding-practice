package design

type MinStack struct {
	datas []int
	mins  []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		datas: make([]int, 0),
		mins:  make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.datas = append(this.datas, x)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, x)
	} else {
		if this.mins[len(this.mins)-1] > x {
			this.mins = append(this.mins, x)
		}
	}
}

func (this *MinStack) Pop() {
	data := this.datas[len(this.datas)-1]
	this.datas = this.datas[:len(this.datas)-1]
	if data == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.datas[len(this.datas)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
