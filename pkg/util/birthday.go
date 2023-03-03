package util

import (
	"math"
	"time"

	log "github.com/sirupsen/logrus"
)

func msUntilBday(bday time.Time) int64 {
	log.Info("---> msUntilBday() - enter")
	defer log.Info("<--- msUntilBday() - exit")

	today := time.Now()
	nextBirthday := time.Date(today.Year(), bday.Month(), bday.Day(), 0, 0, 0, 0, bday.Location())
	if nextBirthday.Before(today) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}
	ms := int64(nextBirthday.Sub(today).Milliseconds())
	return ms
}

func nearestBirthday(spiceGirls map[string]string) (string, int64) {
	log.Info("---> nearestBirthday() - enter")
	defer log.Info("<--- nearestBirthday() - exit")

	var nearestName string
	nearest := int64(math.MaxInt64)
	for name, birthdayStr := range spiceGirls {
		birthday, _ := time.Parse("01/02/2006", birthdayStr)
		ms := msUntilBday(birthday)
		if ms < nearest {
			nearestName = name
			nearest = ms
		}
	}
	return nearestName, nearest
}

func GetNearestMs() int64 {
	log.Info("---> GetNearestMs() - enter")
	defer log.Info("<--- GetNearestMs() - exit")

	spiceGirls := map[string]string{
		"Victoria Beckham": "04/17/1974",
		"Melanie Brown":    "05/29/1975",
		"Emma Bunton":      "01/21/1976",
		"Melanie Chisholm": "01/12/1974",
		"Geri Halliwell":   "08/06/1972",
	}
	_, nearest := nearestBirthday(spiceGirls)
	return nearest
}
