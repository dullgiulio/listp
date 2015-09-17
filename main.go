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
	s := "[arg0, arg0.1], (arg1, (arg2.1, arg2.2),  arg3, arg4, (arg5.1,arg5.2,  arg5.3 ) )"
	p := newParser(s)
	if err := p.parse(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", s)
	printEntries(p.current, 0)
}
