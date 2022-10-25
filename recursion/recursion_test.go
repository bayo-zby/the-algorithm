package recursion_test

import (
	"algorithm/recursion"
	"testing"
)

type tempType struct {
	input  [2]int
	expect int
	name   string
}

var testTemps = []tempType{
	{
		input:  [2]int{1, 2},
		expect: 1,
		name:   "simple",
	},
	{
		input:  [2]int{1, 1},
		expect: 1,
		name:   "same",
	},
	{
		input:  [2]int{0, 1},
		expect: 1,
		name:   "zero first",
	},
	{
		input:  [2]int{1, 0},
		expect: 1,
		name:   "zero last",
	},
}

func testFramework(t *testing.T, f func(int, int) int) {
	for _, test := range testTemps {
		if actual := f(test.input[0], test.input[1]); actual != test.expect {
			t.Errorf("[%v]actual %v,expect %v", test.name, actual, test.expect)
		}
	}
}

func TestEclid(t *testing.T) {
	testFramework(t, recursion.Gcd)
}
