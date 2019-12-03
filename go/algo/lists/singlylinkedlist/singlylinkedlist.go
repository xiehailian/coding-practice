package singlylinkedlist

type element struct {
	value interface{}
	next *element
}

type List struct {
	first *element
	last *element
	size int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{})  {
	for _, value := range values {
		cur := &element{value: value}
		if list.size == 0 {
			list.first = cur
			list.last = cur
		} else {
			list.last.next = cur
			list.last = cur
		}
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	cur := list.first
	for i := 0; i != index; i++ {
		cur = cur.next
	}

	return cur.value, true
}

func (list *List) Remove(index int) {

}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}