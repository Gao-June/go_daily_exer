package no_1_exer

import (
	"reflect"
	"testing"
)

func TestSumsum(t *testing.T) {
	got := Sumsum(1, 2)
	want := 3

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v ", want, got)
	}
}
