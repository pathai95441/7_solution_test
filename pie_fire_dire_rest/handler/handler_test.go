package handler

import (
	"testing"
)

func TestCountMeatOccurrences(t *testing.T) {
	text := `Fatback t-bone t-bone, pastrami .. t-bone. pork, meatloaf jowl enim. Bresaola t-bone.`
	words := cleanText(text)
	meatCount := countMeatOccurrences(words)

	if meatCount["t-bone"] != 4 {
		t.Errorf("Expected 4 occurrences of t-bone, got %d", meatCount["t-bone"])
	}
	if meatCount["pastrami"] != 1 {
		t.Errorf("Expected 1 occurrence of pastrami, got %d", meatCount["pastrami"])
	}
}
