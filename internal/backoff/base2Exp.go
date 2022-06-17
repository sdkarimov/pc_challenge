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

func (b *BackOffBase2Exp) BackOffIteration() (bool, int) {
	if b.CurrentCount <= b.StopCount {
		fmt.Println("Sleep sec = ", b.TimeOut)
		currentTimeout := b.TimeOut
		time.Sleep(time.Duration(b.TimeOut) * time.Second)
		b.TimeOut <<= 1
		b.CurrentCount += 1
		return true, currentTimeout
	}
	return false, b.TimeOut
}
