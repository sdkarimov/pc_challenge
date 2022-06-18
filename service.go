package main

import (
	"context"
	"fmt"

	"time"

	"github.com/sdkarimov/pc_challenge/core"
	"github.com/sdkarimov/pc_challenge/internal/backoff"
	"github.com/sdkarimov/pc_challenge/internal/storage"
	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translatorClient TranslatorAPI
	timeout          backoff.BackOff
	translatorCache  storage.Storage
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
		translatorCache:  storage.NewCache(5, 2),
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {

	fmt.Println("Sleep 0")
	cacheKey := from.String() + to.String() + data
	if v, ok := s.translatorCache.Get(cacheKey); ok {
		fmt.Println("FROM CACHE")
		return v.Value.(string), nil
	}

	result, err := s.translatorClient.Translate(ctx, from, to, data)
	if err == nil {
		s.setToCache(cacheKey, result)
		return result, err
	}

	timeOutOk := true
	sec := 0
	for timeOutOk && err != nil {
		sec, timeOutOk = s.timeout.BackOffIteration()
		fmt.Println("Sec elapsed ", sec)
		if result, err = s.translatorClient.Translate(ctx, from, to, data); err == nil {
			s.setToCache(cacheKey, result)
			return result, err
		}

	}
	return result, err
}

func (s *Service) setToCache(key, result string) {
	val := core.CacheVal{Value: result, CreateDate: time.Now().Unix()}
	s.translatorCache.Set(key, val)
}
