package add

import "testing"

func TestAdd(t *testing.T) {
	tests := [][]int{ // test cases table
		{1, 1, 2},
		{100, 200, 300},
		{-2, 2, 0},
		{-3, -5, -8},
		{999, -1, 998},
	}

	for _, tc := range tests { // 遍历所有测试用例
		if res := Add(tc[0], tc[1]); res != tc[2] {
			t.Errorf("want: %d, got: %d", tc[2], res)
		}
	}
}
