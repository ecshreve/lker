package guess

import (
	"sort"
	"strings"

	"github.com/samsarahq/go/oops"
)

// MAX_ITEM_RANK represents the ceiling of our ranking scale.
const MAX_ITEM_RANK int = 20
const ALPHA string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Item represents a single member of a Cloud.
type Item struct {
	Val   string
	Count int
	Rank  int
}

func initItems(keys []string) map[string]*Item {
	items := make(map[string]*Item)
	for _, k := range keys {
		items[k] = &Item{
			Val:   k,
			Count: 0,
			Rank:  0,
		}
	}
	return items
}

func keysOrDefault(k []string) []string {
	if len(k) > 0 {
		return k
	}
	return strings.Split(ALPHA, "")
}

// Cloud represents a "word cloud", containing some set of keys, each key having
// some count of occurrences, resulting in a natural ability to rank the order of
// key occurrence.
type Cloud struct {
	Keys  []string
	Items map[string]*Item
}

// Args represent available arguments when creating a Cloud.
type Args struct {
	Keys []string
}

// DefaultCloud is a convenience function to return a Cloud initialized with a
// a default set of keys.
func DefaultCloud() *Cloud {
	return NewCloud(Args{})
}

// NewCloud initializes an instance of Cloud with the letters of the alphabet
// as its set of keys.
func NewCloud(a Args) *Cloud {
	keys := keysOrDefault(a.Keys)
	return &Cloud{
		Keys:  keys,
		Items: initItems(keys),
	}
}

// ProcessGuess takes an input string, sanitizes it, applies updates to the word
// cloud according to the contents of the guess.
func (c *Cloud) ProcessGuess(gstr string) error {
	g, err := Sanitize(gstr)
	if err != nil {
		return oops.Wrapf(err, "unable to sanitize")
	}

	for _, s := range g {
		if _, ok := c.Items[string(s)]; ok {
			c.Items[string(s)].Count += 1
		}
	}

	c.UpdateRanks()
	return nil
}

// UpdateRanks iterates through the items in the cloud and updates each item's
// rank field according to current counts.
func (c *Cloud) UpdateRanks() {
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
