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

	git "gopkg.in/src-d/go-git.v4"

	"github.com/logrusorgru/aurora"
)

type PromptState struct {
	Code         int
	Duration     *time.Duration
	CWD          string
	User         string
	Virtualenv   string
	GitBranch    string
	GitUntracked bool
	GitModified  bool
	GitStaged    bool
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

	repo, err := git.PlainOpen(path.Join(state.CWD))
	if err == nil {
		ref, _ := repo.Head()
		state.GitBranch = ref.Name().Short()
		w, _ := repo.Worktree()
		s, _ := w.Status()
		for _, v := range s {
			if v.Staging != git.Untracked && v.Staging != git.Unmodified {
				state.GitStaged = true
			} else if v.Worktree == git.Untracked {
				state.GitUntracked = true
			} else if v.Worktree != git.Unmodified {
				state.GitModified = true
			}
		}
	}

	if strings.HasPrefix(state.CWD, u.HomeDir) {
		state.CWD = strings.Replace(state.CWD, u.HomeDir, "~", 1)
	}

	parts := make([]interface{}, 0)

	if state.Duration != nil {
		parts = append(parts, fmt.Sprintf("x:%d", state.Code))
		parts = append(parts, fmt.Sprintf("t:%s", formatDuration(*state.Duration)))
	}
	if state.Virtualenv != "" {
		parts = append(parts, fmt.Sprintf("v:%s", state.Virtualenv))
	}
	if state.GitBranch != "" {
		g := state.GitBranch
		if state.GitUntracked {
			g += "u"
		}
		if state.GitModified {
			g += "m"
		}
		if state.GitStaged {
			g += "g"
		}
		parts = append(parts, g)
	}

	parts = append(parts, state.User)
	parts = append(parts, state.CWD)

	for i, p := range parts {
		if i%2 == 0 {
			fmt.Print(aurora.BgGray(p).String())
			fmt.Print(aurora.BgCyan(aurora.Gray("▌")).String())
		} else {
			fmt.Print(aurora.BgCyan(p).String())
			fmt.Print(aurora.BgGray(aurora.Cyan("▌")).String())
		}
	}
	fmt.Print(" > ")

	return nil
}
