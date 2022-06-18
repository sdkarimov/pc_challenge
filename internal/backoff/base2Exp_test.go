package backoff

import "testing"

func TestNewBackOffBase2Exp(t *testing.T) {
	if b := NewBackOffBase2Exp(3); b == nil {
		t.Fatal("NewBackOffBase2Exp is nil")
	}
}

func TestBackOffIterationPowerOf3(t *testing.T) {
	b := NewBackOffBase2Exp(3)

	// TODO: testCases array then iterate it
	for id, tout := range []int{1, 2, 4, 8} {
		timeout, ok := b.BackOffIteration()
		if !ok {
			t.Fatal(id, " iteration, timeout not OK ")
		}
		if timeout != tout {
			t.Fatal(id, " iteration, timeout expected ", tout, ", got :", timeout)
		}
	}

	timeout, ok := b.BackOffIteration()
	if ok {
		t.Fatal("Fourth iteration is not expected ")
	}
	if timeout != 16 {
		t.Fatal("Fourth iteration  is not expected, but timeout may 16, got :", timeout)
	}
}
