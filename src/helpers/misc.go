// portainerImporter : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/misc/misc.go
// 4/16/23 21:35:03

package helpers

import (
	"fmt"
	"github.com/jwalton/gchalk"
)

func Changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION		DATE			COMMENT
-------		----			-------
0.200		2023.07.12		token mgmt
0.100		2023.07.05		stub, configs mgmt
`)
}

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

// FIXME : Normal() is the same as White()
func Normal(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithWhite().Bold(sentence))
}
