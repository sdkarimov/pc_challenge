package main

import (
	"context"
	"fmt"
	"offlineChallenge/internal/backoff"
	"time"

	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translatorClient TranslatorAPI
	timeout          backoff.BackOff
}

func NewService() *Service {
	t := newTranslatorStub(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)

	b := backoff.NewBackOffBase2Exp(3)

	return &Service{
		translatorClient: t,
		timeout:          b,
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {

	// 1)  Exp backoff package
	fmt.Println("Sleep 0\n")
	result, err := s.translatorClient.Translate(ctx, from, to, data)
	if err == nil {
		return result, err
	}

	timeOutOk := true
	sec := 0
	for timeOutOk && err != nil {
		timeOutOk, sec = s.timeout.BackOffIteration()
		fmt.Println("Sec elapsed ", sec)
		if result, err = s.translatorClient.Translate(ctx, from, to, data); err == nil {
			return result, err
		}

	}
	return result, err
}
