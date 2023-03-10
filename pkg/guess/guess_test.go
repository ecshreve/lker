package guess_test

import (
	"strings"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/ecshreve/lker/pkg/guess"
	"github.com/stretchr/testify/assert"
)

func TestSanitize(t *testing.T) {
	clk := clock.NewMock()
	now := time.Date(2023, time.April, 20, 0, 0, 0, 0, clk.Now().Location())
	clk.Set(now)

	testcases := []struct {
		desc      string
		target    string
		expected  string
		expectErr bool
	}{
		{
			desc:     "emptyp string",
			target:   "",
			expected: "",
		},
		{
			desc:     "basic valid string",
			target:   "hello world",
			expected: "HELLOWORLD",
		},
		{
			desc:      "over max length",
			target:    strings.Repeat("a", guess.MAX_GUESS_LEN+1),
			expected:  "",
			expectErr: true,
		},
		{
			desc:      "includes profanity",
			target:    "some sentence butt ass other words",
			expected:  "some sentence **** *** other words",
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		actual, err := guess.Sanitize(tc.target)
		if tc.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.expected, actual)
	}
}

func TestCloud(t *testing.T) {
	clk := clock.NewMock()
	now := time.Date(2023, time.April, 20, 0, 0, 0, 0, clk.Now().Location())
	clk.Set(now)

	customArgs := guess.Args{
		Keys: []string{"hello", "world"},
	}
	cld := guess.NewCloud(customArgs)
	assert.Equal(t, customArgs.Keys, cld.Keys)
	assert.Equal(t, len(customArgs.Keys), len(cld.Items))

	def := guess.DefaultCloud()
	assert.Equal(t, strings.Split(guess.ALPHA, ""), def.Keys)
	assert.Equal(t, len(guess.ALPHA), len(def.Items))

	err := def.ProcessGuess("hello")
	assert.NoError(t, err)
	assert.Equal(t, 20, def.Items["L"].Rank)
	assert.Equal(t, 2, def.Items["L"].Count)
	assert.Equal(t, 1, def.Items["O"].Count)

	err = def.ProcessGuess("butt")
	assert.Error(t, err)
	assert.Equal(t, 20, def.Items["L"].Rank)
	assert.Equal(t, 2, def.Items["L"].Count)
	assert.Equal(t, 1, def.Items["O"].Count)
	assert.Equal(t, 0, def.Items["B"].Count)
	assert.Equal(t, 0, def.Items["T"].Count)

}
