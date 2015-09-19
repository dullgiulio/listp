package main

import "errors"

// TODO: put type of brace into tree

var (
	openBrackets  = []byte{'(', '[', '{'}
	closeBrackets = []byte{')', ']', '}'}
)

var (
	errNotBracket     = errors.New("not a bracket")
	errExpectComma    = errors.New("unexpected character, want comma")
	errUnmatchedClose = errors.New("unmatched close bracket")
	errUnmatchedOpen  = errors.New("unmatched open bracket")
)

func isByte(c byte, cs []byte) bool {
	for i := range cs {
		if c == cs[i] {
			return true
		}
	}
	return false
}

func isOpen(c byte) bool {
	return isByte(c, openBrackets)
}

func isClose(c byte) bool {
	return isByte(c, closeBrackets)
}

func oppositeBrace(c byte) (byte, error) {
	for i := range openBrackets {
		if c == openBrackets[i] {
			return closeBrackets[i], nil
		}
	}
	return 0, errNotBracket
}

func isSpace(c byte) bool {
	if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
		return true
	}
	return false
}

type group struct {
	current *list
	brace   byte
}

type parser struct {
	pos     int
	str     string
	groups  []group
	current *list
}

func newParser(s string) *parser {
	return &parser{
		str:     s,
		groups:  make([]group, 0),
		current: newList(),
	}
}

func (p *parser) skipSpaces() {
	for ; p.pos < len(p.str); p.pos++ {
		if !isSpace(p.str[p.pos]) {
			break
		}
	}
}

func (p *parser) nextWord() (end int, err error) {
	for ; p.pos < len(p.str); p.pos++ {
		if p.str[p.pos] == ',' {
			end = p.pos
			p.pos++
			return end, nil
		}
		if isSpace(p.str[p.pos]) {
			break
		}
		if isClose(p.str[p.pos]) {
			break
		}
	}
	end = p.pos
	p.skipSpaces()
	if p.end() || isClose(p.str[p.pos]) {
		return end, nil
	}
	if p.str[p.pos] != ',' {
		return 0, errExpectComma
	}
	return end, nil
}

func (p *parser) end() bool {
	return p.pos >= len(p.str)
}

func (p *parser) parse() error {
	for !p.end() {
		p.skipSpaces()
		if isOpen(p.str[p.pos]) {
			cb, err := oppositeBrace(p.str[p.pos])
			if err != nil {
				return err
			}
			p.pos++
			p.groups = append(p.groups, group{p.current, cb})
			p.current = newList()
			continue
		}
		if isClose(p.str[p.pos]) {
			if len(p.groups) == 0 {
				return errUnmatchedClose
			}
			cur := p.current
			if len(p.groups) > 0 {
				p.current = p.groups[len(p.groups)-1].current
				p.current.add(cur)
			}
			var g group
			g, p.groups = p.groups[len(p.groups)-1], p.groups[:len(p.groups)-1]
			if g.brace != p.str[p.pos] {
				// XXX: Error here is more specific: there is a closing bracket, but
				//      it is not matching (like: "(...}")
				return errUnmatchedClose
			}
			p.pos++
			continue
		}
		begin := p.pos
		end, err := p.nextWord()
		if err != nil {
			return err
		}
		if begin == end {
			continue
		}
		p.current.add(newText(p.str[begin:end]))
	}
	if len(p.groups) > 0 {
		return errUnmatchedOpen
	}
	return nil
}
