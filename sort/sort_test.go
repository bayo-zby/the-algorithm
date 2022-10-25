package sort_test

import (
	"algorithm/sort"
	"reflect"
	"testing"
)

var testTemps = []struct {
	input  []int
	expect []int
	name   string
}{
	//Sorted slice
	{
		input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:   "Sorted Unsigned",
	},
	//Reversed slice
	{
		input:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:   "Reversed Unsigned",
	},
	//Sorted slice
	{
		input:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expect: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:   "Sorted Signed",
	},
	//Reversed slice
	{
		input:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		expect: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:   "Reversed Signed",
	},
	//Reversed slice, even length
	{
		input:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		expect: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:   "Reversed Signed #2",
	},
	//Random order with repetitions
	{
		input:  []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
		expect: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
		name:   "Random order Signed",
	},
	//Single-entry slice
	{
		input:  []int{1},
		expect: []int{1},
		name:   "Singleton",
	},
	// Empty slice
	{
		input:  []int{},
		expect: []int{},
		name:   "Empty Slice",
	},
	// nil slice
	{
		input:  nil,
		expect: nil,
		name:   "nil Slice",
	},
}

func testFramework(t *testing.T, f func([]int) []int) {
	for _, test := range testTemps {
		if actual := f(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("[%v]expect %v,actual %v", test.name, test.expect, actual)
		}
	}
}

func benckmarkFramework(b *testing.B, f func([]int) []int) {
	for i := 0; i < b.N; i++ {
		for _, test := range testTemps {
			f(test.input)
		}
	}
}

// test
func TestSelectSort(t *testing.T) {
	testFramework(t, sort.SelectSort)
}

func TestBubbleSort(t *testing.T) {
	testFramework(t, sort.BubbleSort)
}

func TestInsertionSort(t *testing.T) {
	testFramework(t, sort.InsertionSort)
}

// benchmark
func BenchmarkSelectSort(b *testing.B) {
	benckmarkFramework(b, sort.SelectSort)
}

func BenchmarkBubbleSort(b *testing.B) {
	benckmarkFramework(b, sort.BubbleSort)
}

func BenchmarkInsertionSort(b *testing.B) {
	benckmarkFramework(b, sort.InsertionSort)
}
