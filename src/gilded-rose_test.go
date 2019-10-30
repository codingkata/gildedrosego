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
	assert.True(t,true)
}
