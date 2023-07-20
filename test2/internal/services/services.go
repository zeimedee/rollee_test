package services

import (
	"regexp"
	"strings"
	"sync"
)

type WordService struct {
	words map[string]int
	mutex sync.RWMutex
}

func IsValid(word string) bool {
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	return wordRegex.MatchString(word)
}

func NewWordService() *WordService {
	return &WordService{
		words: make(map[string]int),
		mutex: sync.RWMutex{},
	}
}

func (ws *WordService) StoreWord(word string) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	lowerCaseWord := strings.ToLower(word)

	ws.words[lowerCaseWord]++
}

func (ws *WordService) GetFrequentWord(prefix string) string {
	ws.mutex.RLock()
	defer ws.mutex.RUnlock()

	lowerCasePrefix := strings.ToLower(prefix)
	var frequentWord string
	maxFrequency := 0

	for word, frequency := range ws.words {
		if strings.HasPrefix(word, lowerCasePrefix) && frequency > maxFrequency {
			frequentWord = word
			maxFrequency = frequency
		}
	}

	return frequentWord
}

func (ws *WordService) GetWords() map[string]int {
	return ws.words
}
