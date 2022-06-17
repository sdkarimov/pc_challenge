package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/text/language"
)

// TranslatorAPI in an interface of the 3rd-party translation provider API
//which translates strings from one language to another.
type TranslatorAPI interface {
	Translate(ctx context.Context, from, to language.Tag, data string) (string, error)
}

// translatorStub is a TranslatorAPI implementation which is used
// only for testing purposes
type translatorStub struct {
	minDelay  time.Duration
	maxDelay  time.Duration
	errorProb float64
}

func newTranslatorStub(minDelay, maxDelay time.Duration, errorProbability float64) *translatorStub {
	return &translatorStub{
		minDelay:  minDelay,
		maxDelay:  maxDelay,
		errorProb: errorProbability,
	}
}

// Translate returns fake translation string or error. In any case it delays execution for some time
// to emulate remote service. Error is returned with probability set by errorProb.
func (ts translatorStub) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {
	time.Sleep(ts.randomDuration())

	if rand.Float64() < ts.errorProb {
		return "", errors.New("translation failed")
	}

	res := fmt.Sprintf("%v -> %v : %v -> %v", from, to, data, strconv.FormatInt(rand.Int63(), 10))
	return res, nil
}

func (ts translatorStub) randomDuration() time.Duration {
	delta := ts.maxDelay - ts.minDelay
	var delay time.Duration = ts.minDelay + time.Duration(rand.Int63n(int64(delta)))
	return delay
}
