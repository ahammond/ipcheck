package main

import (
	"testing"
)

var testCases = []struct {
	description string
	input       string
	expected    IpType
}{
	{
		"malformed",
		"Tom-ay-to, tom-aaaah-to.",
		malformed,
	},
	{
		"public",
		"132.15.129.2",
		public,
	},
	{
		"private",
		"192.168.10.2",
		private,
	},
	{
		"loopback",
		"127.0.0.1",
		loopback,
	},
	{
		"multicast",
		"224.0.0.1",
		multicast,
	},
	{
		"linklocal",
		"169.254.0.1",
		linklocal,
	},
	{
		"6to4",
		"192.88.99.1",
		six2four,
	},
	{
		"documentation",
		"203.0.113.1",
		documentation,
	},
	{
		"reserved",
		"240.0.0.1",
		reserved,
	},
}

func TestCategorize(t *testing.T) {
	for _, tc := range testCases {
		actual := Categorize(tc.input)
		if actual != tc.expected {
			msg := `
	%s: %s
	got: %d
	expected: %d`
			t.Fatalf(msg, tc.description, tc.input, actual, tc.expected)
		}
	}
}
