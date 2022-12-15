package main

import (
	"mosplitka-parser/page"
	"mosplitka-parser/utils"
)

func main() {
	c := utils.NewCollector()
	page.CatalogPlitca(c)
}
