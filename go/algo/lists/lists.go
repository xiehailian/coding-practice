package lists

type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	Insert(index int, values ...interface{})
	Set(index int, value interface{})

	Empty() bool
	Size() int
}