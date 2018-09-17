// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

import (
	"sort"
	"testing"
)

func TestTupleArray(t *testing.T) {
	cases := []struct {
		in, want TupleArray
	}{
		{[][]int{{7, 8}, {1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}, {7, 8}}},
		{[][]int{{7, 1}, {1, 2}, {7, 0}}, [][]int{{1, 2}, {7, 0}, {7, 1}}},
	}

	for _, c := range cases {
		sort.Sort(c.in)
		if !equals(c.in, c.want) {
			format := `
Sorted TupleArray is not equal to expected value.
Expected: %v
Actual: %v
`
			t.Errorf(format, c.want, c.in)
		}
	}
}

func equals(a, b TupleArray) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := range a {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}
