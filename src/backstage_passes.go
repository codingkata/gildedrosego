package main

type BackstagePasses struct {
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

