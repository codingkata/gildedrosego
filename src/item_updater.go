package main

type ItemUpdater interface {
	UpdateQuality(item *Item)
}

func createUpdater(item *Item) ItemUpdater {
	if item.name == "Aged Brie" {
		return AgedBrie{}
	} else if item.name == "Backstage passes to a TAFKAL80ETC concert" {
		return BackstagePasses{}
	} else if item.name == "Sulfuras, Hand of Ragnaros" {
		return Sulfuras{}
	} else {
		 return NormalItem{}
	}
}
