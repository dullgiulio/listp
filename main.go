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
	inputs := []string{
		"[arg0, arg0.1], (arg1, (arg2.1, arg2.2),  arg3, arg4, (arg5.1,arg5.2,  arg5.3 ) )",
		"(世界, 世界, 世界), arg1, 世界",
		"arg1, arg2, arg3",
		"(arg1, arg2, arg3)",
		"()",
		"arg1",
		"arg1), arg2, arg3",
		"(arg1, arg2, arg3",
	}
	for i := range inputs {
		fmt.Printf("%s ->\n", inputs[i])
		p := newParser(inputs[i])
		err := p.parse()
		if err != nil {
            log.Print("parse error: ", err)
			continue
		}
		printEntries(p.current, 0)
	}
}
