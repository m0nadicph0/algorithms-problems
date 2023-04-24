package adt

import "testing"

type UpdateParams struct {
	index int
	value int
}

type QueryParams struct {
	start int
	end   int
}

func TestBinaryTreeSegmentTree_Query(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		start int
		end   int
		want  int
	}{
		{
			name:  "Case 1",
			input: []int{1, 2, 3},
			start: 0,
			end:   2,
			want:  6,
		},
		{
			name:  "Case 2",
			input: []int{1, 2, 3, 4, 3, 2, 1},
			start: 0,
			end:   6,
			want:  16,
		},
		{
			name:  "Case 3",
			input: []int{1, 2, 3, 4, 3, 2, 1},
			start: 1,
			end:   3,
			want:  9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinaryTreeSegmentTree(tt.input)
			if got := b.Query(tt.start, tt.end); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTreeSegmentTree_Update(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		update UpdateParams
		query  QueryParams
		want   int
	}{
		{
			name:   "Case 1",
			input:  []int{1, 2, 3},
			update: UpdateParams{index: 0, value: 3},
			query:  QueryParams{start: 0, end: 2},
			want:   8,
		},
		{
			name:   "Case 2",
			input:  []int{1, 2, 3, 4, 3, 2, 1},
			update: UpdateParams{index: 3, value: 8},
			query:  QueryParams{start: 0, end: 6},
			want:   20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinaryTreeSegmentTree(tt.input)
			b.Update(tt.update.index, tt.update.value)
			if got := b.Query(tt.query.start, tt.query.end); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
