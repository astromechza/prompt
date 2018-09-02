package main

import (
	"bytes"
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

	ps1 := new(bytes.Buffer)

	fmt.Fprint(ps1, "[")

	if state.Duration != nil {
		if state.Code != 0 {
			fmt.Fprintf(ps1, "%v ", state.Code)
		}
		fmt.Fprintf(ps1, "%v ", formatDuration(*state.Duration))
	}
	if state.Virtualenv != "" {
		fmt.Fprintf(ps1, "v:%s ", state.Virtualenv)
	}

	if gitState.Branch != "" {
		fmt.Fprintf(ps1, "%s", gitState.Branch)
		if gitState.HasStaged || gitState.HasUntracked || gitState.HasModified {
			fmt.Fprintf(ps1, ":")
			if gitState.HasUntracked {
				fmt.Fprint(ps1, "u")
			}
			if gitState.HasModified {
				fmt.Fprint(ps1, "d")
			}
			if gitState.HasStaged {
				fmt.Fprint(ps1, "s")
			}
		}
		if gitState.Ahead > 0 || gitState.Behind > 0 {
			fmt.Fprintf(ps1, ":")
			if gitState.Ahead > 0 {
				fmt.Fprintf(ps1, "￪%v", gitState.Ahead)
			}
			if gitState.Behind > 0 {
				fmt.Fprintf(ps1, "￬%v", gitState.Ahead)
			}
		}
		fmt.Fprintf(ps1, " ")
	}

	fmt.Fprintf(ps1, "%v ", state.User)

	fmt.Fprintf(ps1, "%v ", strings.Replace(state.CWD, "/", "/", -1))

	fmt.Fprint(ps1, "⟫ ")

	p := ps1.String()
	fmt.Print(p)

	return nil
}
