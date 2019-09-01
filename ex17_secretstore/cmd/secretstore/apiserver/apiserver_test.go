package apiserver

import "testing"

func assertStatus(t *testing.T, received, expected int) {
	t.Helper()
	if received != expected {
		t.Errorf("Expected status to be %d, received %d", expected, received)
	}
}
