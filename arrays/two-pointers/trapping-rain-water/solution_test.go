package trapping_rain_water

import "testing"

func TestTrapWater(t *testing.T) {

	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Case 1",
			heights: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:    6,
		},
		{
			name:    "Case 2",
			heights: []int{4, 2, 0, 3, 2, 5},
			want:    9,
		},
		{
			name:    "Test case 3",
			heights: []int{1, 2, 3, 4, 5},
			want:    0,
		},
		{
			name:    "Test case 4",
			heights: []int{5, 4, 3, 2, 1},
			want:    0,
		},
		{
			name:    "Test case 5",
			heights: []int{1, 0, 1, 0, 1},
			want:    2,
		},
		{
			name:    "Test case 6",
			heights: []int{2, 0, 2},
			want:    2,
		},
		{
			name:    "Test case 7",
			heights: []int{1, 1, 1, 1, 1},
			want:    0,
		},
		{
			name:    "Test case 8",
			heights: []int{1, 1, 1, 2, 2, 2, 1, 1, 1},
			want:    0,
		},
		{
			name:    "Test case 9",
			heights: []int{3, 2, 1, 2, 3},
			want:    4,
		},
		{
			name:    "Test case 10",
			heights: []int{0, 1, 2, 3, 2, 1, 0},
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrapWater(tt.heights); got != tt.want {
				t.Errorf("TrapWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
