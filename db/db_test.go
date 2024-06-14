package db

import "testing"

func TestConnection(t *testing.T) {
	_, err := Db()
	if err != nil {
		t.Fatalf("Could not connect to db")
	}
}
