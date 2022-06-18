package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestNewService(t *testing.T) {

	s := NewService()
	if s == nil {
		t.Fatalf("NewService is nil")
	}
}

func TestTranslate(t *testing.T) {

	ctx := context.Background()
	rand.Seed(time.Now().UTC().UnixNano())
	s := NewService()
	for i, _ := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		fmt.Println("Test Iteration : ", i)
		result, err := s.Translate(ctx, language.English, language.Japanese, "test")
		if err != nil {
			t.Fatal("Translate Err: " + err.Error())
		}
		if !strings.Contains(result, "en -> ja : test") {
			t.Fatal("Failed to translate , error result string")
		}
	}

}
