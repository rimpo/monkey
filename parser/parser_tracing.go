package parser

import "fmt"

const disableLog = true

var p = 0

func trace(name string) string {
	if disableLog {
		return ""
	}
	p += 1
	for i := 0; i < p; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("BEGIN %s\n", name)
	return name
}

func untrace(name string) {
	if disableLog {
		return
	}

	for i := 0; i < p; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("END %s\n", name)
	p -= 1
}

func log(msg string) {
	if disableLog {
		return
	}

	for i := 0; i < p; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("log: %s\n", msg)
}
