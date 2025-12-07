package binary_search

import "testing"

func TestLowerBound(t *testing.T) {
	tests := []struct {
		arr       []int
		target    int
		want      bool
		wantIndex int
	}{
		{[]int{1, 2, 3, 5}, 3, true, 2},
		{[]int{1, 2, 3, 5}, 6, false, -1},
		{[]int{1, 2, 3, 5}, 0, true, 0},
		{[]int{1, 2, 3, 5}, 5, true, 3},
		{[]int{1, 2, 3, 5}, 4, true, 3},
	}
	for _, test := range tests {
		index, ok := LowerBound(test.arr, test.target)
		if ok != test.want {
			t.Errorf("LowerBound(%v, %d) = (%t), want (%t)", test.arr, test.target, ok, test.want)
		}
		if index != test.wantIndex {
			t.Errorf("LowerBound(%v, %d) = (%d), want (%d)", test.arr, test.target, index, test.wantIndex)
		}
	}
}
