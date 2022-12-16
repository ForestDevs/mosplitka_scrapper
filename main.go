package main

import (
	"mosplitka-parser/page"
	"mosplitka-parser/utils"
	"time"

	"github.com/briandowns/spinner"
)

func main() {

	someSet := []string{"[                    ]", "[===>                  ]", "[=====>                ]", "[=======>              ]", "[========>             ]", "[==========>           ]", "[============>         ]", "[==============>       ]", "[================>     ]", "[==================>   ]", "[====================> ]", "[=====================>]"}
	s := spinner.New(someSet, 100*time.Millisecond) // Build our new spinner
	s.Start()
	c := utils.NewCollector()
	page.CatalogPlitca(c)
	s.Stop()
}
