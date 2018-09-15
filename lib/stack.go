// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

//start:NewStack:
type Stack struct {
	st []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(i int) {
	s.st = append([]int{i}, s.st...)
}

func (s *Stack) pop() int {
	if len(s.st) == 0 {
		panic("empty array. cannot pop!")
	}

	head := s.st[0]
	if len(s.st) == 1 {
		s.st = s.st[:0]
	} else {
		s.st = s.st[1:]
	}
	return head
}

func (s *Stack) isEmpty() bool {
	return len(s.st) == 0
}

//end:NewStack:
