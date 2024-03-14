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

	t.Run("complex 2", func(t *testing.T) {
		l := NewList()

		l.PushFront(100) // [100]
		l.PushBack(200)  // [100, 200]
		l.PushBack(300)  // [100, 200, 300]
		l.PushBack(400)  // [100, 200, 300, 400]
		require.Equal(t, 4, l.Len())

		second := l.Front().Next // 200
		require.Equal(t, 200, second.Value)

		third := l.Back().Prev // 300
		require.Equal(t, 300, third.Value)

		l.Remove(second) // [100, 300, 400]
		require.Equal(t, 3, l.Len())

		l.Remove(third) // [100, 400]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{500, 600, 700, 800, 900} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [900, 700, 500, 100, 400, 600, 800]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 900, l.Front().Value)
		require.Equal(t, 800, l.Back().Value)

		l.MoveToFront(l.Front()) // [900, 700, 500, 100, 400, 600, 800]
		l.MoveToFront(l.Back())  // [800, 900, 700, 500, 100, 400, 600]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{800, 900, 700, 500, 100, 400, 600}, elems)
	})

	t.Run("complex 3", func(t *testing.T) {
		l := NewList()

		l.PushBack(200)  // [200]
		l.PushBack(300)  // [200, 300]
		l.PushFront(400) // [400, 200, 300]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 200
		require.Equal(t, 200, middle.Value)

		middle = l.Back().Prev // 200
		require.Equal(t, 200, middle.Value)

		l.Remove(middle) // [400, 300]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{500, 600, 700, 800} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [700, 500, 400, 300, 600, 800]

		require.Equal(t, 6, l.Len())
		require.Equal(t, 700, l.Front().Value)
		require.Equal(t, 800, l.Back().Value)

		l.MoveToFront(l.Front()) // [700, 500, 400, 300, 600, 800]
		l.MoveToFront(l.Back())  // [800, 700, 500, 400, 300, 600]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{800, 700, 500, 400, 300, 600}, elems)
	})
}
