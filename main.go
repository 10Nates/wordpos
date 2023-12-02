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

	noun_file           = "wordnet/dict/data.noun"
	lines_in_noun_file  = 82221
	verb_file           = "wordnet/dict/data.verb"
	lines_in_verb_file  = 13818
	adjective_file      = "wordnet/dict/data.adj"
	lines_in_adj_file   = 18214
	adverb_file         = "wordnet/dict/data.adv"
	lines_in_adv_file   = 3654
	lines_in_disclaimer = 29
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

func GetPOS(text string) ([]POSSet, error) {
	return nil, fmt.Errorf("not implemented")
}

// Whole sentences

func GetNouns(text string) ([]*Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func GetVerbs(text string) ([]*Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func GetAdjectives(text string) ([]*Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func GetAdverbs(text string) ([]*Word, error) {
	return nil, fmt.Errorf("not implemented")
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
