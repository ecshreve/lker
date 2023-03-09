package guess

import (
	"sort"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type ItemRank int

const MAX_ITEM_RANK int = 20
const MAX_GUESS_LEN int = 220

func (i *Item) GetRank(r int) int {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	if i.Count <= 0 {
		return 1
	}
	if r == 0 {
		return MAX_ITEM_RANK
	}
	if r > int(MAX_ITEM_RANK) {
		return 1
	}
	return MAX_ITEM_RANK - r
}

func Validate(gstr string) error {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	if len(gstr) > MAX_GUESS_LEN {
		return oops.Errorf("too long")
	}

	return nil
}

func Sanitize(gstr string) (string, error) {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	if err := Validate(gstr); err != nil {
		return "", oops.Wrapf(err, "validating")
	}

	var sanitizerInstance = bluemonday.StrictPolicy()
	sanStr := sanitizerInstance.Sanitize(gstr)

	var chars []rune
	for _, s := range strings.ToUpper(sanStr) {
		if int(s) < int('A') || int(s) > int('Z') {
			continue
		}
		chars = append(chars, s)
	}

	return string(chars), nil
}

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

func (c *Cloud) ProcessGuess(g string) error {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	if len(g) > MAX_GUESS_LEN {
		return oops.Errorf("guess too long")
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
		item.Rank = item.GetRank(ind)
		c.Items[item.Val] = item
	}
}
