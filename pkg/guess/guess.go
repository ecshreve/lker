package guess

import (
	"strings"

	goaway "github.com/TwiN/go-away"
	"github.com/microcosm-cc/bluemonday"
	"github.com/samsarahq/go/oops"
)

// MAX_GUESS_LEN represents the length of the longest allowed guess string.
const MAX_GUESS_LEN int = 220

func Sanitize(gstr string) (string, error) {
	if len(gstr) > MAX_GUESS_LEN {
		return "", oops.Errorf("input string too long")
	}

	// Strip any html tags from string.
	var sanitizerInstance = bluemonday.StrictPolicy()
	sanStr := sanitizerInstance.Sanitize(gstr)

	// Check for profanity.
	if goaway.IsProfane(sanStr) {
		return goaway.Censor(sanStr), oops.Errorf("profanity detected in guess")
	}

	// Custom sanitizing.
	var chars []rune
	for _, s := range strings.ToUpper(sanStr) {
		if int(s) < int('A') || int(s) > int('Z') {
			continue
		}
		chars = append(chars, s)
	}

	// Double check for profanity.
	retStr := string(chars)
	return goaway.Censor(retStr), nil
}
