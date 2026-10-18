package task

import (
	"errors"
	"strings"
)

type Processor interface {
	Process(input string) (Result, error)
}

type TextProcessor struct{}

func (tp TextProcessor) Process(input string) (Result, error) {
	if strings.TrimSpace(input) == "" {
		return Result{}, errors.New("cannot process empty text")
	}

	words := strings.Fields(input)
	charCount := 0
	vowelCount := 0

	// loop over lowercase runes so we dont have to check uppercase vowels separately
	lower := strings.ToLower(input)
	for _, r := range lower {
		if r != ' ' && r != '\n' && r != '\t' {
			charCount++
		}
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowelCount++
		}
	}

	return Result{
		WordCount:  len(words),
		CharCount:  charCount,
		VowelCount: vowelCount,
	}, nil
}
