package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 10, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	fmt.Println(items)
	GlidedRose()
	fmt.Println(items)
}

func GlidedRose() {
	myitems := items
	UpdateQuality(myitems)
}

func UpdateQuality(myitems []Item) {
	for i := 0; i < len(myitems); i++ {

		if myitems[i].name != "Aged Brie" && myitems[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if myitems[i].quality > 0 {
				if myitems[i].name != "Sulfuras, Hand of Ragnaros" {
					myitems[i].quality = myitems[i].quality - 1
				}
			}
		} else {
			if myitems[i].quality < 50 {
				myitems[i].quality = myitems[i].quality + 1
				if myitems[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if myitems[i].sellIn < 11 {
						if myitems[i].quality < 50 {
							myitems[i].quality = myitems[i].quality + 1
						}
					}
					if myitems[i].sellIn < 6 {
						if myitems[i].quality < 50 {
							myitems[i].quality = myitems[i].quality + 1
						}
					}
				}
			}
		}

		if myitems[i].name != "Sulfuras, Hand of Ragnaros" {
			myitems[i].sellIn = myitems[i].sellIn - 1
		}

		if myitems[i].sellIn < 0 {
			if myitems[i].name != "Aged Brie" {
				if myitems[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if myitems[i].quality > 0 {
						if myitems[i].name != "Sulfuras, Hand of Ragnaros" {
							myitems[i].quality = myitems[i].quality - 1
						}
					}
				} else {
					myitems[i].quality = myitems[i].quality - myitems[i].quality
				}
			} else {
				if myitems[i].quality < 50 {
					myitems[i].quality = myitems[i].quality + 1
				}
			}
		}
	}
}
