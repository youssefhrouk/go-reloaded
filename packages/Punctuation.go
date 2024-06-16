package reloaded

import (
	"regexp"
	"strings"
)

func Punctuation(words []string) string {
	sentence := strings.Join(words, " ")
	// Remove spaces around punctuation except for special cases
	matchPunc1 := regexp.MustCompile(`\s+([,.:;!?])`)
	sentence = matchPunc1.ReplaceAllString(sentence, "$1")

	// Add spaces after punctuation if not followed by special punctuation
	matchPunc2 := regexp.MustCompile(`([,.!?;:])([^,.!?;:\s])`)
	sentence = matchPunc2.ReplaceAllString(sentence, "$1 $2")

	matchA := regexp.MustCompile(`(?i)\b(a)\s+([aeiouh])`)
	for matchA.MatchString(sentence) {
		sentence = matchA.ReplaceAllString(sentence, "${1}n $2")
	}
	return sentence
}
