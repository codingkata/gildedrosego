package main

type ItemUpdater interface {
	UpdateQuality(item *Item)
	UpdateSellIn(item *Item)
	UpdateQualityWhenExpiration(item *Item)
}

func createUpdater(item *Item) ItemUpdater {
	if item.name == "Aged Brie" {
		return AgedBrie{}
	} else if item.name == "Backstage passes to a TAFKAL80ETC concert" {
		return BackstagePasses{}
	} else if item.name == "Sulfuras, Hand of Ragnaros" {
		return Sulfuras{}
	} else if item.name == "Conjured" {
		return Conjured{}
	} else {
		return NormalItem{}
	}
}
