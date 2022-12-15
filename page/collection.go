package page

import (
	"fmt"
	"mosplitka-parser/utils"

	"github.com/gocolly/colly/v2"
)

func Collection(c *colly.Collector, url string) {
	utils.OnRequest(c)

	c.OnXML("//div[@class='m-collection-plitka--detail']", func(x *colly.XMLElement) {
		fmt.Printf("x.ChildText(\"//h1\"): %v\n", x.ChildText("//h1"))
	})

	c.Visit(url)
}
