package main

import (
    "testing"
)

func TestLists(t *testing.T) {
	inputs := []struct {
		s   string
		err error
	}{
		{"[arg0, arg0.1], (arg1, (arg2.1, arg2.2),  arg3, arg4, (arg5.1,arg5.2,  arg5.3 ) )", nil},
		{"(世界, 世界, 世界), arg1, 世界", nil},
		{"arg1, arg2, arg3", nil},
		{"(arg1, arg2, arg3)", nil},
		{"()", nil},
		{"arg1", nil},
		{"arg1), arg2, arg3", errUnmatchedClose},
		{"(arg1, arg2, arg3", errUnmatchedOpen},
	}
	for i := range inputs {
		p := newParser(inputs[i].s)
		err := p.parse()
		if inputs[i].err != nil {
			if err == nil {
				t.Error("err expected but not found")
				continue
			}
			if err != inputs[i].err {
				t.Error("unexpected-err: ", inputs[i].err, " ; expected: ", err)
				continue
			}
			// We got an error but it's the expected one
            continue
		}
		if err != nil {
			t.Error("unexpected-err: ", err)
			continue
		}
	    // TODO: Count elements and compare len
    }
}
