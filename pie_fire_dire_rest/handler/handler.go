package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func BeefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	text, err := fetchText()
	if err != nil {
		http.Error(w, "Failed to fetch text", http.StatusInternalServerError)
	}
	words := cleanText(*text)

	meatCount := countMeatOccurrences(words)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(meatCount); err != nil {
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
	}
}

func cleanText(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z\-]+`)
	words := re.FindAllString(strings.ToLower(text), -1)

	return words
}

func fetchText() (*string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseString := string(body)
	return &responseString, nil
}

func countMeatOccurrences(words []string) map[string]int {
	meatCount := make(map[string]int)
	var mu sync.Mutex

	var wg sync.WaitGroup
	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			mu.Lock()
			meatCount[word]++
			mu.Unlock()
		}(word)
	}

	wg.Wait()
	return meatCount
}
