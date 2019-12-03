package arraylist

import (
	"fmt"
	"strings"
)

type List struct {
	elements []interface{}
	size int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}


func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

func (list *List) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--

	list.shrink()
}

func (list *List) Contains(values ...interface{}) bool {

	for _, value := range values {
		found := false
		for _, element := range list.elements {
			if element == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Insert(index int, values ...interface{}) {

	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

func (list *List) Set(index int, value interface{})  {

	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

func (list *List) Empty() bool {
	return true
}

func (list *List) Size() int {
	return list.size
}

func (list *List) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *List) withinRange(index int) bool  {
	return index >= 0 && index < list.size
}


func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity{
		newCapcity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapcity)
	}
}

func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}


