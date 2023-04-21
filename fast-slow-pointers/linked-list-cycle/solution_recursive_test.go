package linked_list_cycle

import "testing"

func Test_hasCycleRec(t *testing.T) {
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
			if got := hasCycleRec(tt.head); got != tt.want {
				t.Errorf("hasCycleRec() = %v, want %v", got, tt.want)
			}
		})
	}
}
