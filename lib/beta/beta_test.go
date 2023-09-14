/**
 * Author: Mitch Allen
 * File: beta_test.go
 */

package beta

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	a := 1
	b := 2
	expected := a * b

	if got := Multiply(a, b); got != expected {
		t.Errorf("Multiply(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestDivide(t *testing.T) {
	a := 1
	b := 2
	expected := a / b

	if got := Divide(a, b); got != expected {
		t.Errorf("Divide(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}
