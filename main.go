package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type POS string

const (
	POS_Noun      POS = "n"
	POS_Verb      POS = "v"
	POS_Adjective POS = "adj"
	POS_Adverb    POS = "adv"

	noun_file      = "wordnet/dict/data.noun"
	verb_file      = "wordnet/dict/data.verb"
	adjective_file = "wordnet/dict/data.adj"
	adverb_file    = "wordnet/dict/data.adv"
)

type Word struct {
	ID           uint
	Word         string
	PartOfSpeech POS
	Definition   string
}

type POSSet struct {
	Nouns      []*Word
	Verbs      []*Word
	Adjectives []*Word
	Adverbs    []*Word
	Rest       []*Word
}

func GetPOS(text string) (*POSSet, error) {
	nouns, err := GetNouns(text)
	if err != nil {
		return nil, err
	}
	verbs, err := GetVerbs(text)
	if err != nil {
		return nil, err
	}
	adjectives, err := GetAdjectives(text)
	if err != nil {
		return nil, err
	}
	adverbs, err := GetAdverbs(text)
	if err != nil {
		return nil, err
	}

	ps := &POSSet{
		Nouns:      nouns,
		Verbs:      verbs,
		Adjectives: adjectives,
		Adverbs:    adverbs,
	}

	InsertRest(text, ps)

	return ps, nil
}

// Whole sentences

func GetNouns(text string) ([]*Word, error) {
	words := strings.Split(text, " ")
	wordsRes := make([]*Word, 0)

	for i := 0; i < len(words); i++ {
		word, err := LookupNoun(words[i])
		if err == nil {
			wordsRes = append(wordsRes, word)
		} else {
			if err.Error() != "word not found in file(s)" {
				return nil, err
			}
		}
	}

	return wordsRes, nil
}

func GetVerbs(text string) ([]*Word, error) {
	words := strings.Split(text, " ")
	wordsRes := make([]*Word, 0)

	for i := 0; i < len(words); i++ {
		word, err := LookupVerb(words[i])
		if err == nil {
			wordsRes = append(wordsRes, word)
		} else {
			if err.Error() != "word not found in file(s)" {
				return nil, err
			}
		}
	}

	return wordsRes, nil
}

func GetAdjectives(text string) ([]*Word, error) {
	words := strings.Split(text, " ")
	wordsRes := make([]*Word, 0)

	for i := 0; i < len(words); i++ {
		word, err := LookupAdjective(words[i])
		if err == nil {
			wordsRes = append(wordsRes, word)
		} else {
			if err.Error() != "word not found in file(s)" {
				return nil, err
			}
		}
	}

	return wordsRes, nil
}

func GetAdverbs(text string) ([]*Word, error) {
	words := strings.Split(text, " ")
	wordsRes := make([]*Word, 0)

	for i := 0; i < len(words); i++ {
		word, err := LookupAdverb(words[i])
		if err == nil {
			wordsRes = append(wordsRes, word)
		} else {
			if err.Error() != "word not found in file(s)" {
				return nil, err
			}
		}
	}

	return wordsRes, nil
}

func InsertRest(text string, ps *POSSet) {
	words := strings.Split(text, " ")
	wordsRes := make([]*Word, 0)

	// This is woefully inefficient and I hope to change this in the future
	for i := 0; i < len(words); i++ {
		word, err := Lookup(words[i])
		if err != nil {
			wordsRes = append(wordsRes, word)
		}
	}

	ps.Rest = wordsRes
}

// Word only

