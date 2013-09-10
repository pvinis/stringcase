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

func TestParse_Dash(t *testing.T) {
	const s, c = "hello-there-dash", Dash
	out := []string{"hello", "there", "dash"}
	if x := Parse(s, c); !equal(x, out) {
		t.Errorf("Parse(%#v, %#v) = %#v, want %#v", s, c, x, out)
	}
}

func TestParse_Camel1(t *testing.T) {
	const s, c = "HelloThereCamel", Camel
	out := []string{"hello", "there", "Camel"}
	if x := Parse(s, c); !equal(x, out) {
		t.Errorf("Parse(%#v, %#v) = %#v, want %#v", s, c, x, out)
	}
}

func TestParse_Camel2(t *testing.T) {
	const s, c = "helloThereCamel", Camel
	out := []string{"hello", "there", "camel"}
	if x := Parse(s, c); !equal(x, out) {
		t.Errorf("Parse(%#v, %#v) = %#v, want %#v", s, c, x, out)
	}
}

func TestParse_Snake(t *testing.T) {
	const s, c = "hello_there_snake", Snake
	out := []string{"hello", "there", "snake"}
	if x := Parse(s, c); !equal(x, out) {
		t.Errorf("Parse(%#v, %#v) = %#v, want %#v", s, c, x, out)
	}
}
