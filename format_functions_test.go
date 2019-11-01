package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatFuncs(t *testing.T) {
	cases := []struct {
		description string
		funcName    string
		input       string
		expected    string
	}{
		{
			description: "snake",
			funcName:    "snake",
			input:       "SomethingCool",
			expected:    "something_cool",
		},
		{
			description: "camel",
			funcName:    "camel",
			input:       "something_cool",
			expected:    "somethingCool",
		},
		{
			description: "upper",
			funcName:    "upper",
			input:       "IDontKnow",
			expected:    "IDONTKNOW",
		},
		{
			description: "lower",
			funcName:    "lower",
			input:       "ThisTest",
			expected:    "thistest",
		},
		{
			description: "first (keep upper)",
			funcName:    "first",
			input:       "FirstLetter",
			expected:    "F",
		},
		{
			description: "first (keep lower)",
			funcName:    "first",
			input:       "firstLetter",
			expected:    "f",
		},
		{
			description: "first upper",
			funcName:    "first-upper",
			input:       "firstLetter",
			expected:    "F",
		},
		{
			description: "first lower",
			funcName:    "first-lower",
			input:       "FirstLetter",
			expected:    "f",
		},
		{
			description: "capitalize first letter",
			funcName:    "capitalize-first",
			input:       "TheFirstLetterIsUpper",
			expected:    "The first letter is upper",
		},
		{
			description: "capitalize all letter letters",
			funcName:    "capitalize-all",
			input:       "TheFirstLetterIsUpper",
			expected:    "The First Letter Is Upper",
		},
	}

	funcs := FormatFuncs()

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			f, ok := funcs[tc.funcName]

			require.True(t, ok)
			assert.Equal(t, tc.expected, f(tc.input))
		})
	}
}
