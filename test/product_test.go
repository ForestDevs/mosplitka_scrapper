package page

import (
	"mosplitka-parser/page"
	"mosplitka-parser/utils"
	"testing"
)

func TestProduct(t *testing.T) {
	url := "https://mosplitka.ru/product/keramogranit_wonder_belyy_rektifikat_44_8x89_8/"
	c := utils.NewCollector()
	t.Run("product-card", func(t *testing.T) {
		page.Product(c, url)
	})
}
