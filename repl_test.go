package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "  hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HELLO woRlD   ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual slice (%d) doens't match length of expected slice (%d)", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word (%s) doesn't match expectedWord (%s)", word, expectedWord)
			}
		}
	}
}
