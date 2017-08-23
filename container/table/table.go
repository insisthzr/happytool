//TODO auto resize capacity by load factor
package table

import (
	"github.com/insisthzr/happytool/container/list"
)

type Interface interface {
	HashCode() int
}

type element struct {
	key   Interface
	value interface{}
}

type Config struct {
	Capacity int
}

var (
	DefaultConfig = &Config{
		Capacity: 16,
	}
)

type Table struct {
	data     []*list.List
	capacity int
}

func (t *Table) Get(key Interface) (interface{}, bool) {
	index := key.HashCode() % t.capacity
	l := t.data[index]
	if l == nil {
		return nil, false
	}

	return getFromList(l, key)
}

func (t *Table) Set(key Interface, value interface{}) {
	index := key.HashCode() % t.capacity
	l := t.data[index]
	if l == nil {
		t.data[index] = list.NewList()
		l = t.data[index]
	}

	setToList(l, key, value)
}

func getFromList(list *list.List, key Interface) (interface{}, bool) {
	var value interface{}
	var ok bool
	list.For(func(value interface{}) bool {
		elem := value.(element)
		if elem.key == key {
			value = elem.value
			ok = true
			return false
		}
		return true
	})
	return value, ok
}

func setToList(list *list.List, key Interface, value interface{}) {
	var ok bool
	list.For(func(value interface{}) bool {
		elem := value.(element)
		if elem.key == key {
			ok = true
			return false
		}
		return true
	})
	if !ok {
		list.PushBack(&element{key: key, value: value})
	}
}

func NewTableWithConfig(config *Config) *Table {
	if config.Capacity == 0 {
		config.Capacity = DefaultConfig.Capacity
	}
	return &Table{
		capacity: config.Capacity,
		data:     make([]*list.List, config.Capacity),
	}
}

func NewTable() *Table {
	return NewTableWithConfig(DefaultConfig)
}

type Int int

func (i Int) HashCode() int {
	return int(i)
}

type IntTable struct {
	table *Table
}

func (t *IntTable) Get(key int) (interface{}, bool) {
	value, ok := t.table.Get(Int(key))
	return value.(int), ok
}

func (t *IntTable) Set(key int, value interface{}) {
	t.table.Set(Int(key), value)
}
