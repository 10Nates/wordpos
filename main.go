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

// internal
func lookupType(word string, file string, partOfSpeech POS) (*Word, error) {
	word_fmt := strings.ToLower(strings.ReplaceAll(word, " ", "_"))
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
