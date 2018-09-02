package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
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
			fmt.Fprintf(ps1, "%v ", aurora.Red(state.Code))
		}
		fmt.Fprintf(ps1, "%v ", aurora.Cyan(formatDuration(*state.Duration)))
	}
	if state.Virtualenv != "" {
		fmt.Fprintf(ps1, "v:%s ", aurora.Magenta(state.Virtualenv))
	}

	if gitState.Branch != "" {
		fmt.Fprintf(ps1, "%s", aurora.Bold(aurora.Green(gitState.Branch)))
		if gitState.HasStaged || gitState.HasUntracked || gitState.HasModified {
			fmt.Fprintf(ps1, ":")
			if gitState.HasUntracked {
				fmt.Fprint(ps1, aurora.Bold("u"))
			}
			if gitState.HasModified {
				fmt.Fprint(ps1, aurora.Bold(aurora.Brown("d")))
			}
			if gitState.HasStaged {
				fmt.Fprint(ps1, aurora.Bold("s"))
			}
		}
		if gitState.Ahead > 0 || gitState.Behind > 0 {
			fmt.Fprintf(ps1, ":")
			if gitState.Ahead > 0 {
				fmt.Fprintf(ps1, aurora.Sprintf(aurora.Green("￪%v"), gitState.Ahead))
			}
			if gitState.Behind > 0 {
				fmt.Fprintf(ps1, aurora.Sprintf(aurora.Red("￬%v"), gitState.Ahead))
			}
		}
		fmt.Fprintf(ps1, " ")
	}

	fmt.Fprintf(ps1, "%v ", aurora.Bold(aurora.Cyan(state.User)))

	fmt.Fprintf(ps1, "%v ", strings.Replace(state.CWD, "/", aurora.Bold(aurora.Black("/")).String(), -1))

	fmt.Fprint(ps1, aurora.Bold("⟫ "))

	p := ps1.String()
	p = string(regexp.MustCompile("\x1B\\[.*?m").ReplaceAll([]byte(p), []byte("\\[$0\\]")))
	fmt.Print(p)

	return nil
}
