// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

import (
	"math"
)

//start:BellmanFord:
//imports[math]
type EdgeB struct {
	From     int
	To       int
	Distance int
}

func BellmanFord(edges []EdgeB, start, goal int) (connected bool, cost int) {
	if len(edges) == 0 {
		return
	}

	// initialize
	max := math.MaxInt32
	distances := make(map[int]int)
	for _, e := range edges {
		distances[e.From] = max
		distances[e.To] = max
	}

	// solve
	distances[start] = 0
	updated := true

	for updated {
		updated = false
		for _, e := range edges {
			c := distances[e.From] + e.Distance
			if distances[e.From] != max && distances[e.To] > c {
				distances[e.To] = c
				updated = true
			}
		}
	}

	if distances[goal] != max {
		connected = true
	}
	cost = distances[goal]

	return
}

//end:BellmanFord:
