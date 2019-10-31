package main

type NormalItem struct {
}

func (n NormalItem) UpdateQualityWhenExpiration(item *Item) {
	if item.quality > 0 {
		item.quality--
	}
}

func (n NormalItem) UpdateSellIn(item *Item) {
	item.sellIn--
}

func (n NormalItem) UpdateQuality(item *Item) {
	if item.quality > 0 {
		item.quality--
	}
}

