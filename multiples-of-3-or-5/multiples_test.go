package multiples

import "testing"

func TestMultiple3And5(t *testing.T) {

	result := Multiple3And5(10)
	if result != 23 {
		t.Errorf("Incorret. Expecting %d, got %v", 35, result)
	}
}