package main

import "testing"

func TestTrue(t *testing.T) {
	if true != true {
		t.Errorf("expected %v, got %v", true, true)
		t.Fail()
	}
}
