package main

import (
	"bytes"
	"os/exec"
	"strings"
	"sync"
)

type GitState struct {
	Branch       string
	Ahead        int
	Behind       int
	HasModified  bool
	HasStaged    bool
	HasUntracked bool
}

func gitBranchSymbol() string {
	c := exec.Command("git", "symbolic-ref", "-q", "HEAD")
	var b bytes.Buffer
	c.Stdout = &b
	c.Run()
	if b.Len() > 0 {
		return strings.TrimSpace(b.String())
	}
	c = exec.Command("git", "name-rev", "--name-only", "--no-undefined", "--always", "HEAD")
	var d bytes.Buffer
	c.Stdout = &d
	if err := c.Run(); err != nil {
		return ""
	}
	return strings.TrimSpace(d.String())
}

func gitNumAhead() int {
	c := exec.Command("git", "log", "--oneline", "@{u}..")
	var b bytes.Buffer
	c.Stdout = &b
	c.Run()
	output := strings.TrimSpace(b.String())
	if len(output) > 0 {
		lines := strings.Split(output, "\n")
		return int(len(lines))
	}
	return 0
}

func gitNumBehind() int {
	c := exec.Command("git", "log", "--oneline", "..@{u}")
	var b bytes.Buffer
	c.Stdout = &b
	c.Run()
	output := strings.TrimSpace(b.String())
	if len(output) > 0 {
		lines := strings.Split(output, "\n")
		return int(len(lines))
	}
	return 0
}

func gitHasUntracked() bool {
	c := exec.Command("git", "ls-files", "--other", "--exclude-standard")
	return c.Run() != nil
}

func gitHasModified() bool {
	c := exec.Command("git", "diff", "--quiet")
	return c.Run() != nil
}

func gitHasStaged() bool {
	c := exec.Command("git", "diff", "--cached", "--quiet")
	return c.Run() != nil
}

func GetGitState() (*GitState, error) {
	o := new(GitState)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		o.Branch = gitBranchSymbol()
		if strings.HasPrefix(o.Branch, "refs/heads/") {
			o.Branch = o.Branch[11:]
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		o.Ahead = gitNumAhead()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		o.Behind = gitNumBehind()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		o.HasModified = gitHasModified()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		o.HasStaged = gitHasStaged()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		o.HasUntracked = gitHasUntracked()
		wg.Done()
	}()
	wg.Wait()
	return o, nil
}
