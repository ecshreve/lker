package util

import (
	"math"
	"time"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// timeNowDiff returns the absolute ms difference between the given time and Now().
func timeNowDiff(t time.Time) int64 {
	target := time.Date(time.Now().Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return int64(math.Abs(float64(target.UnixMilli() - time.Now().UnixMilli())))
}

// nearest returns the millisecond value of the smalles difference between now
// and the given times.
func nearest(dateStrings []string) int64 {
	nearest := int64(math.MaxInt64)
	for _, d := range dateStrings {
		target, err := time.Parse("01/02/2006", d)
		if err != nil {
			log.Error(oops.Wrapf(err, "parsing time"))
			continue
		}

		if ms := timeNowDiff(target); ms < nearest {
			nearest = ms
		}
	}

	return nearest
}

// GetNearestMs returns the value in ms of the smallest time between now
// and a spice girl's birthday.
func GetNearestMs() int64 {
	datesToDiff := []string{
		"04/17/1974", // "Victoria Beckham":
		"05/29/1975", // "Melanie Brown":
		"01/21/1976", // "Emma Bunton":
		"01/12/1974", // "Melanie Chisholm":
		"08/06/1972", // "Geri Halliwell":
	}
	n := nearest(datesToDiff)
	return n
}
