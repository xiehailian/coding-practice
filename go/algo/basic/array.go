package basic

import (
	"errors"
	"fmt"
	"strconv"
)

type Array struct {
	data []int
	size int
}

func NewArray(capacity int) *Array {
	return &Array{data : make([]int, capacity, capacity), size: 0}
}

// 查
func (a *Array) Get(index int) (int, error){
	if index < 0 || index >= a.size {
		return -1, errors.New("get failed")
	}
	return a.data[index], nil
}

func (a *Array) GetLast() (int, error) {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() (int, error) {
	return a.Get(0)
}

// 改
func (a *Array) Set(index int, value int) error{
	if index < 0 || index >= a.size {
		return errors.New("set failed")
	}
	a.data[index] = value
	return nil
}

// 增
func (a *Array) Add(index int, value int) error {

	if index < 0 || index > a.size {
		return errors.New("Add failed")
	}

	if a.size == cap(a.data) {
		a.resize(2 * cap(a.data))
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = value
	a.size++
	return nil
}

func (a *Array) AddLast(value int) {
	a.Add(a.size, value)
}

func (a *Array) AddFirst(value int)  {
	a.Add(0, value)
}

func (a *Array) Contains(value int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == value {
			return true
		}
	}
	return false
}


// 删
func (a *Array) Remove(index int) (int, error) {

	if index < 0 || index >= a.size {
		return -1, errors.New("remove failed")
	}

	ret := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	a.size--
	a.data[a.size-1] = 0

	if a.size == cap(a.data)/4 && cap(a.data)/2 != 0 {
		a.resize(cap(a.data) / 2)
	}
	return ret, nil
}

func (a *Array) RemoveFirst() (int, error) {
	return a.Remove(0)
}

func (a *Array) RemoveLast() (int, error) {
	return a.Remove(a.size - 1)
}

func (a *Array) resize(capacity int) {
	var newData = make([]int, capacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}


func (a *Array) Print() {
	var res = "["
	for i := 0; i < a.size; i++ {
		res = res + strconv.Itoa(a.data[i])
		if i != a.size - 1 {
			res = res + ", "
		}
	}
	res = res + "]"
	fmt.Println(res)
}