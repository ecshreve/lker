package util

import (
	"math"
	"time"
)

func msUntilBday(bday time.Time) int {
	today := time.Now()
	nextBirthday := time.Date(today.Year(), bday.Month(), bday.Day(), 0, 0, 0, 0, bday.Location())
	if nextBirthday.Before(today) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}
	ms := int(nextBirthday.Sub(today).Milliseconds() / 1000)
	return ms
}

func nearestBirthday(spiceGirls map[string]string) (string, int) {
	var nearestName string
	nearest := math.MaxInt32
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

func GetNearestMs() int {
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
