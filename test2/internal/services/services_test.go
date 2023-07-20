package services_test

import (
	"testing"

	"github.com/zeimedee/test2/internal/services"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		word     string
		expected bool
	}{
		{"Hello", true},
		{"123", false},
		{"ABC", true},
		{"a123", true},
		{"GoodMorning", true},
		{"", false},
	}

	for _, testCase := range testCases {
		result := services.IsValid(testCase.word)
		if result != testCase.expected {
			t.Errorf("Expected IsValid(%q) to be %v, but got %v", testCase.word, testCase.expected, result)
		}
	}
}

func TestWordService_StoreWord(t *testing.T) {
	wordService := services.NewWordService()

	wordService.StoreWord("abc")
	wordService.StoreWord("ab")
	wordService.StoreWord("abc")

	expectedCount := 2

	words := wordService.GetWords()

	if count := words["abc"]; count != expectedCount {
		t.Errorf("Expected word count for 'animal' to be %d, but got %d", expectedCount, count)
	}

	if count := words["ab"]; count != 1 {
		t.Errorf("Expected word count for 'house' to be 1, but got %d", count)
	}
}

func TestWordService_GetMostFrequentWord(t *testing.T) {
	wordService := services.NewWordService()

	wordService.StoreWord("abc")
	wordService.StoreWord("ab")
	wordService.StoreWord("ab")

	expectedWord := "ab"

	if word := wordService.GetFrequentWord("a"); word != expectedWord {
		t.Errorf("Expected most frequent word with prefix 'a' to be '%s', but got '%s'", expectedWord, word)
	}

	expectedWord = "abc"

	if word := wordService.GetFrequentWord("abc"); word != expectedWord {
		t.Errorf("Expected most frequent word with prefix 'abc' to be '%s', but got '%s'", expectedWord, word)
	}

	if word := wordService.GetFrequentWord("d"); word != "" {
		t.Errorf("Expected most frequent word with prefix 'd' to be '', but got '%s'", word)
	}
}
