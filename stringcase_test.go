package stringcase

import "testing"

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParse(t *testing.T) {
	const s, c = "hello-there-dash", Dash
	out := []string{"hello", "there", "dash"}
	if x := Parse(s, c); !equal(x, out) {
		t.Errorf("Parse(%#v, %#v) = %#v, want %#v", s, c, x, out)
	}
}
