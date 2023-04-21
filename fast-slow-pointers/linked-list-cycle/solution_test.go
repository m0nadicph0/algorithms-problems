package linked_list_cycle

import "testing"

func Test_hasCycle(t *testing.T) {

	tests := []struct {
		name string
		head *Node
		want bool
	}{
		{
			name: "Case 1",
			head: listWithCycle(),
			want: true,
		},
		{
			name: "Case 2",
			head: listWithoutCycle(),
			want: false,
		},
		{
			name: "Case 3",
			head: &Node{1, nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func listWithoutCycle() *Node {
	return &Node{1, &Node{2, &Node{3, nil}}}
}

func listWithCycle() *Node {
	head := &Node{3, nil}
	node1 := &Node{2, nil}
	node2 := &Node{0, nil}
	node3 := &Node{-4, nil}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1
	return head
}
