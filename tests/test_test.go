package test

import (
	"testing"
)

// Dummy test to check if Github Actions are working as expected.
func TestDummy(t *testing.T) {
	one := 1
	if one != 1 {
		t.Errorf("1 isn't equal to one.")
	}
}
