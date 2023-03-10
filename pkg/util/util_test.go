package util_test

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/ecshreve/lker/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestGetNearestNms(t *testing.T) {
	clk := clock.NewMock()

	now := time.Date(2023, time.April, 20, 0, 0, 0, 0, clk.Now().Location())
	clk.Set(now)

	testcases := []struct {
		desc      string
		target    []string
		expected  int64
		expectErr bool
	}{
		{
			desc:     "same value as now, expect 0",
			target:   []string{"04/20/2023"},
			expected: 0,
		},
		{
			desc:     "now and future value, expect 0",
			target:   []string{"04/20/2023", "04/30/2023"},
			expected: 0,
		},
		{
			desc:     "past and future value, expect 1d in ms",
			target:   []string{"04/19/2023", "04/30/2023"},
			expected: 86400000,
		},
		{
			desc:     "use real consts, expect 259200000 in ms",
			target:   util.SpiceGirlBirthdates,
			expected: 259200000,
		},
		{
			desc:     "pass nil slice, expect 259200000 in ms",
			target:   nil,
			expected: 259200000,
		},
		{
			desc:      "bad date, expect error and -1 return",
			target:    []string{"04/90/2023"},
			expected:  -1,
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		actual, err := util.GetNearestMs(clk, tc.target)
		if tc.expectErr {
			assert.Error(t, err)
		}
		assert.Equal(t, tc.expected, actual)
	}
}
