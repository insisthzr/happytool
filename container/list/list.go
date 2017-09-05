package list

type Element struct {
	prev  *Element
	next  *Element
	Value interface{}
	list  *List
}

func (e *Element) Next() *Element {
	if e.list == nil {
		return nil
	}
	if e.next == e.list.root {
		return nil
	}
	return e.next
}

func (e *Element) Prev() *Element {
	if e.list == nil {
		return nil
	}
	if e.prev == e.list.root {
		return nil
	}
	return e.prev
}

type List struct {
	root *Element
	len  int
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) insertAfter(elem *Element, at *Element) *Element {
	if at.list != l {
		return nil
	}

	n := at.next
	at.next = elem
	elem.next = n
	elem.prev = at
	n.prev = elem

	l.len++
	elem.list = l

	return elem
}

func (l *List) insertBefore(elem *Element, at *Element) *Element {
	return l.insertAfter(elem, at.prev)
}

func (l *List) InsertAfter(value interface{}, at *Element) *Element {
	elem := &Element{Value: value}
	return l.insertAfter(elem, at)
}

func (l *List) InsertBefore(value interface{}, at *Element) *Element {
	elem := &Element{Value: value}
	return l.insertBefore(elem, at)
}

func (l *List) PushBack(value interface{}) *Element {
	return l.InsertAfter(value, l.root.prev)
}

func (l *List) PushFront(value interface{}) *Element {
	return l.InsertAfter(value, l.root)
}

func (l *List) Remove(elem *Element) *Element {
	if elem.list != l {
		return nil
	}

	elem.prev.next = elem.next
	elem.next.prev = elem.prev
	elem.prev = nil
	elem.next = nil
	elem.list = nil

	l.len--

	return elem
}

func (l *List) PopFront() (interface{}, bool) {
	if l.len == 0 {
		return nil, false
	}
	return l.Remove(l.root.next).Value, true
}

func (l *List) PopBack() (interface{}, bool) {
	if l.len == 0 {
		return nil, false
	}
	return l.Remove(l.root.prev).Value, true
}

func (l *List) For(fn func(value interface{}) bool) {
	if l.Len() == 0 {
		return
	}
	for elem := l.root.next; elem != nil; elem = elem.Next() {
		conti := fn(elem.Value)
		if !conti {
			break
		}
	}
}

func NewList() *List {
	list := &List{}
	root := &Element{}
	root.next = root
	root.prev = root
	root.list = list
	list.root = root
	return list
}
