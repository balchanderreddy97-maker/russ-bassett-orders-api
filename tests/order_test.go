package tests

import "testing"

func TestSample(t *testing.T) {
	expected := 1
	actual := 1

	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}