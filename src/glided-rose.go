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
	UpdateQualityForAll(items)
}

func UpdateQualityForAll(myitems []Item) {
	for i := 0; i< len(myitems);i++  {
		passOneDay(&myitems[i])
	}
}

func passOneDay(item *Item) {
	updater:=createUpdater(item)
	updater.UpdateQuality(item)
	updater.UpdateSellIn(item)
	if item.sellIn < 0 {
		updater.UpdateQualityWhenExpiration(item)
	}
}