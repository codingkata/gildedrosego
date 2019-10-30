This Kata was originally created by Terry Hughes (http://twitter.com/#!/TerryHughes). It is already on GitHub [here](https://github.com/NotMyself/GildedRose). See also [Bobby Johnson's description of the kata](http://iamnotmyself.com/2011/02/13/refactor-this-the-gilded-rose-kata/).

======================================
Gilded Rose 需求描述
======================================


欢迎来到镶金玫瑰(Gilded Rose)团队。
如你所知，我们是主城中的一个小旅店，店主非常友好，名叫Allison。我们也售卖最好的物品。但有些物品的品质会随着销售期限的临近而不断下降。

我们有一个软件系统来更新仓库中的库存物品信息。系统是由一个无名之辈Leeroy所开发的，他已经不在这了。

你的任务是添加新功能，这样我们就可以售卖新的物品。

先介绍一下我们的软件系统：

    - 每种物品都包含三种信息，如下所示：
        - 名称`name`，是物品名称。
        - 剩余天数`SellIn`，表示我们要在多少天之内把物品卖出去，即销售期。
        - 品质`quality`，表示物品的品相。
	
	- 每天结束时，系统会自动降低每种物品的`SellIn`和`quality`这两个值

注意：还有一些信息是你必须知道的：

    - 每种物品的品质（quality）不会超过50，也永远不会为负。
	- 每种物品每在仓库放一天，品质`Quality`会下降1；
	- 每种物品一旦超过销售期，品质`Quality`会以双倍速度加速下降
	
但是，以下为一些特殊物品，需要特殊处理：
   
    - 名为"Sulfuras"的物品是一种吉祥物，因此
	    - 它永不到期
	    - 品质`Quality`永远是80，而且永远不变。
	- "Aged Brie"的品质`Quality`会随着时间推移而提高，每天提高1
	- "Backstage passes"与Aged Brie类似，其品质`Quality`也会随着时间推移而提高；它自己的特点在于：
	    - 当还剩10天或更少的时候，品质`quality`每天提高2；
	    - 当还剩5天或更少的时候，品质`quality`每天提高3；
	    - 然而，一旦过期，品质就会降为0
     
我们最近签约了一个特殊物品供应商。这就要对我们的系统进行升级：

	- "Conjured"物品的品质`Quality`下降速度比正常物品快一倍

请随意对UpdateQuality函数进行修改和添加新代码，只要系统还能正常工作。

然而，不要修改Item类或其属性，因为那属于角落里的地精，他会非常愤怒地爆你头，因为他不相信代码共享所有制（如果你愿意，你可以将UpdateQuality方面和Items属性改为静态的，我们会掩护你的）。