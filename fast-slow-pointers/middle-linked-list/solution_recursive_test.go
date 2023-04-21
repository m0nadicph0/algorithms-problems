package middle_linked_list

import (
	"testing"
)

func TestMiddleNodeRecursive(t *testing.T) {
	tests := []struct {
		name string
		head *Node
		want int
	}{
		{
			name: "Case 1",
			head: createList(1, 2, 3, 4, 5),
			want: 3,
		},
		{
			name: "Case 2",
			head: createList(1, 2, 3, 4, 5, 6),
			want: 4,
		},
		{
			name: "Case 3",
			head: createList(1, 2, 3),
			want: 2,
		},
		{
			name: "Case 4",
			head: createList(4, 3, 2, 1, 0),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleNodeRecursive(tt.head); got.Value != tt.want {
				t.Errorf("MiddleNode() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}
