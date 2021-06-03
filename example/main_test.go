package example_bin

import "testing"

func TestHundred(t *testing.T) {
	res := hundred()

	if res != 100 {
		t.Errorf("Expected 100 but got %d", res)
	}
}
