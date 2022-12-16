package page

import (
	"mosplitka-parser/models"
	"mosplitka-parser/utils"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	productMainBlock           = "//div[@class='container']"                                                     // main block
	productTitle               = "//div[@class='tile-hero']//h1"                                                 // name
	productPrice               = "//div[@class='tile-hero']//p[contains(@class,'tile-shop__price')]"             // main price
	productImages              = "//img[@class='pop-images-big-item__img']"                                      // images
	productFeaturesItems       = "//div[@class='communication-prop__col'][1]//li"                                // product feateruse titiles from product card
	productFeaturesDescription = "//div[@class='communication-prop__col'][1]//p[@class='tile-prop-tabs__value']" // valuse feauters
)

func productFeaturesCollector(x *colly.XMLElement, col map[string]string) map[string]string {
	elementsCount := len(x.ChildTexts(productFeaturesItems))
	if elementsCount != 0 {
		for i := 0; i <= elementsCount; i++ {
			title := x.ChildText("//div[@class='communication-prop__col'][1]//li[" + strconv.Itoa(i) + "]//span[@class='tile-prop-tabs__name-wrap']")
			manyDescCount := len(x.ChildTexts("//div[@class='communication-prop__col'][1]//li[" + strconv.Itoa(i) + "]//span[@class='tile-prop-tabs__value-name']//span[@class='tile-prop-tabs__row']"))
			if manyDescCount > 1 {
				desc := strings.Join(x.ChildTexts("//div[@class='communication-prop__col'][1]//li["+strconv.Itoa(i)+"]//span[@class='tile-prop-tabs__value-name']//span[@class='tile-prop-tabs__row']"), ",")
				col[title] = desc
			} else {
				desc := x.ChildText("//div[@class='communication-prop__col'][1]//li[" + strconv.Itoa(i) + "]//span[@class='tile-prop-tabs__value-name']")
				if desc == "" {
					desc = x.ChildText("//div[@class='communication-prop__col'][1]//li[" + strconv.Itoa(i) + "]//a[@class='tile-prop-tabs__value-name']")
				}
				col[title] = desc
			}
		}
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
