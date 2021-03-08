package sloths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertIsSlothful(t *testing.T, s string, expected bool) {
	if IsSlothful(s) != expected {
		if expected {
			t.Errorf("%s is supposed to be slothful", s)
		} else {
			t.Errorf("%s is not supposed to be slothful", s)
		}
	}
}

type isSlothfulTestCase struct {
	str      string
	expected bool
}

var isSlothfulTestCases = []isSlothfulTestCase{
	{str: "hello, world!", expected: false},
	{str: "hello, slothful world!", expected: true},
	{str: "Sloths rule!", expected: true},
	{str: "Nothing like an iced hibiscus tea! 🧊🌺", expected: true},
	{str: "Get your 🌺 flowers! They're going fast! 🏎️", expected: false},
}

func TestIsSlothful(t *testing.T) {
	for _, c := range isSlothfulTestCases {
		assertIsSlothful(t, c.str, c.expected)
	}
}

func TestIsSlothfulWithTestify(t *testing.T) {
	for _, c := range isSlothfulTestCases {
		assert.Equal(t, c.expected, IsSlothful(c.str))
	}
}
