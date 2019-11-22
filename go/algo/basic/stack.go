package basic

type Stack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
}

type ArrayStack struct {
	array *Array
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{array: NewArray(capacity)}
}

func (as *ArrayStack) Push(value int) {
	as.array.AddLast(value)
}

func (as *ArrayStack) Pop() (int, error) {
	return as.array.RemoveLast()
}

func (as *ArrayStack) Peek() (int, error) {
	return as.array.GetLast()
}


type LinkedListStack struct {
	list *LinkedList
}

func NewLinkedListStack() *LinkedListStack{
	return &LinkedListStack{list: NewLinkedList()}
}

func (ls *LinkedListStack) Push(value int)  {
	ls.list.PushFront(value)
}


