package main

import (
	"fmt"
	"mosplitka-parser/page"
	"mosplitka-parser/utils"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	fmt.Println("                           Start \n \n \n                    ")
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond) // Build our new spinner
	s.Start()
	c := utils.NewCollector()
	page.CatalogPlitca(c)
	s.Stop()
}
