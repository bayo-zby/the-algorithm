package array_test

import (
	"algorithm/array"
	"testing"
)

type searchArrayTemp struct {
	input_arr   []int  // 搜索目标
	input_value int    //待查找数值
	expect      int    //正确结果
	name        string // 测试组别名
}

var TestTemps = []searchArrayTemp{
	{
		[]int{1, 2, 3, 4, 5}, 1, 0, "在首位",
	},
}

func benckmarkFramework(testTemps []searchArrayTemp, b *testing.B, f func([]int, int) int) {
	for i := 0; i < b.N; i++ {
		for _, test := range testTemps {
			if actual := f(test.input_arr, test.input_value); actual != test.expect {
				b.Errorf("[{%v}]actual %v,expect %v", test.name, actual, test.expect)
			}
		}

	}
}

func BenchmarkBinarySearch(b *testing.B) {
	benckmarkFramework(TestTemps, b, array.BinarySearch)
}
