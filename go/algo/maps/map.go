package maps

import _ "github.com/emirpasic/gods/maps/treemap"

type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Remove(key interface{})
	Keys() []interface{}

	Size() int
	Empty() bool
}