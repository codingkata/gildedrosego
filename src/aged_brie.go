package main

type AgedBrie struct {
}

func (a AgedBrie) UpdateQualityWhenExpiration(item *Item) {
	if item.quality < 50 {
		item.quality++
	}
}

func (a AgedBrie) UpdateSellIn(item *Item) {
	item.sellIn--
}

func (a AgedBrie) UpdateQuality(item *Item) {
	if item.quality < 50 {
		item.quality++
	}
}

