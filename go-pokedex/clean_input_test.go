package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   I'm looking up and down side to side like a roller coaster",
			expected: []string{"i'm", "looking", "up", "and", "down", "side", "to", "side", "like", "a", "roller", "coaster"},
		},
		{
			input:    "this                 is        another test",
			expected: []string{"this", "is", "another", "test"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "SIMPLY LOVELY!",
			expected: []string{"simply", "lovely!"},
		},
		{
			input:    "JuSt An         InChIDeNt!",
			expected: []string{"just", "an", "inchident!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("word lengths do not match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected output '%v' got '%v'", expectedWord, word)
			}
		}
	}
}
