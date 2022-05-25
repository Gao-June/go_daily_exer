/*
	练习 单元测试
*/

package main

import (
	"testing"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.FailNow()
	}

	if add(4, 5) != 9 {
		t.FailNow()
	}
}

func TestSub(t *testing.T) {
	if sub(1, 2) != -1 {
		t.FailNow()
	}
	if sub(6, 2) != 4 {
		t.FailNow()
	}
}
