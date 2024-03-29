// jsonInjector
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/misc.go
// Original timestamp: 2024/03/20 14:35

package helpers

import (
	"bytes"
	"fmt"
	"strconv"
)

// CustomError implements the error interface
type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func ChangeLog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Printf("CHANGELOG\n=========\n")
	fmt.Print(`
VERSION		DATE			COMMENTAIRE
-------		----			-----------
2.03.00		2024.03.27		Tout les secrets sont encodés (Base64) par défaut
2.00.00		2024.03.xx		Les endpoints sont mieux définis. COBRA est de retour dans la solution (logiciel)
1.11.00		2024.03.20		Les paramètres passés à l'outil sont mieux validés
1.00.04		2024.03.19		L'outil est maintenant "container-aware"
1.00.03		2024.03.19		Force l'architecture AMD64 pour l'outil et le container Docker 
1.00.00		2024.03.15		Version initiale
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
