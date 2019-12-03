package design

type CircularDeque struct {
	data []int
	front int
	last int

}

func NewCircularDeque(k int) CircularDeque {
	return CircularDeque{
		data: make([]int, k+1, k+1),
		front: 0,
		last: 0,
	}
}

func (this *CircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.front = (len(this.data) + this.front - 1) % len(this.data)
	this.data[this.front] = value

	return true
}

func (this *CircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.data[this.last] = value
	this.last = (this.last + 1) % len(this.data)
	return true
}

func (this *CircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1) % len(this.data)
	return true
}

func (this *CircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.last = (len(this.data) + this.last - 1) % len(this.data)
	return true
}

func (this *CircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[this.front]
}

func (this *CircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[(len(this.data) + this.last - 1) % len(this.data)]
}

func (this *CircularDeque) IsEmpty() bool {
	return this.last == this.front
}

func (this *CircularDeque) IsFull() bool {
	return (this.last + 1) % len(this.data) == this.front
}
