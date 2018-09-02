package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func mainInner() error {
	flag.Parse()
	if flag.NArg() == 0 {
		return fmt.Errorf("missing before/after subcommand")
	}

	switch flag.Arg(0) {
	case "before":
		return Before(flag.CommandLine)
	case "fix":
		FixCursor(fmt.Sprint(aurora.Bold(aurora.Black("%"))))
		return nil
	case "after":
		return After(flag.CommandLine)
	default:
		return fmt.Errorf("unknown subcommand '%s'", flag.Arg(0))
	}
}

func main() {
	if err := mainInner(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	// _, b, _ := cursorPosition()
	// if b > 1 {
	// 	fmt.Println("%")
	// // }
	// fmt.Printf("%v\n", os.Args[1:])
	// fmt.Println(terminfo.GetStdoutDimensions())
	// fmt.Println(aurora.Cyan("══════════════════════════════"))
}
