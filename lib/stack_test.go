package cutil

import "testing"

func TestIntStack(t *testing.T) {
	stack := NewIntStack()
	if !stack.isEmpty() {
		t.Error("The IntStack must be empty. But it's not empty.")
	}

	stack.push(100)
	if stack.isEmpty() {
		t.Error("The IntStack must not be empty. But it's empty.")
	}

	stack.push(200)
	stack.push(300)
	if len(stack.st) != 3 {
		t.Errorf("The IntStack size must be 3. But got %v", len(stack.st))
	}

	i := stack.pop()
	if i != 300 {
		t.Errorf("The value must be 300. But got %v", i)
	}

	j := stack.pop()
	if j != 200 {
		t.Errorf("The value must be 200. But got %v", j)
	}
}
