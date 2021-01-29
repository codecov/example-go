package Calculator

import "testing"

func TestCalculatorAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", 3, result)
	}
}
