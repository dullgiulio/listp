package main

import (
	"fmt"
	"log"
)

func printEntries(e entry, tabs int) {
	for i := 0; i < tabs; i++ {
		fmt.Print("\t")
	}
	fmt.Printf("'%s'\n", e.String())
	children := e.Children()
	if children != nil {
		for i := range children {
			printEntries(children[i], tabs+1)
		}
	}
}

func main() {
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
		fmt.Printf("%s\n", inputs[i].s)
		p := newParser(inputs[i].s)
		err := p.parse()
		if inputs[i].err != nil {
			if err == nil {
				log.Print("err expected but not found")
				continue
			}
			if err != inputs[i].err {
				log.Print("unexpected-err: ", inputs[i].err, " ; expected: ", err)
				continue
			}
			log.Print("expected-err: ", err)
			continue
		}
		if err != nil {
			log.Print("unexpected-err: ", err)
			continue
		}
		printEntries(p.current, 0)
	}
}
