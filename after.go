package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"time"
)

type PromptState struct {
	Code       int
	Duration   *time.Duration
	CWD        string
	User       string
	Virtualenv string
}

func formatDuration(d time.Duration) string {
	if d.Hours() > 1.0 {
		return fmt.Sprintf("%.1fh", d.Hours())
	}
	if d.Minutes() > 1.0 {
		return fmt.Sprintf("%.1fm", d.Minutes())
	}
	return fmt.Sprintf("%.1fs", d.Seconds())
}

func After(flags *flag.FlagSet) error {
	afterTime := time.Now()
	if flag.NArg() != 3 {
		return fmt.Errorf("incorrect number of args %d != 3 (%v)", flag.NArg(), flag.Args())
	}
	uid := flag.Arg(1)
	code, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		return fmt.Errorf("unable to parse exit code '%s': %s", flag.Arg(2), err)
	}

	beforeState, err := TryPopState(uid)
	if err != nil {
		return fmt.Errorf("unable to pop state: %s", err)
	}

	state := new(PromptState)
	state.Code = code
	state.CWD, _ = os.Getwd()
	if !beforeState.Time.IsZero() {
		d := afterTime.Sub(beforeState.Time)
		state.Duration = &d
	}

	u, _ := user.LookupId(fmt.Sprint(os.Getuid()))
	state.User = u.Username

	state.Virtualenv = os.Getenv("VIRTUAL_ENV")
	if state.Virtualenv != "" {
		state.Virtualenv = path.Base(state.Virtualenv)
	}

	gitState, err := GetGitState()
	if err != nil {
		return fmt.Errorf("unable to get git state: %s", err)
	}

	if strings.HasPrefix(state.CWD, u.HomeDir) {
		state.CWD = strings.Replace(state.CWD, u.HomeDir, "~", 1)
	}

	if state.Duration != nil {
		fmt.Printf("(%d %s) ", state.Code, formatDuration(*state.Duration))
	}
	if state.Virtualenv != "" {
		fmt.Printf("(%s) ", state.Virtualenv)
	}
	if gitState.Branch != "" {
		fmt.Printf("(%s:", gitState.Branch)
		if gitState.HasUntracked {
			fmt.Printf("u")
		}
		if gitState.HasModified {
			fmt.Printf("d")
		}
		if gitState.HasStaged {
			fmt.Printf("s")
		}
		if gitState.Ahead > 0 {
			fmt.Printf("/%d", gitState.Ahead)
		}
		if gitState.Behind > 0 {
			fmt.Printf("/%d", gitState.Behind)
		}
		fmt.Printf(") ")
	}

	fmt.Printf("%s ", state.User)
	fmt.Printf("%s ", state.CWD)
	fmt.Print("$ ")

	return nil
}
