// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cutil

import (
	"bytes"
	"testing"
)

func strTest(t *testing.T, inputs, expectedValues []string, setfn func(*Scanner)) {
	t.Helper()
	for i, v := range inputs {
		input := bytes.NewBufferString(v)
		s := NewScanner(input)
		setfn(s)
		r := s.ScanStr()
		if r != expectedValues[i] {
			t.Errorf("Input: %v, Expected: %v, Actual: %v", v, expectedValues[i], r)
		}
	}
}

func TestScanner_ScanLine(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("test\nstr\ning")
	s := NewScanner(input)
	s.SetSplitter("\n")
	r1 := s.ScanLine()
	if r1 != "test" {
		t.Errorf("Input: %v, Expected: %v, Actual: %v", input, "test", r1)
	}
	r2 := s.ScanLine()
	if r2 != "str" {
		t.Errorf("Input: %v, Expected: %v, Actual: %v", input, "str", r2)
	}
	r3 := s.ScanLine()
	if r3 != "ing" {
		t.Errorf("Input: %v, Expected: %v, Actual: %v", input, "ing", r3)
	}
}

func TestScanner_ScanStr(t *testing.T) {
	var inputs = []string{
		"test_string",
		"test string testing",
		"test\nstring\ntesting",
		"te,st",
		"12345",
	}

	t.Run("split by new line", func(t *testing.T) {
		t.Parallel()
		expectedValues := []string{
			"test_string",
			"test string testing",
			"test",
			"te,st",
			"12345",
		}
		strTest(t, inputs, expectedValues, func(s *Scanner) { s.SetSplitter("\n") })
	})

	t.Run("split by space", func(t *testing.T) {
		t.Parallel()
		expectedValues := []string{
			"test_string",
			"test",
			"test",
			"te,st",
			"12345",
		}
		strTest(t, inputs, expectedValues, func(s *Scanner) { s.SetSplitter(" ") })
	})

	t.Run("split by rune", func(t *testing.T) {
		t.Parallel()
		expectedValues := []string{
			"t",
			"t",
			"t",
			"t",
			"1",
		}
		strTest(t, inputs, expectedValues, func(s *Scanner) { s.SetSplitter("") })
	})

	t.Run("split by new line", func(t *testing.T) {
		t.Parallel()
		expectedValues := []string{
			"test_string",
			"test string testing",
			"test",
			"te",
			"12345",
		}
		strTest(t, inputs, expectedValues, func(s *Scanner) { s.SetSplitterByRune(',') })
	})
}

func intTest(t *testing.T, inputs []string, expectedValues []int, setfn func(*Scanner)) {
	t.Helper()
	for i, v := range inputs {
		input := bytes.NewBufferString(v)
		s := NewScanner(input)
		setfn(s)
		r := s.ScanInt()
		if r != expectedValues[i] {
			t.Errorf("Input: %v, Expected: %v, Actual: %v", v, expectedValues[i], r)
		}
	}
}

func TestScanner_ScanInt(t *testing.T) {
	inputs := []string{
		"12345",
		"123\n45",
		"123 45",
		"123  45",
	}
	t.Run("split by space", func(t *testing.T) {
		t.Parallel()
		expectedValues := []int{
			12345,
			123,
			123,
			123,
		}
		intTest(t, inputs, expectedValues, func(s *Scanner) { s.SetSplitter(" ") })
	})

	t.Run("panic when input was not int", func(t *testing.T) {
		panicTest(t, func() {
			input := bytes.NewBufferString("isNotInt")
			s := NewScanner(input)
			s.ScanInt()
		})
	})
}

func strsTest(t *testing.T, inputs []string, expectedValues [][]string) {
	t.Helper()
	for i, v := range inputs {
		input := bytes.NewBufferString(v)
		s := NewScanner(input)
		s.SetSplitter(" ")
		r := s.ScanStrs(3)
		if !strsEquals(r, expectedValues[i]) {
			t.Errorf("Input: %v, Expected: %v, Actual: %v", v, expectedValues[i], r)
		}
	}
}

func TestScanner_ScanStrs(t *testing.T) {
	inputs := []string{
		"test\nstring\nslice",
		"test string slice",
	}
	outputs := [][]string{
		{"test", "string", "slice"},
		{"test", "string", "slice"},
	}
	strsTest(t, inputs, outputs)
}

func intsTest(t *testing.T, inputs []string, expectedValues [][]int) {
	t.Helper()
	for i, v := range inputs {
		input := bytes.NewBufferString(v)
		s := NewScanner(input)
		s.SetSplitter(" ")
		r := s.ScanInts(3)
		if !intsEquals(r, expectedValues[i]) {
			t.Errorf("Input: %v, Expected: %v, Actual: %v", v, expectedValues[i], r)
		}
	}
}

func TestScanner_ScanInts(t *testing.T) {
	inputs := []string{
		"123\n456\n789",
		"987 654 321",
	}
	outputs := [][]int{
		{123, 456, 789},
		{987, 654, 321},
	}
	intsTest(t, inputs, outputs)
}

func panicTest(t *testing.T, f func()) {
	t.Helper()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("should be panic, but was not.")
			}
		}()
		// This function should cause a panic
		f()
	}()
}

func strsEquals(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func intsEquals(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
