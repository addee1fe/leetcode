package solution

// Ugly hardcoded tests

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	link := "https://leetcode.com/problems/design-tinyurl"
	encoded := c.encode(link)
	decoded := c.decode(encoded)

	if decoded != link {
		t.Errorf("Unable to decode encoded string: %q, encoded: %q, decoded: %q", link, encoded, decoded)
	}
}

func TestCodecEncode(t *testing.T) {
	c := Constructor()
	tests := []string{
		"https://apple.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://google.com",
		"https://leetcode.com/problems/design-tinyurl",
		"https://microsoft.com",
		"https://netflix.com",
	}
	res := make([]string, len(tests))

	for i, v := range tests {
		res[i] = c.encode(v)
	}

	for i, v := range res {
		if tests[i] != c.decode(v) {
			t.Errorf("Unable to decode encoded string: expected: %q, got: %q", tests[i], c.decode(v))
		}
	}
}

func TestCodecDecode(t *testing.T) {
	c := Constructor()
	tests := []string{
		"https://apple.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://google.com",
		"https://leetcode.com/problems/design-tinyurl",
		"https://microsoft.com",
		"https://netflix.com",
	}
	res := make([]string, len(tests))

	for i, v := range tests {
		res[i] = c.encode(v)
	}

	for i, v := range res {
		if tests[i] != c.decode(v) {
			t.Errorf("Unable to decode encoded string: expected: %q, got: %q", tests[i], c.decode(v))
		}
	}
}
