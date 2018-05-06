package compiler

import (
	"testing"
)

func TestToGoName(t *testing.T) {
	tests := map[string]string{
		"a":      "A",
		"aBC":    "ABC",
		"aaBbCc": "AaBbCc",
		"_A":     "Prefix__A",
		"_a":     "Prefix__a",
		"1":      "Prefix_1",
		" a":     "Prefix_A",
		"-a":     "Prefix_A",
		"aa.bb":  "AaBb",
		"-":      "Prefix_",
		"--":     "Prefix_",
	}
	for name, want := range tests {
		goName := toGoName(name, "Prefix_")
		if goName != want {
			t.Errorf("%q: got %q, want %q", name, goName, want)
		}
	}
}
