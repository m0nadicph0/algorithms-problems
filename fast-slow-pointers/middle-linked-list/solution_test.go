package middle_linked_list

import (
	"testing"
)

func TestMiddleNode(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleNode(tt.head); got.Value != tt.want {
				t.Errorf("MiddleNode() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func createList(nums ...int) *Node {
	head := &Node{
		Value: nums[0],
		Next:  nil,
	}
	current := head
	for _, n := range nums[1:] {
		current.Next = &Node{Value: n, Next: nil}
		current = current.Next
	}
	return head
}
