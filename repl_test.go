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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   test tests",
			expected: []string{"test", "tests"},
		},
		{
			input:    "EVEN mOrE  TestS   ",
			expected: []string{"even", "more", "tests"},
		},
		{
			input:    "last test for this",
			expected: []string{"last", "test", "for", "this"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v,\nexpected: %v", c.input, actual, c.expected)
			}
		}
	}
}
