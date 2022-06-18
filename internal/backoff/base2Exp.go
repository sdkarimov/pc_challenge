package backoff

import (
	"fmt"
	"time"
)

type BackOffBase2Exp struct {
	StopCount    int
	CurrentCount int
	TimeOut      int
}

func NewBackOffBase2Exp(iterationCount int) *BackOffBase2Exp {
	return &BackOffBase2Exp{StopCount: iterationCount, CurrentCount: 0, TimeOut: 1}
}

func (b *BackOffBase2Exp) BackOffIteration() (int, bool) {
	if b.CurrentCount <= b.StopCount {
		fmt.Println("Back off iteration = ", b.CurrentCount)
		fmt.Println("Back off sleep sec = ", b.TimeOut)

		currentTimeout := b.TimeOut
		time.Sleep(time.Duration(b.TimeOut) * time.Second)
		b.TimeOut <<= 1
		b.CurrentCount += 1
		return currentTimeout, true
	}
	return b.TimeOut, false
}
