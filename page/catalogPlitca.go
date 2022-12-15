package page

import (
	"mosplitka-parser/utils"
	"strconv"

	"github.com/gocolly/colly/v2"
)

const url = utils.Domain + "/catalog/plitka/"

func scrapCollections(c *colly.Collector, url string) {
	utils.OnRequest(c)

	// find collections card on page
	c.OnXML("//div[@class='card__name']/a", func(x *colly.XMLElement) {
		cInstance := utils.NewCollector()
		Collection(cInstance, utils.Domain+x.Attr("href"))
	})

	c.Visit(url)
}

func pagination(c *colly.Collector) {
	c.OnXML("//ul[@class='pagination-catalog__items']/li[last()]", func(x *colly.XMLElement) {
		lastPage, _ := strconv.Atoi(x.Text)
		for i := 1; i <= lastPage; i++ {
			cInstance := utils.NewCollector()
			scrapCollections(cInstance, url+"?PAGEN_1="+strconv.Itoa(i))
		}
	})
}

func CatalogPlitca(c *colly.Collector) {
	utils.OnRequest(c)
	pagination(c)
	c.Visit(url)
}
