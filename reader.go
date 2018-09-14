// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//start:ScanInt:
//imports[fmt]
func ScanInt() (i int) {
	fmt.Scan(&i)
	return
}

//end:ScanInt:

//start:NewReader:
//imports[bufio,os]
type Reader struct {
	rdr  *bufio.Reader
	size int
}

func NewReader(size int) *Reader {
	return &Reader{
		bufio.NewReaderSize(os.Stdin, size),
		size,
	}
}

//end:NewReader:

//start:ReadLine:
func ReadLine(r *Reader) string {
	buf := make([]byte, 0, r.size)
	for {
		line, isPrefix, err := r.rdr.ReadLine()
		if err != nil {
			panic(err)
		}
		buf = append(buf, line...)
		if !isPrefix {
			break
		}
	}
	return string(buf)
}

//end:ReadLine:

//start:ReadStrs:
//imports[strings]
//dependsOn[ReadLine]
func ReadStrs(r *Reader, sep string) []string {
	return strings.Split(ReadLine(r), sep)
}

//end:ReadStrs:

//start:ReadInt:
//dependsOn[ReadLine]
func ReadInt(r *Reader) (n int) {
	if i, e := strconv.Atoi(ReadLine(r)); e == nil {
		n = i
	}
	return
}

//end:ReadInt:

//start:ReadInts:
//imports[strconv]
//dependsOn[ReadStrs]
func ReadInts(r *Reader, sep string) []int {
	a := ReadStrs(r, sep)
	n := make([]int, 0)
	for _, v := range a {
		if i, e := strconv.Atoi(v); e == nil {
			n = append(n, i)
		}
	}
	return n
}

//end:ReadInts:

//start:ReadInt64:
//imports[strconv]
//dependsOn[ReadLine]
func ReadInt64(r *Reader) int64 {
	v, _ := strconv.ParseInt(ReadLine(r), 10, 64)
	return v
}

//end:ReadInt64:

//start:ReadInts64:
//imports[strconv]
//dependsOn[ReadStrs]
func ReadInts64(r *Reader, sep string) []int64 {
	a := ReadStrs(r, sep)
	n := make([]int64, 0)
	for _, v := range a {
		if i, e := strconv.ParseInt(v, 10, 64); e == nil {
			n = append(n, i)
		}
	}
	return n
}

//end:ReadInts64:
