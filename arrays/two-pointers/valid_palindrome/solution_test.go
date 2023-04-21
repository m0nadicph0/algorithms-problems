package valid_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Case 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Case 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Case 3",
			s:    " ",
			want: true,
		},
		{
			name: "Case 4",
			s:    " madam ",
			want: true,
		},
		{
			name: "Case 5",
			s:    "Able was I ere I saw Elba",
			want: true,
		},
		{
			name: "Case 6",
			s:    "Eva, can I see bees in a cave?",
			want: true,
		},
		{
			name: "Case 7",
			s:    "Step on no pets.",
			want: true,
		},
		{
			name: "Case 8",
			s:    "No lemon, no melon",
			want: true,
		},
		{
			name: "Case 9",
			s:    "Was it a car or a cat I saw?",
			want: true,
		},
		{
			name: "Case 10",
			s:    "Sir, I demand, I am a maid named Iris",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
