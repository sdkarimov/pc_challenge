package main

import (
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/sdkarimov/pc_challenge/core"
	"github.com/sdkarimov/pc_challenge/internal/backoff"
	"github.com/sdkarimov/pc_challenge/internal/storage"
	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translatorClient  TranslatorAPI
	timeout           backoff.BackOff
	translatorCache   storage.Storage
	processingStorage storage.Storage
}

func NewService() *Service {
	t := newTranslatorStub(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)

	return &Service{
		translatorClient:  t,
		timeout:           backoff.NewBackOffBase2Exp(3),
		translatorCache:   storage.NewCache(10, 2),
		processingStorage: storage.NewCache(20, 10),
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {

	cacheKey := from.String() + to.String() + data

	// dedublicate
	if _, ok := s.processingStorage.Get(cacheKey); ok {
		return data, errors.New("Service is processing request with params:  " +
			from.String() + "; " + to.String() + "; " + data)
	} else {
		s.processingStorage.Set(cacheKey, core.CacheVal{Value: true, CreateDate: time.Now().Unix()})
	}

	defer s.processingStorage.Delete(cacheKey)

	if v, ok := s.translatorCache.Get(cacheKey); ok {
		fmt.Println("FROM CACHE")
		return v.(core.CacheVal).Value.(string), nil
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
