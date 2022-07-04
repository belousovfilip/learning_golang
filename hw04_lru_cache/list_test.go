package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("removeLast", func(t *testing.T) {
		l := NewList()
		l.PushBack(20)
		l.PushFront(10)
		l.PushBack(30)
		length := l.Len()
		l.Remove(l.Back())
		require.Equal(t, length-1, l.Len())
		require.Equal(t, 20, l.Back().Value)
		require.Nil(t, l.Back().Next)
	})

	t.Run("removeFirst", func(t *testing.T) {
		l := NewList()
		l.PushBack(20)
		l.PushFront(10)
		l.PushBack(30)
		length := l.Len()
		l.Remove(l.Front())
		require.Equal(t, length-1, l.Len())
		require.Equal(t, 20, l.Front().Value)
		require.Nil(t, l.Front().Prev)
	})

	t.Run("removeMiddle", func(t *testing.T) {
		l := NewList()
		l.PushBack(20)
		l.PushFront(10)
		l.PushBack(30)
		length := l.Len()
		l.Remove(l.Front().Next)
		require.Equal(t, length-1, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()
		l.PushFront(10) // [10]
		l.PushBack(30)  // [10, 20, 30]
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
