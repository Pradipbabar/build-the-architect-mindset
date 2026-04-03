package main

import (
	"testing"
)

func TestParseProblems(t *testing.T) {
	tests := []struct {
		name    string
		records [][]string
		want    []problem
	}{
		{
			name:    "basic parsing",
			records: [][]string{{"5+5", "10"}, {"1+1", "2"}},
			want:    []problem{{q: "5+5", a: "10"}, {q: "1+1", a: "2"}},
		},
		{
			name:    "trims answer whitespace",
			records: [][]string{{"5+5", " 10 "}},
			want:    []problem{{q: "5+5", a: "10"}},
		},
		{
			name:    "empty input",
			records: [][]string{},
			want:    []problem{},
		},
		{
			name:    "ignores malformed records",
			records: [][]string{{"5+5", "10"}, {"broken"}},
			want:    []problem{{q: "5+5", a: "10"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseProblems(tt.records)

			if len(got) != len(tt.want) {
				t.Errorf("len(got) = %d, len(tt.want) = %d", len(got), len(tt.want))
			}

			for i, problem := range got {
				if problem.q != tt.want[i].q || problem.a != tt.want[i].a {
					t.Errorf("got %v, want %v", problem, tt.want[i])
				}
			}
		})
	}
}
