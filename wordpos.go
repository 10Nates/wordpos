package wordpos

import (
	"bytes"
	_ "embed"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type POS string

const (
	POS_Noun           POS = "n"
	POS_Verb           POS = "v"
	POS_Adjective      POS = "adj"
	POS_Adverb         POS = "adv"
	POS_Other          POS = "o"
	file_header_length     = 29
)

var (
	//go:embed wordnet/dict/data.noun
	noun_file []byte
	//go:embed wordnet/dict/data.verb
	verb_file []byte
	//go:embed wordnet/dict/data.adj
	adjective_file []byte
	//go:embed wordnet/dict/data.adv
	adverb_file []byte
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
	// goal &{[0xc0000149c0 0xc000014a80 0xc000014b40 0xc000014c80 0xc000014e40] [0xc000015080 0xc000015140 0xc000015400] [0xc000015580 0xc000015640 0xc000015900] [0xc000015b00 0xc000015d40] [0xc000166140 0xc0001665c0]}
	// curr &{[0xc000014b80 0xc000014c40 0xc000014d00 0xc000015000 0xc000015280] [] [0xc0000151c0] [] [0xc000014ac0 0xc000014f40]}
	wordsarr := strings.Split(strings.ToLower(text), " ")
	words := []string{}
	wordscheck := make(map[string]bool, len(wordsarr))
	// make unique set
	for _, word := range wordsarr {
		if wordscheck[word] {
			continue
		}
		wordscheck[word] = true
		words = append(words, word)
	}

	ps := &POSSet{}

	// This is woefully inefficient and I hope to change this in the future
	for i := 0; i < len(words); i++ {
		wordl, err := Lookup(words[i], true)
		if err != nil {
			ps.Rest = append(ps.Rest, &Word{
				ID:           0,
				Word:         words[i],
				PartOfSpeech: POS_Other,
				Definition:   "",
			})
		} else {
			for _, word := range wordl {
				switch word.PartOfSpeech {
				case POS_Noun:
					ps.Nouns = append(ps.Nouns, word)
				case POS_Verb:
					ps.Verbs = append(ps.Verbs, word)
				case POS_Adjective:
					ps.Adjectives = append(ps.Adjectives, word)
				case POS_Adverb:
					ps.Adverbs = append(ps.Adverbs, word)
				}
			}
		}
	}

	return ps, nil
}

// Whole sentences

func GetNouns(text string) ([]*Word, error) {
	wordsarr := strings.Split(text, " ")
	words := []string{}
	wordscheck := make(map[string]bool, len(wordsarr))
	// make unique set
	for _, word := range wordsarr {
		if wordscheck[word] {
			continue
		}
		wordscheck[word] = true
		words = append(words, word)
	}
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
	wordsarr := strings.Split(text, " ")
	words := []string{}
	wordscheck := make(map[string]bool, len(wordsarr))
	// make unique set
	for _, word := range wordsarr {
		if wordscheck[word] {
			continue
		}
		wordscheck[word] = true
		words = append(words, word)
	}
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
	wordsarr := strings.Split(text, " ")
	words := []string{}
	wordscheck := make(map[string]bool, len(wordsarr))
	// make unique set
	for _, word := range wordsarr {
		if wordscheck[word] {
			continue
		}
		wordscheck[word] = true
		words = append(words, word)
	}
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
	wordsarr := strings.Split(text, " ")
	words := []string{}
	wordscheck := make(map[string]bool, len(wordsarr))
	// make unique set
	for _, word := range wordsarr {
		if wordscheck[word] {
			continue
		}
		wordscheck[word] = true
		words = append(words, word)
	}
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

func Lookup(word string, checkAll bool) ([]*Word, error) {
	words := []*Word{}
	data, err := LookupNoun(word)
	if err == nil && !checkAll {
		return []*Word{data}, nil
	} else if err == nil {
		words = append(words, data)
	} else if err != nil && err.Error() != "word not found in file(s)" {
		return nil, err
	}
	data, err = LookupVerb(word)
	if err == nil && !checkAll {
		return []*Word{data}, nil
	} else if err == nil {
		words = append(words, data)
	} else if err != nil && err.Error() != "word not found in file(s)" {
		return nil, err
	}
	data, err = LookupAdjective(word)
	if err == nil && !checkAll {
		return []*Word{data}, nil
	} else if err == nil {
		words = append(words, data)
	} else if err != nil && err.Error() != "word not found in file(s)" {
		return nil, err
	}
	data, err = LookupAdverb(word)
	if err == nil && !checkAll {
		return []*Word{data}, nil
	} else if err == nil {
		words = append(words, data)
	} else if err != nil && err.Error() != "word not found in file(s)" {
		return nil, err
	}

	return words, nil
}

// internal generic
func lookupType(word string, file *[]byte, partOfSpeech POS) (*Word, error) {
	word_fmt := strings.ToLower(strings.ReplaceAll(word, " ", "_"))
	// regex: matches the correct word, then captures the ID and definition.
	regex_string := strings.Replace(`(?im)^(\d+?) \d\d [^\\\n]+? WORD .+?\| (.+)$`, "WORD", word_fmt, 1)
	regex, err := regexp.Compile(regex_string)
	if err != nil {
		return nil, err
	}

	matching_bytes := regex.FindSubmatch(*file)

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
	return lookupType(word, &noun_file, POS_Noun)
}

func LookupVerb(word string) (*Word, error) {
	return lookupType(word, &verb_file, POS_Verb)
}

func LookupAdjective(word string) (*Word, error) {
	return lookupType(word, &adjective_file, POS_Adjective)
}

func LookupAdverb(word string) (*Word, error) {
	return lookupType(word, &adverb_file, POS_Adverb)
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
func randType(file *[]byte, partOfSpeech POS, startsWith string, count uint) ([]*Word, error) {
	sw_fmt := strings.ToLower(strings.ReplaceAll(startsWith, " ", "_"))
	regex_string := strings.Replace(`(?im)^(\d+?) \d\d . \d\d (\S+) .+?\| (.+)$`, "START", sw_fmt, 1)
	regex, err := regexp.Compile(regex_string)
	if err != nil {
		return nil, err
	}

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := *file
	if startsWith == "" { // significant speedup compared to full-file regex for nonspecific random words
		blist := bytes.Split(b, []byte{'\n'})
		blistcut := [][]byte{}
		for i := 0; i < int(count); i++ { // number of random items
			randomIndex := randomizer.Intn(len(blist)-file_header_length) + file_header_length
			blistcut = append(blistcut, blist[randomIndex])
		}
		b = bytes.Join(blistcut, []byte{'\n'})
	}

	matching_bytes := regex.FindAllSubmatch(b, -1)

	if len(matching_bytes) == 0 {
		return nil, fmt.Errorf("word not found in file(s)")
	}

	data := make([]*Word, count)

	for i := 0; i < int(count); i++ {
		randomIndex := randomizer.Intn(len(matching_bytes))
		if startsWith == "" {
			randomIndex = i // fix double cycle
		}

		word_id, err := strconv.ParseUint(string(matching_bytes[randomIndex][1]), 10, 64)
		if err != nil {
			return nil, err
		}

		data[i] = &Word{
			ID:           uint(word_id),
			Word:         sw_fmt + string(matching_bytes[randomIndex][2]),
			PartOfSpeech: partOfSpeech,
			Definition:   string(matching_bytes[randomIndex][3]),
		}
	}

	return data, nil
}

func RandNoun(startsWith string, count uint) ([]*Word, error) {
	return randType(&noun_file, POS_Noun, startsWith, count)
}

func RandVerb(startsWith string, count uint) ([]*Word, error) {
	return randType(&verb_file, POS_Verb, startsWith, count)
}

func RandAdjective(startsWith string, count uint) ([]*Word, error) {
	return randType(&adjective_file, POS_Adjective, startsWith, count)
}

func RandAdverb(startsWith string, count uint) ([]*Word, error) {
	return randType(&adverb_file, POS_Adverb, startsWith, count)
}
