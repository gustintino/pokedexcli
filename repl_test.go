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
			input:    "                  HELLO                                         WORld            ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for n, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) == 0 {
			t.Errorf("the resulting slice is empty")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words '%s' and '%s' do not match in test case %d", word, expectedWord, n)
			}
		}
	}

}
