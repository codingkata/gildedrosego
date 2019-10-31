package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 10, 20},
	{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	fmt.Println(items)
	GlidedRose()
	fmt.Println(items)
}

func GlidedRose() {
	myitems := items
	UpdateQualityForAll(myitems)
}

func UpdateQualityForAll(myitems []Item) {
	for i := 0; i < len(myitems); i++ {
		item := &myitems[i]
		passOneDay(item)
	}
}

func passOneDay(item *Item) {
	UpdateItemQuality(item)
	UpdateSellIn(item)
	if item.sellIn < 0 {
		updateQualityWhenExpiration(item)
	}
}

func UpdateItemQuality(item *Item) {
	if item.name == "Aged Brie" {
		a:=AgedBrie{}
		a.UpdateQuality(item)
	} else if item.name == "Backstage passes to a TAFKAL80ETC concert" {
		b:= BackstagePasses{}
		b.UpdateQuality(item)
	} else if item.name == "Sulfuras, Hand of Ragnaros" {
		s:= Sulfuras{}
		s.UpdateQuality(item)
	} else {
		n:=NormalItem{}
		n.UpdateQuality(item)
	}
}
type NormalItem struct{

}
func(n NormalItem) UpdateQuality(item *Item) {
	if item.quality > 0 {
		item.quality--
	}
}

func UpdateSellIn(item *Item) {
	if item.name != "Sulfuras, Hand of Ragnaros" {
		item.sellIn--
	}
}

func updateQualityWhenExpiration(item *Item) {
	if item.name != "Aged Brie" {
		if item.name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.quality > 0 {
				if item.name != "Sulfuras, Hand of Ragnaros" {
					item.quality--
				}
			}
		} else {
			item.quality = 0
		}
	} else {
		if item.quality < 50 {
			item.quality++
		}
	}
}
