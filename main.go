package main

import (
	"encoding/json"
	"io/ioutil"
	"mosplitka-parser/page"
	"mosplitka-parser/utils"
	"time"

	"github.com/briandowns/spinner"
)

type Config struct {
	Collections []string `json: "collections"`
}

func main() {

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	cfg := Config{}

	if err := json.Unmarshal(file, &cfg); err != nil {
		panic(err)
	}

	s := spinner.New(spinner.CharSets[80], 100*time.Millisecond) // Build our new spinner
	s.Start()
	c := utils.NewCollector()

	if len(cfg.Collections) != 0 {
		for _, url := range cfg.Collections {
			page.Collection(c, url)
		}
	} else {
		page.CatalogPlitca(c)
	}

	s.Stop()
}
