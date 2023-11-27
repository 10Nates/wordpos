package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func GetPOS(text string) (*[]POS, error) {
	return nil, fmt.Errorf("not implemented")
}

// Whole sentences

func GetNouns(text string) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func GetVerbs(text string) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func GetAdjectives(text string) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func GetAdverbs(text string) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

// Word only

func IsNoun(word string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func IsVerb(word string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func IsAdjective(word string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func IsAdverb(word string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func Lookup(word string) (*Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func LookupNoun(word string) (*Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func LookupVerb(word string) (*Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func LookupAdjective(word string) (*Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func LookupAdverb(word string) (*Word, error) {
	word_fmt := strings.ToLower(strings.ReplaceAll(word, " ", "_"))
	regex_string := strings.Replace(`(?im)^(\d+?) \d\d r [^\\]+? WORD .+?\| (.+)$`, "WORD", word_fmt, 1)
	regex, err := regexp.Compile(regex_string)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(adverb_file)
	if err != nil {
		return nil, err
	}

	matching_bytes := regex.FindSubmatch(b)

	if len(matching_bytes) == 0 {
		return nil, fmt.Errorf("word not found in adverbs")
	}

	word_id, err := strconv.ParseUint(string(matching_bytes[1]), 10, 64)
	if err != nil {
		return nil, err
	}

	data := &Word{
		ID:           uint(word_id),
		Word:         word_fmt,
		PartOfSpeech: POS_Adverb,
		Definition:   string(matching_bytes[2]),
	}

	return data, nil
}

// Random lookup. Leave startsWith as "" for any.

func Rand(startsWith string, count uint) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func RandNoun(startsWith string, count uint) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func RandVerb(startsWith string, count uint) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func RandAdjective(startsWith string, count uint) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}

func RandAdverb(startsWith string, count uint) (*[]Word, error) {
	return nil, fmt.Errorf("not implemented")
}
