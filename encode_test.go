package oauth1

import (
	"testing"
)

func TestPercentEncode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},
		{"%", "%25"},
		{"-._", "-._"},
	}
	for _, c := range cases {
		if output := PercentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}
