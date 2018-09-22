// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

import "sort"

//start:BigInts:
type BigInts []int64

func (a BigInts) Len() int {
	return len(a)
}

func (a BigInts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a BigInts) Less(i, j int) bool {
	return a[i] < a[j]
}

//end:BigInts:

//start:SearchBigInts:
//imports[sort]
func SearchBigInts(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

//end:SearchBigInts:
