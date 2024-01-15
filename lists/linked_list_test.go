package lists

import (
	"fmt"
	"testing"
)

var linkedListFromIntSliceTests = [][]int{
	{1, 2, 3, 4},
	{1, 1, 1, 1, 1},
	{},
	{-42},
}

var linkedListFromFloat64SliceTests = [][]float64{
	{1.0, 2.3, 3.1, 7.5},
	{1.0, 1.0, 1.0, 1.0, 1.0},
	{},
	{-42.42},
}

func testLinkedListFromSingleSlice[T comparable](t *testing.T, s []T) {
	errorStart := fmt.Sprintf("LinkedListFromSlice(%v) ", s)

	cur := LinkedListFromSlice[T](s)

	if len(s) == 0 && cur != nil {
		t.Errorf("%q must return <nil>", errorStart)
	}

	for i := 0; i < len(s); i++ {
		if cur == nil {
			t.Errorf("%q returned a LinkedList with missing element %v", errorStart, i)
			break
		}
		if cur.Value != s[i] {
			t.Errorf("%q returned a LinkedList with %v (element %v) != s[%v] = %v", errorStart, cur.Value, i, i, s[i])
		}
		cur = cur.Next
	}
}

func testLinkedListFromSlice[T comparable](t *testing.T, slices [][]T) {
	for _, s := range slices {
		testLinkedListFromSingleSlice[T](t, s)
	}
}

func TestLinkedListFromSlice(t *testing.T) {
	testLinkedListFromSlice[int](t, linkedListFromIntSliceTests)
	testLinkedListFromSlice[float64](t, linkedListFromFloat64SliceTests)
}
