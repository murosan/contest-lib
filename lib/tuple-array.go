// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

//start:TupleArray:
type TupleArray [][]int

func (a TupleArray) Len() int {
	return len(a)
}

func (a TupleArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a TupleArray) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return a[i][0] < a[j][0]
}

//end:TupleArray:
