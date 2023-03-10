package util

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
)

func TestTimeNowDiff(t *testing.T) {
	clk := clock.NewMock()
	now := time.Date(2023, time.April, 20, 0, 0, 0, 0, clk.Now().Location())
	clk.Set(now)

	testcases := []struct {
		desc     string
		target   time.Time
		expected int64
	}{
		{
			desc:     "same value as now, expect 0",
			target:   now,
			expected: 0,
		},
		{
			desc:     "future value, expect 1h in ms",
			target:   now.Add(time.Hour),
			expected: 3600000,
		},
		{
			desc:     "future value, expect 24h in ms",
			target:   now.Add(time.Hour * 24),
			expected: 86400000,
		},
		{
			desc:     "past value, expect 1h in ms",
			target:   now.Add(-time.Hour),
			expected: 3600000,
		},
		{
			desc:     "future value, expect 24h in ms",
			target:   now.Add(-time.Hour * 24),
			expected: 86400000,
		},
	}

	for _, tc := range testcases {
		actual := timeDiff(now, tc.target)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestNearest(t *testing.T) {
	clk := clock.NewMock()
	now := time.Date(2023, time.April, 20, 0, 0, 0, 0, clk.Now().Location())
	clk.Set(now)

	testcases := []struct {
		desc     string
		target   []time.Time
		expected int64
	}{
		{
			desc:     "same value as now, expect 0",
			target:   []time.Time{now},
			expected: 0,
		},
		{
			desc:     "future value, expect 1h in ms",
			target:   []time.Time{now.Add(time.Hour)},
			expected: 3600000,
		},
		{
			desc:     "now and future value, expect 0",
			target:   []time.Time{now, now.Add(time.Hour)},
			expected: 0,
		},
		{
			desc:     "past and future value, expect 30m in ms",
			target:   []time.Time{now.Add(-time.Minute * 30), now.Add(time.Hour)},
			expected: 1800000,
		},
		{
			desc:     "large past and future values, expect 0",
			target:   []time.Time{now.Add(-time.Hour * 24 * 10), now.Add(time.Hour * 24 * 100)},
			expected: 864000000,
		},
	}

	for _, tc := range testcases {
		actual := nearest(now, tc.target)
		assert.Equal(t, tc.expected, actual)
	}
}
