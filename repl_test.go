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
			input:    "This is a Much longer  Sentence with multiple  spaces and  capitalized words",
			expected: []string{"this", "is", "a", "much", "longer", "sentence", "with", "multiple", "spaces", "and", "capitalized", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word: \"%v\", but got \"%v\"", expectedWord, word)
			}
		}
	}
}
