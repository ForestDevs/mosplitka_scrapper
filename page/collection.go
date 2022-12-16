package page

import (
	"mosplitka-parser/utils"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func trimPriceString(str string) string {
	r := regexp.MustCompile("\n\\s+")
	return strings.ReplaceAll(r.ReplaceAllString(str, ""), "от", "")
}

func Collection(c *colly.Collector, url string) {
	utils.OnRequest(c)

	c.OnXML("//div[@class='m-collection-plitka--detail']", func(x *colly.XMLElement) {
		// title := x.ChildText("//h1")
		//price:= trimPriceString(x.ChildText("//div[@class='rate'][text()]"))
		// image := utils.Domain + x.ChildAttr("//div[@class='m-collection-plitka--detail-top']//div[@class='swiper-slide']//a", "href")
	})

	c.Visit(url)
}
