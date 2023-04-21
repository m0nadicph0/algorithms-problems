package two_sum_2

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{
			name:   "case 1",
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 2},
		},
		{
			name:   "case 2",
			input:  []int{2, 3, 4},
			target: 6,
			want:   []int{1, 3},
		},
		{
			name:   "case 3",
			input:  []int{-1, 0},
			target: -1,
			want:   []int{1, 2},
		},
		{
			name:   "case 4",
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 17,
			want:   []int{7, 10},
		},
		{
			name:   "case 5",
			input:  []int{1, 2, 3, 4, 5, 7, 8, 22},
			target: 3,
			want:   []int{1, 2},
		},
		{
			name:   "case 6",
			input:  []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			target: 30,
			want:   []int{5, 10},
		},
		{
			name:   "case 7",
			input:  []int{-8, -4, -3, -2, -1, 0, 2, 4, 6},
			target: -7,
			want:   []int{2, 3},
		},
		{
			name:   "case 8",
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target: 20,
			want:   []int{5, 15},
		},
		{
			name:   "case 9",
			input:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
			target: -19,
			want:   []int{1, 2},
		},
		{
			name:   "case 10",
			input:  []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100},
			target: 160,
			want:   []int{12, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumII(tt.input, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumII() = %v, want %v", got, tt.want)
			}
		})
	}
}
