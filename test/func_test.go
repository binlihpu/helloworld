package test

import "testing"

// TestAbs TestAbs
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func BenchmarkAbs(b *testing.B) {
	result := Abs(-9)
	if result != 9 {
		b.Error("Unexpected result: ", result)
	}
}
