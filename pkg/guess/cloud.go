package guess

import (
	"sort"
	"strings"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// MAX_ITEM_RANK represents the ceiling of our ranking scale.
const MAX_ITEM_RANK int = 20

type Item struct {
	Val   string
	Count int
	Rank  int
}

type Cloud struct {
	Keys  []string
	Items map[string]*Item
}

func NewCloud() *Cloud {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	alphabet := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	items := make(map[string]*Item)
	for _, k := range alphabet {
		items[k] = &Item{
			Val:   k,
			Count: 0,
			Rank:  0,
		}
	}
	return &Cloud{
		Keys:  alphabet,
		Items: items,
	}
}

func (c *Cloud) ProcessGuess(gstr string) error {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	g, err := Sanitize(gstr)
	if err != nil {
		return oops.Wrapf(err, "unable to sanitize")
	}

	for _, s := range strings.ToUpper(g) {
		if int(s) < int('A') || int(s) > int('Z') {
			continue
		}

		if _, ok := c.Items[string(s)]; !ok {
			continue
		}

		c.Items[string(s)].Count += 1
	}

	c.UpdateRanks()
	return nil
}

func (c *Cloud) UpdateRanks() {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	var items []*Item
	for _, i := range c.Items {
		items = append(items, i)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})

	for ind, item := range items {
		rankVal := MAX_ITEM_RANK - ind
		if item.Count <= 0 || ind > MAX_ITEM_RANK {
			rankVal = 1
		}
		c.Items[item.Val].Rank = rankVal
	}
}
