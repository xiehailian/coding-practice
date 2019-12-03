package sets

type Set interface {
	Add(values ...interface{})
	Remove(values ...interface{})
	Contains(values ...interface{}) bool

	Size() int
	Empty() bool
}