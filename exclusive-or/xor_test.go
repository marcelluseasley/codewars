package kata

import "testing"

func TestXor(t *testing.T) {

	tables := []struct {
		a bool
		b bool
		expected bool
	}{
		{false,false,false},
		{true,false, true},
		{false, true, true},
		{true, true, false},
	}

	for _, table := range tables {
		actual := Xor(table.a, table.b)
		if actual != table.expected {
			t.Errorf("Xor of %t and %t was incorrect, got: %t, want: %t",
				table.a, table.b, actual, table.expected)
		}
	}
	
	
}