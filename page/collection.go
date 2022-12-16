package page

import (
	"mosplitka-parser/models"
	"mosplitka-parser/utils"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	collectionMainBlock           = "//div[@class='m-collection-plitka--detail']" // collection cards block
	collectionTitle               = "//h1"                                        // collection name
	collectionPrice               = "//div[@class='rate'][text()]"
	collectionImage               = "//div[@class='m-collection-plitka--detail-top']//div[@class='swiper-slide']//a"                                                       // first collection image
	collectionFeaturesTitle       = "//div[contains(@class,'size-use-item')]//div[@class='title']"                                                                         // titiles from features block
	collectionFeaturesDescription = "//div[contains(@class,'size-use-item')]//div[@class='desc']"                                                                          // descriptions from featrues block
	collectionProductsCard        = "//div[contains(@class,'products product-list-block plitka_new')]//div[contains(@class,'product--wrap')]//a[contains(@class,'title')]" // products cards hrefs
)

func productsCollector(x *colly.XMLElement) []models.Product {
	var products []models.Product
	for _, cardHref := range x.ChildAttrs(collectionProductsCard, "href") {
		cInstance := utils.NewCollector()
		products = append(products, Product(cInstance, utils.Domain+cardHref))
	}
	return products
}

func collectionFeaturesCollector(x *colly.XMLElement, col map[string]string) map[string]string {
	titles := x.ChildTexts(collectionFeaturesTitle)
	descs := x.ChildTexts(collectionFeaturesDescription)
	for i := 0; i < len(titles); i++ {
		col[titles[i]] = strings.ReplaceAll(" "+descs[i], "\u00a0", " ")
	}
	return col
}

func trimPriceString(str string) string {
	r := regexp.MustCompile("\n\\s+")
	return strings.ReplaceAll(r.ReplaceAllString(str, ""), "от", "")
}

func Collection(c *colly.Collector, url string) {
	utils.OnRequest(c)

	c.OnXML(collectionMainBlock, func(x *colly.XMLElement) {
		var collectionFeatures map[string]string = make(map[string]string)
		name := x.ChildText(collectionTitle)
		price := trimPriceString(x.ChildText(collectionPrice))
		image := utils.Domain + x.ChildAttr(collectionImage, "href")
		collectionFeatures = collectionFeaturesCollector(x, collectionFeatures)
		products := productsCollector(x)
		collection := models.NewCollection(name, price, image, collectionFeatures, products)
		utils.ExcelWrite(collection)
	})

	c.Visit(url)
}
