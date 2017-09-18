package table

import (
	"fmt"

	"github.com/insisthzr/happytool/container/list"
)

func hash(key Interface) int {
	if key == nil {
		return 0
	}
	h := key.HashCode()
	return h
	//return h ^ (h >> 16) //?
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
	Capacity   int
	LoadFactor float64
}

var (
	DefaultConfig = Config{
		Capacity:   1,
		LoadFactor: 0.75,
	}
)

type Table struct {
	data       []*list.List
	capacity   int
	length     int
	loadFactor float64
}

func (t *Table) String() string {
	return fmt.Sprintf("{Data: %v, Capacity: %d, Length: %d, LoadFactor: %f}",
		t.data, t.capacity, t.length, t.loadFactor)
}

func (t *Table) resize() {
	oldCap := t.capacity
	newCap := 2 * oldCap
	newData := make([]*list.List, newCap)
	for i := 0; i < oldCap; i++ {
		l := t.data[i]
		if l == nil {
			continue
		}
		l.For(func(e *list.Element) bool {
			elem := e.Value.(*element)
			newI := hash(elem.key) % newCap
			newL := newData[newI]
			if newL == nil {
				newL = list.NewList()
			}
			newL.PushBack(elem)
			newData[newI] = newL
			return true
		})
	}
	t.data = newData
	t.capacity = newCap
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
		if float64(t.length) >= float64(t.capacity)*t.loadFactor {
			t.resize()
		}
	}
}

func getFromList(l *list.List, key Interface) (interface{}, bool) {
	var value interface{}
	var ok bool
	l.For(func(e *list.Element) bool {
		elem := e.Value.(*element)
		if elem.key.Equals(key) {
			value = elem.value
			ok = true
			return false
		}
		return true
	})
	return value, ok
}

func setToList(l *list.List, key Interface, value interface{}) bool {
	var ok bool
	l.For(func(e *list.Element) bool {
		elem := e.Value.(*element)
		if elem.key.Equals(key) {
			ok = true
			elem.value = value
			return false
		}
		return true
	})
	if !ok {
		l.PushBack(&element{key: key, value: value})
	}
	return !ok
}

func NewTableWithConfig(config Config) *Table {
	if config.Capacity == 0 {
		config.Capacity = DefaultConfig.Capacity
	}
	if config.LoadFactor == 0 {
		config.LoadFactor = DefaultConfig.LoadFactor
	}
	return &Table{
		capacity:   config.Capacity,
		loadFactor: config.LoadFactor,
		data:       make([]*list.List, config.Capacity),
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
	table := NewTableWithConfig(config)
	return &IntTable{
		table: table,
	}
}

func NewIntTable() *IntTable {
	return NewIntTableWithConfig(DefaultConfig)
}
