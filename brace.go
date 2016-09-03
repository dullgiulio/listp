package main

import "errors"

var (
	errNotBracket = errors.New("not a bracket")
)

type brace byte

var (
	openBrackets  = []byte{'(', '[', '{'}
	closeBrackets = []byte{')', ']', '}'}
	pairCache     = []string{"()", "[]", "{}"}
)

func isByte(c byte, cs []byte) bool {
	for i := range cs {
		if c == cs[i] {
			return true
		}
	}
	return false
}

func (b brace) pairOf(list []byte) string {
	for i := 0; i < len(list); i++ {
		if b == brace(list[i]) {
			if i < len(pairCache) {
				return pairCache[i]
			}
			return ""
		}
	}
	return ""
}

func (b brace) pair() string {
	if p := b.pairOf(openBrackets); p != "" {
		return p
	}
	return b.pairOf(closeBrackets)
}

func (b brace) getOpen() brace {
	if !b.isOpen() {
		if ob, err := b.oppositeBrace(); err == nil {
			return ob
		}
	}
	return b
}

func (b brace) getClose() brace {
	if !b.isClose() {
		if ob, err := b.oppositeBrace(); err == nil {
			return ob
		}
	}
	return b
}

func (b brace) isOpen() bool {
	return isByte(byte(b), openBrackets)
}

func (b brace) isClose() bool {
	return isByte(byte(b), closeBrackets)
}

func (b brace) oppositeBrace() (brace, error) {
	for i := range openBrackets {
		if byte(b) == openBrackets[i] {
			return brace(closeBrackets[i]), nil
		}
	}
	for i := range closeBrackets {
		if byte(b) == closeBrackets[i] {
			return brace(openBrackets[i]), nil
		}
	}
	return 0, errNotBracket
}
