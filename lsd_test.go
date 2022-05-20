package lsd

import "testing"

func TestDistance(t *testing.T) {
	var tests = []struct {
		lhs  string
		rhs  string
		want int
	}{
		{"こんにちわ世界", "こんにちわ世界", 0},
		{"こんにちわ世界", "こにゃちわ世界", 2},
		{"こんにちわ世界", "こにゃにゃちわ世界", 3},
		{"こんにちわ世界", "こんばんわ世界", 2},
		{"こんにちわ世界", "こんにちわ", 2},
		{"こんにちわ世界", "こんばんわ", 4},
		{"こんにちわ世界", "世界", 5},
	}

	for _, test := range tests {
		got := StringDistance(test.lhs, test.rhs)
		if got != test.want {
			t.Fatalf("want %v but %v: %v vs %v", test.want, got, test.lhs, test.rhs)
		}
	}
}
