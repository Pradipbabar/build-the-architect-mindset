package main

import "testing"

func TestEcho(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		sep      string
		newline  bool
		expected string
	}{
		{"single arg", []string{"hello"}, " ", false, "hello"},
		{"two args default sep", []string{"hello", "world"}, " ", false, "hello world"},
		{"custom separator", []string{"a", "b", "c"}, ",", false, "a,b,c"},
		{"newline added", []string{"hello"}, " ", true, "hello\n"},
		{"newline with custom sep", []string{"a", "b"}, "-", true, "a-b\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := echo(tt.args, tt.sep, tt.newline)
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
