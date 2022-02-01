package iterator

import (
	"container/list"
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	slice := []int{1,2,3,4,5,6}
	iSlice := make([]interface{}, len(slice))
	for i := 0; i < len(slice); i++ {
		iSlice[i] = slice[i]
	}
	sliceIter := NewSliceIterator(iSlice)
	m := make(map[interface{}]interface{})
	for i := 7; i < 12; i++ {
		m[i] = i
	}
	mIter := NewMapIterator(m)
	l := makeList(iSlice)
	lIte := NewListIterator(l)

	iterate(sliceIter)
	iterate(mIter)
	iterate(lIte)
}

func iterate(i IteratorI) {
	for i.HasNext() {
		fmt.Println(i.Next())
	}
}

func makeList(s []interface{}) *list.List {
	l := list.New()
	for i := 0; i < len(s); i++ {
		l.PushBack(s[i])
	}
	return l
}

