// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

//start:Min:
func Min(n ...int) int {
	min := n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

//end:Min:

//start:Max:
func Max(n ...int) int {
	max := n[0]
	for _, v := range n[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

//end:Max:

//start:MinMax:
func MinMax(n ...int) (int, int) {
	min := n[0]
	max := n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

//end:MinMax:

//start:Abs:
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

//end:Abs:

//start:Product:
func Product(n ...int) int {
	p := n[0]
	for _, v := range n[1:] {
		p *= v
	}
	return p
}

//end:Product:
