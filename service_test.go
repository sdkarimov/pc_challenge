package main

import "testing"

func TestNewService(t *testing.T) {

	s := NewService()
	if s == nil {
		t.Fatalf("NewService is nil")
	}
}
