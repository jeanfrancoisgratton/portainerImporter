// jsonInjector
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/terminal.go
// Original timestamp: 2024/03/20 14:35

package helpers

import (
	"fmt"
	"github.com/jwalton/gchalk"
	"syscall"
	"unsafe"
)

const (
	terminalEscape = "\x1b"
)

// TERMINAL FUNCTIONS
func GetTerminalSize() (int, int) {
	var size struct {
		rows    uint16
		cols    uint16
		xpixels uint16
		ypixels uint16
	}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&size)))
	if err != 0 {
		return 0, 0
	}
	return int(size.cols), int(size.rows)
}

// COLOR FUNCTIONS
// ===============
func Red(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightRed().Bold(sentence))
}

func Green(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightGreen().Bold(sentence))
}

func White(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightWhite().Bold(sentence))
}

func Yellow(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightYellow().Bold(sentence))
}
