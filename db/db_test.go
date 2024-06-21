package db

import "testing"

func TestConnection(t *testing.T) {
	_, err := Connection()
	if err != nil {
		t.Fatalf("Could not connect to db: %v", err)
	}
}
