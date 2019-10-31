package main

type NormalItem struct {
}

func (n NormalItem) UpdateQuality(item *Item) {
	if item.quality > 0 {
		item.quality--
	}
}

