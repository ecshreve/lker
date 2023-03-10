package guess

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysOrDefault(t *testing.T) {
	testcases := []struct {
		desc     string
		target   []string
		expected []string
	}{
		{
			desc:     "nil list",
			target:   nil,
			expected: strings.Split(ALPHA, ""),
		},
		{
			desc:     "empty list",
			target:   []string{},
			expected: strings.Split(ALPHA, ""),
		},
		{
			desc:     "valid list",
			target:   []string{"hello", "world"},
			expected: []string{"hello", "world"},
		},
	}

	for _, tc := range testcases {
		actual := keysOrDefault(tc.target)
		assert.Equal(t, tc.expected, actual)
	}
}
