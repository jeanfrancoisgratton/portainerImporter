// portainerImporter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/errorMessages.go
// Original timestamp: 2024/03/30 07:50

package helpers

//import (
//	"fmt"
//	"strconv"
//)
//
//// CustomError implements the error interface
//type CustomError struct {
//	Title   string
//	Message string
//	Code    int
//	IsFatal bool
//}
//
//func (e CustomError) Error() string {
//	if e.IsFatal {
//		return e.Fatal()
//	} else {
//		return e.Warning()
//	}
//}
//
//func (e CustomError) Fatal() string {
//	builtString := ""
//	if e.Title != "" {
//		if e.Code != 0 {
//			builtString = fmt.Sprintf("%s (code: %d): %s\n", Red(e.Title), White(strconv.Itoa(e.Code)), Red(e.Message))
//		} else {
//			builtString = fmt.Sprintf("%s: %s\n", Red(e.Title), White(e.Message))
//		}
//	} else {
//		if e.Code != 0 {
//
//			builtString = fmt.Sprintf("%s (code: %d): %s\n", Red(e.Message), White(strconv.Itoa(e.Code)))
//		} else {
//			builtString = fmt.Sprintf("%s: %s\n", Red(e.Title), White(e.Message))
//		}
//	}
//	return builtString
//}
//
//func (e CustomError) Warning() string {
//	builtString := ""
//	if e.Title != "" {
//		if e.Code != 0 {
//			builtString = fmt.Sprintf("%s (code: %d): %s\n", Red(e.Title), White(strconv.Itoa(e.Code)), Red(e.Message))
//		} else {
//			builtString = fmt.Sprintf("%s: %s\n", Red(e.Title), White(e.Message))
//		}
//	} else {
//		if e.Code != 0 {
//
//			builtString = fmt.Sprintf("%s (code: %d): %s\n", Yellow(e.Message), White(strconv.Itoa(e.Code)))
//		} else {
//			builtString = fmt.Sprintf("%s: %s\n", Yellow(e.Title), White(e.Message))
//		}
//	}
//	return builtString
//}
