package adder

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumbers(1, 2)
	if result != 3 {
		t.Errorf("addNumbers() = %d, want 3", result)
	}
}
