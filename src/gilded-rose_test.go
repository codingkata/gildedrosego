package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

- 每种物品都有一个：
	- 'name'值。
	- `SellIn`值，表示我们要在多少天之内把物品卖出去，即销售期
	- `Quality`值，表示物品的品质，它永远不会为负值

- 每天结束时，系统通常都会降低各种物品的'SellIn'和'Quality'，规则如下：
	- 每过去一天，sellIn值会减少1，quality 也会减少1
	- 一旦销售期过了（sellIn<0），`Quality`会以双倍的速度加速下降

- 但也有一些特殊的物品，其quality的变化与一般的物品不同，例如：

	- "Aged Brie"
		- `quality`会随着时间推移而提高，每天提高1
		- `quality`也永远不会超过50

	- "Sulfuras, Hand of Ragnaros"
		- 永远也不会过期
		- `Quality`也永远不会变化，一直是80

	- "Backstage passes to a TAFKAL80ETC concert"
		-`quality`会随着时间推移而提高，升高的方式如下，当销售期：
			- 大于10天时，`Quality`不发生变化;
			- 还剩10天或更少的时候，品质`Quality`每天提高2;
			- 还剩5天或更少的时候，`Quality`每天提高3；
			- 一旦过期，品质就会降为0

* 新需求如下：
	* 新增加一种物品，名为Conjured，其遵守如下规则：
		* `Quality`下降速度比一般物品快一倍
		* 其它规则与一般物品相同
 */

func Test_GildedRose(t *testing.T) {
	_, items:=newItem("Aged Brie",1,1)
	UpdateQuality(items)
	assert.Equal(t, items[0].sellIn,0)
}
func Test_sellIn_should_decrease_1_per_day(t *testing.T) {
	names:=[] string{
		"normal item",
		"Backstage passes to a TAFKAL80ETC concert",
		"Aged Brie",
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T){
			sellIn := 10
			anyQuality :=12
			_,items:=newItem(name,sellIn,anyQuality)
			UpdateQuality(items)
			assert.Equal(t,sellIn-1,items[0].sellIn)
		})
	}
}
func Test_quality_should_decrease_one_per_day_until_expiration_for_normal_item(t *testing.T) {
	qualities:=[] int{
		5,
		1,
	}
	for _, quality := range qualities {
		t.Run("quality of normal should decrease by 1", func(t *testing.T){
			unexpiredDay := 10
			_,items:=newItem("normal", unexpiredDay,quality)
			UpdateQuality(items)
			assert.Equal(t,quality-1,items[0].quality)
		})
	}
}

func Test_quality_should_decrease_by_2_after_expiration_for_normal_item(t *testing.T) {
	qualities:=[] int{
		5,
		4,
	}
	for _, quality := range qualities {
		expiredDay := 0
		t.Run("quality of normal should decrease by 1", func(t *testing.T){
			_,items:=newItem("normal", expiredDay,quality)
			UpdateQuality(items)
			assert.Equal(t,quality-2,items[0].quality)
			assert.Equal(t,expiredDay-1,items[0].sellIn)
		})
	}
}

func Test_quality_should_not_be_negative_for_every_item(t *testing.T) {
	qualities:=[] int{
		0,
	}
	for _, quality := range qualities {
		t.Run("quality_should_not_be_negative_for_every_item", func(t *testing.T){
			sellIn :=10
			_,items:=newItem("normal",sellIn,quality)
			UpdateQuality(items)
			assert.Equal(t,0,items[0].quality)
		})
	}
}

func Test_Aged_Brie(t *testing.T) {
	tests:=map[string] struct{
		sellIn int
		quality int
		expectSellIn int
		expectquality int
	}{
		"quality_should_not_more_than_50"	:{10,50,9,50},
		"quality of normal should increase by 1"	:{5,10,4,11},

	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T){
			_,items:=newItem("Aged Brie",testcase.sellIn,testcase.quality)
			UpdateQuality(items)
			assert.Equal(t,testcase.expectquality,items[0].quality)
			assert.Equal(t,testcase.expectSellIn,items[0].sellIn)
		})
	}
}

func Test_Sulfuras_should_not_expired_and_quality_is_80_forever(t *testing.T) {
	tests:=map[string] struct{
		sellIn int
		quality int
		expectSellIn int
		expectquality int
	}{
		"quality_should_not_change_for Sulfuras"	:{20,80,20,80},
		"sellIn_should_not_be_negative_for Sulfuras":{0,80,0,80},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T){
			_,items:=newItem("Sulfuras, Hand of Ragnaros",testcase.sellIn,testcase.quality)
			UpdateQuality(items)
			assert.Equal(t,testcase.expectquality,items[0].quality)
			assert.Equal(t,testcase.expectSellIn,items[0].sellIn)
		})
	}
}

func Test_quality_should_increase_according_to_sellIn_for_Backstage_passes(t *testing.T) {
	tests:=map[string] struct{
		sellIn int
		quality int
		expectquality int
	}{
		"sellin is more than 10 days，quality should not change"	:{12,50,50},
		"sellin is less than 10 days,quality should decrease by 2"	:{10,40,42},
		"sellin is less than 5 days, quality should decrease by 3"	:{5,30,33},
		"sellin is less than 0 day, quality should be 0"			:{0,30,0},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T){
			_,items:=newItem("Backstage passes to a TAFKAL80ETC concert",testcase.sellIn,testcase.quality)
			UpdateQuality(items)
			assert.Equal(t,testcase.expectquality,items[0].quality)
		})
	}
}

func newItem(name string, sellIn int, quality int) (item *Item, items []Item){
	item = &Item{name, sellIn, quality}
	items = []Item{*item}
	return item, items
}