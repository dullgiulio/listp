package main

import "fmt"

type entry interface {
	String() string
	Children() []entry
}

type text string

func newText(s string) *text {
	t := text(s)
	return &t
}

func (t *text) String() string {
	return string(*t)
}

func (t *text) Children() []entry {
	return nil
}

type list struct {
	entries []entry
	brace   brace
}

func newList() *list {
	return &list{entries: make([]entry, 0)}
}

func (l *list) add(e entry) {
	l.entries = append(l.entries, e)
}

func (l *list) String() string {
	return fmt.Sprintf("[list %s %d]", l.brace.pair(), len(l.entries))
}

func (l *list) Children() []entry {
	return l.entries
}
