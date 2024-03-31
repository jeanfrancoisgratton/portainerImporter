// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/misc.go
// Original timestamp: 2024/03/28 21:34

package helpers

import (
	"bytes"
	"fmt"
	"strconv"
)

func ChangeLog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Printf("CHANGELOG\n=========\n")
	fmt.Print(`
VERSION		DATE			COMMENTAIRE
-------		----			-----------
1.00.00		2024.03.29		Initial version
`)
}

// NUMBER FORMATTING FUNCTIONS
// ===========================

// This function was originally written in 1993, in C, by my friend Jean-François Gauthier (jief@brebis.dyndns.org)
// I've ported it in C# in 2011. It is then a third iteration of this function
// This function transforms a multi-digit number in International Notation; 1234567 thus becomes 1,234,567
func SI(nombre uint64) string {
	var strN string
	var strbR bytes.Buffer
	var nLen, nPos int

	strN = strconv.FormatUint(nombre, 10)
	strN = ReverseString(strN)
	nLen = len(strN)

	for nPos < nLen {
		if nPos != 0 && nPos%3 == 0 {
			strbR.WriteString(",")
			strbR.WriteString(string(strN[nPos]))
		} else {
			strbR.WriteString(string(strN[nPos]))
		}
		nPos++
	}

	strN = strbR.String()
	strN = ReverseString(strN)

	return strN
}

// OTHER FUNCTIONS
// ===============

// This function takes a string and returns its reverse
// Thus, "12345" becomes "54321"
func ReverseString(inputStr string) string {
	runes := []rune(inputStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
