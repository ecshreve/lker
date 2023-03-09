package util

import (
	"math"
	"time"

	log "github.com/sirupsen/logrus"
)

func timeNowDiff(t time.Time) int64 {
	log.Trace("---> - enter")
	defer log.Trace("<--- -  exit")

	today := time.Now()
	target := time.Date(today.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return int64(math.Abs(float64(target.UnixMilli() - today.UnixMilli())))
}

func nearest(dateStrings []string) int64 {
	log.Trace("---> - enter")
	defer log.Trace("<--- -  exit")

	nearest := int64(math.MaxInt64)
	for _, d := range dateStrings {
		target, _ := time.Parse("01/02/2006", d)
		if ms := timeNowDiff(target); ms < nearest {
			nearest = ms
		}
	}
	return nearest
}

func GetNearestMs() int64 {
	log.Trace("---> - enter")
	defer log.Trace("<--- -  exit")

	datesToDiff := []string{
		"04/17/1974", //" Victoria Beckham":
		"05/29/1975", // "Melanie Brown":
		"01/21/1976", // "Emma Bunton":
		"01/12/1974", //		"Melanie Chisholm":
		"08/06/1972", //		"Geri Halliwell":
	}
	n := nearest(datesToDiff)
	return n
}
