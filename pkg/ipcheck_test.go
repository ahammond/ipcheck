package ipcheck

import (
	"testing"
)

var cases = []struct {
	description string
	input       string
	expected    IpType
	expectedErr IpError
}{
	{
		"public",
		"132.15.129.2",
		Public,
		nil,
	},
	{
		"private",
		"192.168.10.2",
		Private,
		nil,
	},
	{
		"loopback",
		"127.0.0.1",
		Loopback,
		nil,
	},
	{
		"multicast",
		"224.0.0.1",
		Multicast,
		nil,
	},
	{
		"linklocal",
		"169.254.0.1",
		LinkLocal,
		nil,
	},
	{
		"6to4",
		"192.88.99.1",
		Six2Four,
		nil,
	},
	{
		"documentation",
		"203.0.113.1",
		Documentation,
		nil,
	},
	{
		"reserved",
		"240.0.0.1",
		Reserved,
		nil,
	},
}

func TestType(t *testing.T) {
	for _, c := range cases {
		r, err := Type(c.input)
		if r != c.expected || err != c.expectedErr {
			msg := `
	%s input: %s
	got: %+v, %+v
	expected: %+v, %+v`
			t.Fatalf(msg,
				c.description, c.input,
				r, err,
				c.expected, c.expectedErr)
		}
	}
}

func TestMalformed(t *testing.T) {
	i := "Tom-ay-to, tom-aaaah-to."
	_, err := Type(i)
	errStr := err.Error()
	if errStr != "malformed" {
		t.Fatalf(`
	malformed input: %s
	got: %+v
	expected: malformed`,
			i, errStr)
	}
}
