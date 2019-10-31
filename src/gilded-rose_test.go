package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

- 每种物品都有一个：
	- 'name'值。
	- `SellIn`值，表示我们要在多少天之内把物品卖出去，即销售期
	- `Quality`值，表示物品的品质

- 每天结束时，系统通常都会降低各种物品的'SellIn'和'Quality'，规则如下：
	- 每过去一天，sellIn值会减少1，quality 也会减少1
	- 一旦销售期过了（sellIn<0），`Quality`会以双倍的速度加速下降
	- `Quality`永远不能为负值

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
func Test_All_items_sellIn_should_decrease_1_per_day_except_for_Sulfuras(t *testing.T) {
	tests := [] struct {
		input    Item
		expected Item
	}{
		{Item{"normal item", 1, 1}, Item{"normal item", 0, 0}},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 1, 1}, Item{"Backstage passes to a TAFKAL80ETC concert", 0, 0}},
		{Item{"Aged Brie", 1, 1}, Item{"Aged Brie", 0, 0}},
		{Item{"Sulfuras, Hand of Ragnaros", 1, 1}, Item{"Sulfuras, Hand of Ragnaros", 1, 1}},
	}
	for _, item := range tests {
		t.Run(item.input.name, func(t *testing.T) {
			_, items := newItem(item.input.name, item.input.sellIn, item.input.quality)
			UpdateQualityForAll(items)
			assert.Equal(t, item.expected.sellIn, items[0].sellIn)
		})
	}
}

func Test_All_items_quality_should_not_negative_at_all(t *testing.T) {
	names := [] string{
		"normal item",
		"Backstage passes to a TAFKAL80ETC concert",
		"Aged Brie",
		"Sulfuras, Hand of Ragnaros",
	}
	sellIns := [] int{
		0, 1,
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			for _, eachSellIn := range sellIns {
				nonNegativeQuality := 0
				_, items := newItem(name, eachSellIn, nonNegativeQuality)
				UpdateQualityForAll(items)
				assert.GreaterOrEqual(t,  items[0].quality,0)
			}
		})
	}
}

func Test_Normal_item_quality_should_decrease_1_per_day_until_expired(t *testing.T) {
	names := [] string{
		"normal item",
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			sellIn := 0
			nonNegativeQuality := 0
			_, items := newItem(name, sellIn, nonNegativeQuality)
			UpdateQualityForAll(items)
			assert.Equal(t, sellIn-1, items[0].sellIn)
			assert.Equal(t, 0, items[0].quality)
		})
	}
}

func Test_Normal_item_quality_should_decrease_by_2_once_expiration(t *testing.T) {
	qualities := [] int{
		5,
		4,
	}
	for _, quality := range qualities {
		expiredDay := 0
		t.Run("quality of normal should decrease by 1", func(t *testing.T) {
			_, items := newItem("normal", expiredDay, quality)
			UpdateQualityForAll(items)
			assert.Equal(t, quality-2, items[0].quality)
			assert.Equal(t, expiredDay-1, items[0].sellIn)
		})
	}
}

func Test_Aged_Brie_quality_should_increase_quality_by_1_per_day_until_expiration(t *testing.T) {
	_, items := newItem("Aged Brie", 1, 1)
	UpdateQualityForAll(items)
	assert.Equal(t, 2, items[0].quality)
	assert.Equal(t, 0, items[0].sellIn)
}

func Test_Aged_Brie_quality_should_increase_by_2_per_day_once_expiration(t *testing.T) {
	_, items := newItem("Aged Brie", -1, 1)
	UpdateQualityForAll(items)
	assert.Equal(t, 3, items[0].quality)
	assert.Equal(t, -2, items[0].sellIn)
}

func Test_Aged_Brie_quality_should_not_more_than_50(t *testing.T) {
	tests := map[string]struct {
		sellIn        int
		quality       int
		expectSellIn  int
		expectquality int
	}{
		"quality_49_should_not_more_than_50":        {10, 49, 9, 50},
		"quality_50_should_not_more_than_50":        {10, 50, 9, 50},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			_, items := newItem("Aged Brie", testcase.sellIn, testcase.quality)
			UpdateQualityForAll(items)
			assert.Equal(t, testcase.expectquality, items[0].quality)
			assert.Equal(t, testcase.expectSellIn, items[0].sellIn)
		})
	}
}

func Test_Sulfuras_should_not_expired_and_quality_is_80_forever(t *testing.T) {
	tests := map[string]struct {
		sellIn        int
		quality       int
		expectSellIn  int
		expectquality int
	}{
		"quality_should_not_change_for Sulfuras":     {20, 80, 20, 80},
		"sellIn_should_not_be_negative_for Sulfuras": {0, 80, 0, 80},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			_, items := newItem("Sulfuras, Hand of Ragnaros", testcase.sellIn, testcase.quality)
			UpdateQualityForAll(items)
			assert.Equal(t, testcase.expectquality, items[0].quality)
			assert.Equal(t, testcase.expectSellIn, items[0].sellIn)
		})
	}
}

func Test_Backstage_passes_quality_should_increase_according_to_sellIn_for(t *testing.T) {
	tests := map[string]struct {
		sellIn        int
		quality       int
		expectquality int
	}{
		"when sellin is more than 10 days，quality should not change":    {12, 50, 50},
		"when sellin is less than 10 days,quality should decrease by 2": {10, 40, 42},
		"when sellin is less than 5 days, quality should decrease by 3": {5, 30, 33},
		"when sellin is less than 0 day, quality should be 0":           {0, 30, 0},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			_, items := newItem("Backstage passes to a TAFKAL80ETC concert", testcase.sellIn, testcase.quality)
			UpdateQualityForAll(items)
			assert.Equal(t, testcase.expectquality, items[0].quality)
		})
	}
}

func newItem(name string, sellIn int, quality int) (item *Item, items []Item) {
	item = &Item{name, sellIn, quality}
	items = []Item{*item}
	return item, items
}