func IsNoun(word string) (bool, error) {
	_, err := LookupNoun(word)
	if err != nil {
		if err.Error() == "word not found in file(s)" { // irrelevant error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsVerb(word string) (bool, error) {
	_, err := LookupVerb(word)
	if err != nil {
		if err.Error() == "word not found in file(s)" { // irrelevant error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsAdjective(word string) (bool, error) {
	_, err := LookupAdjective(word)
	if err != nil {
		if err.Error() == "word not found in file(s)" { // irrelevant error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsAdverb(word string) (bool, error) {
	_, err := LookupAdverb(word)
	if err != nil {
		if err.Error() == "word not found in file(s)" { // irrelevant error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Lookup(word string) (*Word, error) {
	data, err := LookupNoun(word)
	if err == nil {
		return data, nil
	}
	data, err = LookupVerb(word)
	if err == nil {
		return data, nil
	}
	data, err = LookupAdjective(word)
	if err == nil {
		return data, nil
	}
	data, err = LookupAdverb(word)
	if err == nil {
		return data, nil
	}

	return nil, err
}

// internal generic
func lookupType(word string, file string, partOfSpeech POS) (*Word, error) {
	word_fmt := strings.ToLower(strings.ReplaceAll(word, " ", "_"))
	// regex: matches the correct word, then captures the ID and definition.
	regex_string := strings.Replace(`(?im)^(\d+?) \d\d [^\\]+? WORD .+?\| (.+)$`, "WORD", word_fmt, 1)
	regex, err := regexp.Compile(regex_string)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	matching_bytes := regex.FindSubmatch(b)

	if len(matching_bytes) == 0 {
		return nil, fmt.Errorf("word not found in file(s)")
	}

	word_id, err := strconv.ParseUint(string(matching_bytes[1]), 10, 64)
	if err != nil {
		return nil, err
	}

	data := &Word{
		ID:           uint(word_id),
		Word:         word_fmt,
		PartOfSpeech: partOfSpeech,
		Definition:   string(matching_bytes[2]),
	}

	return data, nil
}

func LookupNoun(word string) (*Word, error) {
	return lookupType(word, noun_file, POS_Noun)
}

func LookupVerb(word string) (*Word, error) {
	return lookupType(word, verb_file, POS_Verb)
}

func LookupAdjective(word string) (*Word, error) {
	return lookupType(word, adjective_file, POS_Adjective)
}

func LookupAdverb(word string) (*Word, error) {
	return lookupType(word, adverb_file, POS_Adverb)
}

// Random lookup. Leave startsWith as "" for any.
func Rand(startsWith string, count uint) ([]*Word, error) {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch randomizer.Intn(4) {
	case 0:
		return RandNoun(startsWith, count)
	case 1:
		return RandVerb(startsWith, count)
	case 2:
		return RandAdjective(startsWith, count)
	case 3:
		return RandAdverb(startsWith, count)
	}
	return nil, fmt.Errorf("problem selecting part of speech")
}

// internal generic
func randType(file string, partOfSpeech POS, startsWith string, count uint) ([]*Word, error) {
	sw_fmt := strings.ToLower(strings.ReplaceAll(startsWith, " ", "_"))
	regex_string := strings.Replace(`(?im)^(\d+?) \d\d [^\\]+? START(\S+) .+?\| (.+)$`, "START", sw_fmt, 1)
	regex, err := regexp.Compile(regex_string)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	matching_bytes := regex.FindAllSubmatch(b, -1)

	if len(matching_bytes) == 0 {
		return nil, fmt.Errorf("word not found in file(s)")
	}

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	data := make([]*Word, count)

	for i := 0; i < int(count); i++ {
		randomIndex := randomizer.Intn(len(matching_bytes))

		word_id, err := strconv.ParseUint(string(matching_bytes[randomIndex][1]), 10, 64)
		if err != nil {
			return nil, err
		}

		data[i] = &Word{
			ID:           uint(word_id),
			Word:         string(matching_bytes[randomIndex][2]),
			PartOfSpeech: partOfSpeech,
			Definition:   string(matching_bytes[randomIndex][3]),
		}
	}

	return data, nil
}

func RandNoun(startsWith string, count uint) ([]*Word, error) {
	return randType(noun_file, POS_Noun, startsWith, count)
}

func RandVerb(startsWith string, count uint) ([]*Word, error) {
	return randType(verb_file, POS_Verb, startsWith, count)
}

func RandAdjective(startsWith string, count uint) ([]*Word, error) {
	return randType(adjective_file, POS_Adjective, startsWith, count)
}

func RandAdverb(startsWith string, count uint) ([]*Word, error) {
	return randType(adverb_file, POS_Adverb, startsWith, count)
}
