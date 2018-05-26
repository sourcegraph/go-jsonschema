package compiler

import "testing"

func TestLineComments(t *testing.T) {
	tests := map[string]string{
		"foo":      "// foo",
		"foo\nbar": "// foo\n// bar",
	}
	for input, want := range tests {
		got := lineComments(input)
		if got != want {
			t.Errorf("%q:  got %q, want %q", input, got, want)
		}
	}
}
