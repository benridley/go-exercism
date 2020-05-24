package cryptosquare

import "testing"

var tests = []struct {
	pt string // plain text
	ct string // cipher text
}{
	{
		"s#$%^&plunk",
		"su pn lk",
	},
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		if ct := Encode(test.pt); ct != test.ct {
			t.Fatalf(`Encode(%q):
got  %q
want %q`, test.pt, ct, test.ct)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Encode(test.pt)
		}
	}
}
