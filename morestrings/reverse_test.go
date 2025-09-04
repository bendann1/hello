package morestrings

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, World!", "!dlroW ,olleH"},
		{" ", " "},
		{"!@#$%^&*()", ")(*&^%$#@!"},
		{"1234567890", "0987654321"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"日本語", "語本日"},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
