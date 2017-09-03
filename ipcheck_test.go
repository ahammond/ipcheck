package main

import (
	"testing"
)

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		"malformed",
		"Tom-ay-to, tom-aaaah-to.",
		"malformed",
	},
	{
		"malformed-short",
		"123.123.1",
		"malformed",
	},
	{
		"malformed-long",
		"123.123.12.41.14",
		"malformed",
	},
	{
		"malformed-large-number",
		"123.123.12.256",
		"malformed",
	},
	{
		"malformed-two-dots",
		"123..123.1.33",
		"malformed",
	},
	{
		"malformed-comma",
		"123.123,1.35",
		"malformed",
	},
	{
		"malformed-leading-zero",
		"0.12.42.131",
		"malformed",
	},
	{
		"public",
		"132.15.129.2",
		"public",
	},
	{
		"private10",
		"10.168.10.2",
		"private",
	},
	{
		"private172",
		"172.16.10.2",
		"private",
	},
	{
		"private192",
		"192.168.10.2",
		"private",
	},
	{
		"private100",
		"192.168.10.2",
		"private",
	},
	{
		"private198",
		"198.18.10.2",
		"private",
	},
	{
		"loopback",
		"127.0.0.1",
		"loopback",
	},
	{
		"loopback0",
		"0.0.0.0",
		"loopback",
	},
	{
		"multicast",
		"224.0.0.1",
		"multicast",
	},
	{
		"linklocal",
		"169.254.0.1",
		"linklocal",
	},
	{
		"6to4",
		"192.88.99.1",
		"six2four",
	},
	{
		"documentation192",
		"192.0.2.1",
		"documentation",
	},
	{
		"documentation198",
		"198.51.100.1",
		"documentation",
	},
	{
		"documentation203",
		"203.0.113.1",
		"documentation",
	},
	{
		"reserved",
		"240.0.0.1",
		"reserved",
	},
}

func TestCategorize(t *testing.T) {
	for _, tc := range testCases {
		actual := IpCheck(tc.input)
		if actual != tc.expected {
			msg := `
	%s: %s
	got: %s
	expected: %s`
			t.Fatalf(msg, tc.description, tc.input, actual, tc.expected)
		}
	}
}
