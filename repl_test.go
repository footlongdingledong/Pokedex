package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input:	" hello world ",
			expected: []string{"hello", "world"},
		},{
			input: " joe  mama",
			expected: []string{"joe", "mama"},
		},{
			input: "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("non-matching input lengths")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedword := c.expected[i]
			if word != expectedword {
				t.Errorf("words do not match")
			}
		}
	}
}
