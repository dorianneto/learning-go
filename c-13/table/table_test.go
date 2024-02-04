package table

import (
	"testing"
)

func Test_DoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 3, "+", 5, ""},
		{"subtraction", 8, 2, "-", 6, ""},
		{"multiplication", 6, 3, "*", 18, ""},
		{"division", 14, 2, "/", 7, ""},
		{"bad_division", 4, 0, "/", 0, "division by zero"},
		{"unknown_operator", 4, 0, "?", 0, "unknown operator ?"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}

			if err == nil {
				return
			}

			errMsg := err.Error()
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
