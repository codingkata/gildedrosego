package main

type AgedBrie struct {
}

func (a AgedBrie) UpdateQuality(item *Item) {
	if item.quality < 50 {
		item.quality++
	}
}

