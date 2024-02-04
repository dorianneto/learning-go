package adder_test

import (
	"c_13/adder"
	"testing"
)

func Test_AddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Errorf("incorrect result: expected %d, got %d", 5, result)
	}
}
