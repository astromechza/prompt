package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func cursorPosition() (int, int, error) {
	fd := int(os.Stdin.Fd())
	termios, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	if err != nil {
		return 0, 0, err
	}
	newState := *termios
	newState.Lflag &^= (unix.ECHO |
		unix.ICANON |
		unix.IGNBRK |
		unix.BRKINT |
		unix.IGNPAR |
		unix.PARMRK |
		unix.INPCK |
		unix.ISTRIP |
		unix.INLCR |
		unix.IGNCR |
		unix.ICRNL |
		unix.IXON |
		unix.ICANON |
		unix.OPOST |
		unix.ISIG |
		unix.IXANY |
		unix.IMAXBEL |
		0)
	if err := unix.IoctlSetTermios(fd, ioctlWriteTermios, &newState); err != nil {
		return 0, 0, err
	}

	defer unix.IoctlSetTermios(fd, ioctlWriteTermios, termios)

	// same as $ echo -e "\033[6n"
	cmd := exec.Command("tput", "u7")
	randomBytes := &bytes.Buffer{}
	cmd.Stdout = randomBytes
	_ = cmd.Start()

	// capture keyboard output from echo command
	reader := bufio.NewReader(os.Stdin)
	cmd.Wait()

	// by printing the command output, we are triggering input
	fmt.Print(randomBytes)
	// capture the triggered stdin from the print
	text, _ := reader.ReadSlice('R')

	// check for the desired output
	if strings.Contains(string(text), ";") {
		re := regexp.MustCompile(`\d+;\d+`)
		line := re.FindString(string(text))
		parts := strings.Split(line, ";")
		y, _ := strconv.Atoi(parts[0])
		x, _ := strconv.Atoi(parts[1])
		return y - 1, x - 1, nil
	}
	return 0, 0, fmt.Errorf("bad line: '%s'", text)
}

func GetSize(fd int) (width, height int, err error) {
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	return int(ws.Col), int(ws.Row), nil
}
