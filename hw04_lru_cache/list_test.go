package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	value int
}

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("add one", func(t *testing.T) {
		l := NewList()
		ts := testStruct{1}
		li := l.PushBack(ts)
		require.Equal(t, 1, l.Len())
		require.Equal(t, li, l.Front())
		require.Equal(t, li, l.Back())
	})

	t.Run("add one, remove", func(t *testing.T) {
		l := NewList()
		ts := testStruct{0}
		item := l.PushBack(ts)
		require.Equal(t, item, l.Front())
		require.Equal(t, item, l.Back())
		l.Remove(item)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("add two, remove", func(t *testing.T) {
		l := NewList()
		item1 := l.PushBack(testStruct{1})
		item2 := l.PushBack(testStruct{2})
		require.Equal(t, 2, l.Len())
		require.Equal(t, item1, l.Front())
		require.Equal(t, item2, l.Back())

		l.Remove(item1)
		require.Equal(t, 1, l.Len())
		require.Equal(t, item2, l.Front())
		require.Equal(t, item2, l.Back())

		l.InsertBefore(item1, item2)
		require.Equal(t, 2, l.Len())
		require.Equal(t, item1, l.Front())
		require.Equal(t, item2, l.Back())
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()
		item1 := l.PushBack(testStruct{1})
		item2 := l.PushBack(testStruct{2})

		l.MoveToFront(item2)
		require.Equal(t, 2, l.Len())
		require.Equal(t, item2, l.Front())
		require.Equal(t, item1, l.Back())
	})

	t.Run("push front", func(t *testing.T) {
		l := NewList()
		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		require.Equal(t, l.Front().Value.(int), 3)
		require.Equal(t, l.Front().Next.Value.(int), 2)
		require.Equal(t, l.Back().Value.(int), 1)
		require.Equal(t, l.Back().Prev.Value.(int), 2)
	})

	t.Run("pop back", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)
		l.PopBack()
		require.Equal(t, 2, l.Len())
		require.Equal(t, l.Back().Value.(int), 2)
		l.PopBack()
		l.PopBack()
		require.Equal(t, 0, l.Len())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()
		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
