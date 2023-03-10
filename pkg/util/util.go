package util

import (
	"math"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/samsarahq/go/oops"
)

// SpiceGirlBirthdates is a list of date strings containing the birthdates
// of each of the original spice girls.
var SpiceGirlBirthdates = []string{
	"04/17/1974", // "Victoria Beckham":
	"05/29/1975", // "Melanie Brown":
	"01/21/1976", // "Emma Bunton":
	"01/12/1974", // "Melanie Chisholm":
	"08/06/1972", // "Geri Halliwell":
}

// timeDiff returns the absolute ms difference between the two given times. In
// this implementationi t1 is expected to be Now().
func timeDiff(t1, t2 time.Time) int64 {
	target := time.Date(t1.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute(), t2.Second(), t2.Nanosecond(), t2.Location())
	return int64(math.Abs(float64(target.UnixMilli() - t1.UnixMilli())))
}

// nearest returns the millisecond value of the time until the closes date on the
// given list of date strings.
func nearest(t time.Time, dts []time.Time) int64 {
	nearest := int64(math.MaxInt64)
	for _, d := range dts {
		if ms := timeDiff(t, d); ms < nearest {
			nearest = ms
		}
	}

	return nearest
}

// GetNearestMs returns the value in ms of the smallest time between now
// and a spice girl's birthday.
func GetNearestMs(clk clock.Clock, args []string) (int64, error) {
	if args == nil {
		args = SpiceGirlBirthdates
	}

	dateTimesForDiff := make([]time.Time, len(args))
	for i, dstr := range args {
		if target, err := time.Parse("01/02/2006", dstr); err == nil {
			dateTimesForDiff[i] = target
		} else {
			return -1, oops.Wrapf(err, "parsing date: %s", dstr)
		}
	}

	return nearest(clk.Now(), dateTimesForDiff), nil
}
