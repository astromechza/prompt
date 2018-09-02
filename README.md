# prompt

My personal opinionated Bash prompt.

**Features**:

- Shows exit code of previous command if it failed
- Shows execution time of the previous command
- Shows Git branch and status of working tree
- Shows active Python virtualenv name if there is one
- Adds a coloured, trailing '%' to the previous output if it didn't add a newline

## How to install

You'll need a Go development environment for now, I have no plans to release prebuilt binaries at this time.

```
$ go get github.com/AstromechZA/prompt
$ cp $GOPATH/bin/prompt /usr/local/bin/prompt  # Rename it or put it somewhere else if you need to
```

And then in your Bash rc/profiles files:

```
source <(/usr/local/bin/prompt setup)
```
