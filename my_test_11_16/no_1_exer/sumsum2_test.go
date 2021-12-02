package no_1_exer

import (
	"reflect"
	"testing"
)

func TestSumsum2(t *testing.T) {
	got := Sumsum(3, 4)
	want := 8

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v ", want, got)
	}
}
