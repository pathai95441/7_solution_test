package handler

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func BeefSummaryHandler() (*map[string]int32, error) {
	text, err := fetchText()
	if err != nil {
		return nil, err
	}
	words := cleanText(*text)

	meatCount := countMeatOccurrences(words)

	return &meatCount, nil
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

func countMeatOccurrences(words []string) map[string]int32 {
	meatCount := make(map[string]int32)
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
