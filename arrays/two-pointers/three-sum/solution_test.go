package three_sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Case 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "Case 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Case 3",
			nums: []int{0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "basic case",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "empty input",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "no triplet sum to zero",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "two distinct triplets",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "two identical triplets",
			nums: []int{-2, -2, -2, 4},
			want: [][]int{{-2, -2, 4}},
		},
		{
			name: "one positive element",
			nums: []int{1},
			want: [][]int{},
		},
		{
			name: "one negative element",
			nums: []int{-1},
			want: [][]int{},
		},
		{
			name: "all negative elements",
			nums: []int{-1, -2, -3, -4},
			want: [][]int{},
		},
		{
			name: "all positive elements",
			nums: []int{1, 2, 3, 4},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
