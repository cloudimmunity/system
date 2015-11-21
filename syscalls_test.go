package system

import "testing"

func TestCallNumTableIsOkX86Family32(t *testing.T) {
	if !callNumTableIsOkX86Family32() {
		t.Error("Malformed table!")
	}
}

func TestCallNumTableIsOkX86Family64(t *testing.T) {
	if !callNumTableIsOkX86Family64() {
		t.Error("Malformed table!")
	}
}