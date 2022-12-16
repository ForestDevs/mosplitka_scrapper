package page

import (
	"mosplitka-parser/models"
	"mosplitka-parser/utils"

	"github.com/gocolly/colly/v2"
)

const (
	productMainBlock           = "//div[@class='container']"                                                            // main block
	productTitle               = "//div[@class='tile-hero']//h1"                                                        // name
	productPrice               = "//div[@class='tile-hero']//p[contains(@class,'tile-shop__price')]"                    // main price
	productImages              = "//img[@class='pop-images-big-item__img']"                                             // images
	productFeaturesTitle       = "//div[@class='communication-prop__col'][1]//span[@class='tile-prop-tabs__name-wrap']" // product feateruse titiles from product card
	productFeaturesDescription = "//div[@class='communication-prop__col'][1]//p[@class='tile-prop-tabs__value']"        // valuse feauters
)

func productFeaturesCollector(x *colly.XMLElement, col map[string]string) map[string]string {
	titles := x.ChildTexts(productFeaturesTitle)
	descs := x.ChildTexts(productFeaturesDescription)
	for i := 0; i < len(titles); i++ {
		col[titles[i]] = " " + descs[i] + " "
	}
	return col
}

func imagesCollector(x *colly.XMLElement) []string {
	images := make([]string, 0)
	for _, src := range x.ChildAttrs(productImages, "src") {
		images = append(images, utils.Domain+src)
	}
	return images
}

func Product(c *colly.Collector, url string) models.Product {
	utils.OnRequest(c)
	var product models.Product
	c.OnXML(productMainBlock, func(x *colly.XMLElement) {
		var productFeatures map[string]string = make(map[string]string)
		name := x.ChildText(productTitle)
		price := x.ChildText(productPrice)
		images := imagesCollector(x)
		productFeatures = productFeaturesCollector(x, productFeatures)
		product = models.NewProduct(name, price, images, productFeatures)
	})
	c.Visit(url)
	return product
}
