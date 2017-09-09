package table

import (
	"github.com/insisthzr/happytool/container/list"
)

func hash(key Interface) int {
	if key == nil {
		return 0
	}
	h := key.HashCode()
	return h ^ (h >> 16) //?
}

type Interface interface {
	HashCode() int
	Equals(Interface) bool
}

type element struct {
	key   Interface
	value interface{}
}

type Config struct {
	Capacity int
}

var (
	DefaultConfig = Config{
		Capacity: 16,
	}
)

type Table struct {
	data     []*list.List
	capacity int
	length   int
}

func (t *Table) Get(key Interface) (interface{}, bool) {
	h := hash(key)
	index := h % t.capacity
	l := t.data[index]
	if l == nil {
		return nil, false
	}
	return getFromList(l, key)
}

func (t *Table) Set(key Interface, value interface{}) {
	h := hash(key)
	index := h % t.capacity
	l := t.data[index]
	if l == nil {
		t.data[index] = list.NewList()
		l = t.data[index]
	}
	new := setToList(l, key, value)
	if new {
		t.length++
	}
}

func getFromList(l *list.List, key Interface) (interface{}, bool) {
	var value interface{}
	var ok bool
	l.For(func(val interface{}) bool {
		elem := val.(*element)
		if elem.key.Equals(key) {
			value = elem.value
			ok = true
			return false
		}
		return true
	})
	return value, ok
}

func setToList(list *list.List, key Interface, value interface{}) bool {
	var ok bool
	list.For(func(val interface{}) bool {
		elem := val.(*element)
		if elem.key.Equals(key) {
			ok = true
			elem.value = value
			return false
		}
		return true
	})
	if !ok {
		list.PushBack(&element{key: key, value: value})
	}
	return !ok
}

func NewTableWithConfig(config Config) *Table {
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

func (i Int) Equals(other Interface) bool {
	j, ok := other.(Int)
	if !ok {
		return false
	}
	return int(i) == int(j)
}

type IntTable struct {
	table *Table
}

func (t *IntTable) Get(key int) (interface{}, bool) {
	value, ok := t.table.Get(Int(key))
	return value, ok
}

func (t *IntTable) Set(key int, value interface{}) {
	t.table.Set(Int(key), value)
}

func NewIntTableWithConfig(config Config) *IntTable {
	if config.Capacity == 0 {
		config.Capacity = DefaultConfig.Capacity
	}
	table := &Table{
		capacity: config.Capacity,
		data:     make([]*list.List, config.Capacity),
	}
	return &IntTable{
		table: table,
	}
}

func NewIntTable() *IntTable {
	return NewIntTableWithConfig(DefaultConfig)
}
