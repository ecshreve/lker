package guess

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// MAX_GUESS_LEN represents the length of the longest allowed guess string.
const MAX_GUESS_LEN int = 220

func Sanitize(gstr string) (string, error) {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	if len(gstr) > MAX_GUESS_LEN {
		return "", oops.Errorf("input string too long")
	}

	// Strip any html tags from string.
	var sanitizerInstance = bluemonday.StrictPolicy()
	sanStr := sanitizerInstance.Sanitize(gstr)

	// Custom sanitizing.
	var chars []rune
	for _, s := range strings.ToUpper(sanStr) {
		if int(s) < int('A') || int(s) > int('Z') {
			continue
		}
		chars = append(chars, s)
	}

	return string(chars), nil
}
