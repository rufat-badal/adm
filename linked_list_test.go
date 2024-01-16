package adm

import (
	"fmt"
	"testing"
)

type sliceTestData[T comparable] struct {
	slice        []T
	notContained []T
}

var intSliceTests = []sliceTestData[int]{
	{[]int{1, 2, 3, 4}, []int{-10, -4, 5, 7}},
	{[]int{1, 1, 1, 1, 1}, []int{2, 3, 4, 5, -10}},
	{[]int{}, []int{1, 2, 3, 4, 5, -10}},
	{[]int{-42}, []int{23, 25, 67, 1992}},
}

var float64SliceTests = []sliceTestData[float64]{
	{[]float64{1.0, 2.3, 3.1, 7.5}, []float64{1.1, 1.0001, 3.2, 200.3}},
	{[]float64{1.0, 1.0, 1.0, 1.0, 1.0}, []float64{1.1, 1.0001, 3.2, 200.3}},
	{[]float64{}, []float64{1.1, 1.0001, 3.2, 200.3}},
	{[]float64{-42.42}, []float64{-42.0, -44.0, 33.7834}},
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

func testLinkedListFromSlice[T comparable](t *testing.T, data []sliceTestData[T]) {
	for _, e := range data {
		testLinkedListFromSingleSlice[T](t, e.slice)
	}
}

func TestLinkedListFromSlice(t *testing.T) {
	testLinkedListFromSlice[int](t, intSliceTests)
	testLinkedListFromSlice[float64](t, float64SliceTests)
}

func testSearchSingleData[T comparable](t *testing.T, data sliceTestData[T]) {
	if len(data.slice) == 0 {
		return
	}

	l := LinkedListFromSlice[T](data.slice)

	for _, x := range data.slice {
		errorStart := fmt.Sprintf("Searching for %v in linked list (should be present):", x)
		found := Search(l, x)
		if found == nil {
			t.Errorf("%q element was not found", errorStart)
			continue
		}
		if found.Value != x {
			t.Errorf("%q found wrong element with wrong Value = %v", errorStart, found.Value)
		}
	}
}

func testSearch[T comparable](t *testing.T, data []sliceTestData[T]) {
	for _, e := range data {
		testSearchSingleData[T](t, e)
	}
}

func TestSearch(t *testing.T) {
	testSearch[int](t, intSliceTests)
	testSearch[float64](t, float64SliceTests)
}

func testItemAheadSingleData[T comparable](t *testing.T, data sliceTestData[T]) {
	l := LinkedListFromSlice[T](data.slice)

	for itemAheadWant := l; itemAheadWant != nil; itemAheadWant = itemAheadWant.Next {
		node := itemAheadWant.Next
		errorStart := fmt.Sprintf("Searching for item ahead of %v in linked list:", node)
		itemAhead := ItemAhead(l, node)
		if itemAhead == nil {
			t.Errorf("%q element was not found", errorStart)
			continue
		}
		if itemAhead.Next != node {
			t.Errorf("%q found element %v that is not followed by %v", errorStart, itemAhead, node)
		}
	}
}

func testItemAhead[T comparable](t *testing.T, data []sliceTestData[T]) {
	for _, e := range data {
		testItemAheadSingleData[T](t, e)
	}
}

func TestItemAhead(t *testing.T) {
	testItemAhead[int](t, intSliceTests)
	testItemAhead[float64](t, float64SliceTests)
}
