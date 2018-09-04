package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/logrusorgru/aurora"
)

var promptDebug = os.Getenv("PROMPT_DEBUG") != ""

func promptDebugLog(format string, args ...interface{}) {
	if promptDebug {
		log.Printf(format, args...)
	}
}

func runCmdWithDebug(cmd *exec.Cmd) error {
	s := time.Now()
	e := cmd.Run()
	promptDebugLog("%v: %s", cmd.Args, time.Since(s))
	return e
}

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
	case "setup":
		thisBinary, err := filepath.Abs(os.Args[0])
		if err != nil {
			return fmt.Errorf("failed to find abs path to %s", os.Args[0])
		}
		fmt.Println(`PROMPT_PID=$$
_prompt() {
	` + thisBinary + ` fix
	PS1="$(` + thisBinary + ` after ${PROMPT_PID} $1)"
}
PS0='$(` + thisBinary + ` before '${PROMPT_PID}')'
PROMPT_COMMAND='_prompt $?'`)
		return nil
	default:
		return fmt.Errorf("unknown subcommand '%s'", flag.Arg(0))
	}
}

func main() {
	start := time.Now()
	if err := mainInner(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	promptDebugLog("%v: %s", os.Args, time.Since(start))
}
