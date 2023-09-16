/**
 * Author: Mitch Allen
 * File: secrets_test.go
 */

package secrets

import (
	"testing"
)

func TestSecrets(t *testing.T) {
	expected := "my-secret"

	if got := GetSecret(); got != expected {
		t.Errorf("GetSecret() = %s, didn't return %s", got, expected)
	}
}
