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
	timeout, ok := b.BackOffIteration()
	if !ok {
		t.Fatal("Zero iteration, timeout not OK ")
	}
	if timeout != 1 {
		t.Fatal("Zero iteration, timeout expected 1, got :", timeout)
	}

	timeout, ok = b.BackOffIteration()
	if !ok {
		t.Fatal("First iteration, timeout not OK ")
	}
	if timeout != 2 {
		t.Fatal("First iteration, timeout expected 2, got :", timeout)
	}

	timeout, ok = b.BackOffIteration()
	if !ok {
		t.Fatal("Second iteration, timeout not OK ")
	}
	if timeout != 4 {
		t.Fatal("Second iteration, timeout expected 4, got :", timeout)
	}

	timeout, ok = b.BackOffIteration()
	if !ok {
		t.Fatal("Third iteration, timeout not OK ")
	}
	if timeout != 8 {
		t.Fatal("Third iteration, timeout expected 8, got :", timeout)
	}

	timeout, ok = b.BackOffIteration()
	if ok {
		t.Fatal("Fourth iteration is not expected ")
	}
	if timeout != 16 {
		t.Fatal("Fourth iteration  is not expected, but timeout may 16, got :", timeout)
	}
}
