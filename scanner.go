// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

import (
	"bufio"
	"os"
	"strconv"
	"unicode/utf8"
)

// 1. set the separator
//    ex1. SetSplitter(" ")
//    ex2. SetSplitterByRune(',')
// 2. run scanning

//start:NewScanner:
//imports[bufio,os]
var sc = bufio.NewScanner(os.Stdin)
//end:NewScanner:

//start:ScanStr:
//dependsOn[NewScanner]
func ScanStr() string {
	sc.Scan()
	return sc.Text()
}

//end:ScanStr:

//start:ScanInt:
//imports[strconv]
//dependsOn[NewScanner]
func ScanInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

//end:ScanInt:

//start:ScanLine:
//dependsOn[NewScanner]
func ScanLine() string {
	sc.Scan()
	return sc.Text()
}

//end:ScanLine:

//start:ScanStrs:
//dependsOn[ScanStr]
func ScanStrs(len int) []string {
	s := make([]string, len)
	for i := 0; i < len; i++ {
		s[i] = ScanStr()
	}
	return s
}

//end:ScanStrs:

//start:ScanInts:
//dependsOn[ScanInt]
func ScanInts(len int) []int {
	s := make([]int, len)
	for i := 0; i < len; i++ {
		s[i] = ScanInt()
	}
	return s
}

//end:ScanInts:

//start:SetSplitter:
//dependsOn[NewScanner]
func SetSplitter(sep string) {
	switch sep {
	case "":
		sc.Split(bufio.ScanRunes)
	case " ":
		sc.Split(bufio.ScanWords)
	default:
		sc.Split(bufio.ScanLines)
	}
}

//end:SetSplitter:

//start:SetSplitterByRune:
//imports[unicode/utf8]
//dependsOn[NewScanner]
//splitter from: https://golang.org/src/bufio/scan.go?s=13093:13171#L380
func SetSplitterByRune(sep rune) {
	sc.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Skip leading separator.
		start := 0
		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if r != sep {
				break
			}
		}
		// Scan until separator, marking end of word.
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if r == sep || r == '\n' {
				return i + width, data[start:i], nil
			}
		}
		// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}
		// Request more data.
		return start, nil, nil
	})
}

//end:SetSplitterByRune:
