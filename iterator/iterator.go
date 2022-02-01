// file: iterator.go
// author: TaoZer
// Date: 2022/2/1
// Description: iterator pattern

package iterator

import "container/list"

type IteratorI interface {
	HasNext() bool
	Next() interface{}
}

type SliceIterator struct {
	pos int
	slice []interface{}
}

func NewSliceIterator(slice []interface{}) *SliceIterator {
	return &SliceIterator{
		slice: slice,
		pos: 0,
	}
}

func (s *SliceIterator) HasNext() bool {
	return s.pos < len(s.slice)
}

func (s *SliceIterator) Next() interface{} {
	v := s.slice[s.pos]
	s.pos++
	return v
}

type MapIterator struct {
	pos int
	vSlice []interface{}
}

func NewMapIterator(m map[interface{}]interface{}) *MapIterator {
	var slice []interface{}
	for _, v := range m {
		slice = append(slice, v)
	}
	return &MapIterator{
		vSlice: slice,
		pos: 0,
	}
}

func (m *MapIterator) HasNext() bool {
	return m.pos < len(m.vSlice)
}

func (m *MapIterator) Next() interface{} {
	v := m.vSlice[m.pos]
	m.pos++
	return v
}

type ListIterator struct {
	l *list.List
	pos int
	p *list.Element
}

func NewListIterator(list *list.List) *ListIterator {
	return &ListIterator{
		l: list,
		pos: 0,
		p: list.Front(),
	}
}

func (l *ListIterator) HasNext() bool {
	return l.pos < l.l.Len()
}

func (l *ListIterator) Next() interface{} {
	v := l.p.Value
	l.pos++
	l.p = l.p.Next()
	return v
}