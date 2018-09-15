package cutil

import "testing"

func TestReverseStr(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, World", "dlroW ,olleH"},
		{"cutil123", "321lituc"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseStr(c.in)
		if got != c.want {
			format :=
				`
ReverseStr(%q) was not equal to %q.
Want: %q
Actual: %q
`
			t.Errorf(format, c.in, c.want, c.want, got)
		}
	}
}
