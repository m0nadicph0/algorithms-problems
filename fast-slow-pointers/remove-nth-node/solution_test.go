package remove_nth_node

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {

	tests := []struct {
		name string
		head *Node
		n    int
		want []int
	}{
		{
			name: "Case 1",
			head: createList(1, 2, 3, 4, 5),
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Case 2",
			head: createList(1),
			n:    1,
			want: []int{},
		},
		{
			name: "Case 3",
			head: createList(1, 2),
			n:    1,
			want: []int{1},
		},
		{
			name: "Case 5",
			head: createList(1, 2, 3, 4),
			n:    1,
			want: []int{1, 2, 3},
		},
		{
			name: "Case 6",
			head: createList(1, 2),
			n:    2,
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.head, tt.n).ToList()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createList(nums ...int) *Node {
	head := &Node{nums[0], nil}
	current := head
	for _, num := range nums[1:] {
		current.Next = &Node{num, nil}
		current = current.Next
	}
	return head
}
