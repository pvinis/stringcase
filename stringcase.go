package stringcase

import (
	"regexp"
)

type Case int
const (
	Camel Case = iota
	Snake
	Dash
)

var dashre = regexp.MustCompile("([A-Za-z]+)")

func Parse(s string, c Case) (o []string) {
	switch c {
	case Dash:
		o = dashre.FindAllString(s, -1)
	}
	return
}
