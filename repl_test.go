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
			input: "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input: "Good evening TO YouU I gUess",
			expected: []string{"good", "evening", "to", "youu", "i", "guess"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length are not equal: %v vs %v", len(actual), len(c.expected))
			continue
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not equal %v", word, expectedWord)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}