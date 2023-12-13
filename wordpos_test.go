// Assisted with ChatGPT

package wordpos

import (
	"testing"
)

func TestGetPOS(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog"
	ps, err := GetPOS(text)
	if err != nil {
		t.Errorf("GetPOS(%s) returned an error: %v", text, err)
	}

	t.Log(ps)

	// Add assertions based on the expected results.
	// Example:
	if len(ps.Nouns) == 2 {
		t.Errorf("Expected two nouns, but got %d", len(ps.Nouns))
	}
}

// test Is[POS]

func TestIsNoun(t *testing.T) {
	word := "fox"
	isNoun, err := IsNoun(word)
	if err != nil {
		t.Errorf("IsNoun(%s) returned an error: %v", word, err)
	}

	t.Log(isNoun)

	// Add assertions based on the expected results.
	// Example:
	if !isNoun {
		t.Errorf("Expected %s to be a noun, but it wasn't", word)
	}
}

func TestIsVerb(t *testing.T) {
	word := "jump"
	isVerb, err := IsVerb(word)
	if err != nil {
		t.Errorf("IsNoun(%s) returned an error: %v", word, err)
	}

	t.Log(isVerb)

	// Add assertions based on the expected results.
	// Example:
	if !isVerb {
		t.Errorf("Expected %s to be a verb, but it wasn't", word)
	}
}

func TestIsAdjective(t *testing.T) {
	word := "brown"
	isAdj, err := IsAdjective(word)
	if err != nil {
		t.Errorf("IsNoun(%s) returned an error: %v", word, err)
	}

	t.Log(isAdj)

	// Add assertions based on the expected results.
	// Example:
	if !isAdj {
		t.Errorf("Expected %s to be a adjective, but it wasn't", word)
	}
}

func TestIsAdverb(t *testing.T) {
	word := "gentle"
	isAdv, err := IsAdjective(word)
	if err != nil {
		t.Errorf("IsNoun(%s) returned an error: %v", word, err)
	}

	t.Log(isAdv)

	// Add assertions based on the expected results.
	// Example:
	if !isAdv {
		t.Errorf("Expected %s to be a adverb, but it wasn't", word)
	}
}

// Test Lookup

func TestLookup(t *testing.T) {
	word := "quick"
	wordInfo, err := Lookup(word)
	if err != nil {
		t.Errorf("Lookup(%s) returned an error: %v", word, err)
	}

	t.Log(wordInfo)

	// Add assertions based on the expected results.
	// Example:
	if wordInfo == nil {
		t.Errorf("Expected information for %s, but got nil", word)
	}
}

func TestLookupNoun(t *testing.T) {
	word := "fox"
	wordInfo, err := LookupNoun(word)
	if err != nil {
		t.Errorf("Lookup(%s) returned an error: %v", word, err)
	}

	t.Log(wordInfo)

	// Add assertions based on the expected results.
	// Example:
	if wordInfo == nil {
		t.Errorf("Expected information for %s, but got nil", word)
	}
}

func TestLookupVerb(t *testing.T) {
	word := "leap"
	wordInfo, err := LookupVerb(word)
	if err != nil {
		t.Errorf("Lookup(%s) returned an error: %v", word, err)
	}

	t.Log(wordInfo)

	// Add assertions based on the expected results.
	// Example:
	if wordInfo == nil {
		t.Errorf("Expected information for %s, but got nil", word)
	}
}

func TestLookupAdjective(t *testing.T) {
	word := "lazy"
	wordInfo, err := LookupAdjective(word)
	if err != nil {
		t.Errorf("Lookup(%s) returned an error: %v", word, err)
	}

	t.Log(wordInfo)

	// Add assertions based on the expected results.
	// Example:
	if wordInfo == nil {
		t.Errorf("Expected information for %s, but got nil", word)
	}
}

func TestLookupAdverb(t *testing.T) {
	word := "ultimately"
	wordInfo, err := LookupAdverb(word)
	if err != nil {
		t.Errorf("Lookup(%s) returned an error: %v", word, err)
	}

	t.Log(wordInfo)

	// Add assertions based on the expected results.
	// Example:
	if wordInfo == nil {
		t.Errorf("Expected information for %s, but got nil", word)
	}
}

// Test random

func TestRandNoun(t *testing.T) {
	startsWith := "a"
	count := 3
	words, err := RandNoun(startsWith, uint(count))
	if err != nil {
		t.Errorf("RandNoun(%s, %d) returned an error: %v", startsWith, count, err)
	}

	for _, word := range words {
		t.Log(word)
	}

	// Add assertions based on the expected results.
	// Example:
	if len(words) != count {
		t.Errorf("Expected %d words, but got %d", count, len(words))
	}
}

func TestRandAdjective(t *testing.T) {
	startsWith := "am"
	count := 3
	words, err := RandAdjective(startsWith, uint(count))
	if err != nil {
		t.Errorf("RandNoun(%s, %d) returned an error: %v", startsWith, count, err)
	}

	for _, word := range words {
		t.Log(word)
	}

	// Add assertions based on the expected results.
	// Example:
	if len(words) != count {
		t.Errorf("Expected %d words, but got %d", count, len(words))
	}
}

func TestRandVerb(t *testing.T) {
	startsWith := ""
	count := 3
	words, err := RandVerb(startsWith, uint(count))
	if err != nil {
		t.Errorf("RandNoun(%s, %d) returned an error: %v", startsWith, count, err)
	}

	for _, word := range words {
		t.Log(word)
	}

	// Add assertions based on the expected results.
	// Example:
	if len(words) != count {
		t.Errorf("Expected %d words, but got %d", count, len(words))
	}
}

func TestRandAdverb(t *testing.T) {
	startsWith := "si"
	count := 3
	words, err := RandAdverb(startsWith, uint(count))
	if err != nil {
		t.Errorf("RandNoun(%s, %d) returned an error: %v", startsWith, count, err)
	}

	for _, word := range words {
		t.Log(word)
	}

	// Add assertions based on the expected results.
	// Example:
	if len(words) != count {
		t.Errorf("Expected %d words, but got %d", count, len(words))
	}
}
