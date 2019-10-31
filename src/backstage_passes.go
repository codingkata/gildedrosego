package main

type BackstagePasses struct {
}

func (bp BackstagePasses) UpdateQualityWhenExpiration(item *Item) {
	item.quality = 0
}

func (bp BackstagePasses) UpdateSellIn(item *Item) {
	item.sellIn--
}

func (bp BackstagePasses) UpdateQuality(item *Item) {
	if item.quality < 50 {
		item.quality++
		if item.sellIn < 11 {
			if item.quality < 50 {
				item.quality++
			}
		}
		if item.sellIn < 6 {
			if item.quality < 50 {
				item.quality++
			}
		}
	}
}

