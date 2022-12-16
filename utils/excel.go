package utils

import (
	"fmt"
	"log"
	"mosplitka-parser/models"
	"os"
	"strconv"
	"strings"

	"github.com/takuoki/clmconv"

	"github.com/xuri/excelize/v2"
)

func productsSheet(f *excelize.File, products []models.Product) {
	var titles []string = []string{"Название", "Цена", "Картинки"}
	var features map[string]int = make(map[string]int)
	index := f.NewSheet("Товары")
	f.SetActiveSheet(index)
	productIndex := 2
	lastIndex := 0

	for i, title := range titles {
		f.SetCellValue("Товары", clmconv.Itoa(i)+"1", title)
		lastIndex = i
	}

	lastIndex++

	for _, proproduct := range products {
		f.SetCellValue("Товары", clmconv.Itoa(0)+strconv.Itoa(productIndex), proproduct.Name)
		f.SetCellValue("Товары", clmconv.Itoa(1)+strconv.Itoa(productIndex), proproduct.Price)
		f.SetCellValue("Товары", clmconv.Itoa(2)+strconv.Itoa(productIndex), strings.Join(proproduct.Images, ";"))
		for k, v := range proproduct.Features {
			_, ok := features[k]
			if !ok {
				f.SetCellValue("Товары", clmconv.Itoa(lastIndex)+"1", k)
				features[k] = lastIndex
				lastIndex++
			}
			f.SetCellValue("Товары", clmconv.Itoa(features[k])+strconv.Itoa(productIndex), v)
		}
		productIndex++
	}
}

func collectionSheet(f *excelize.File, collection models.Collection) {
	var titles []string = []string{"Название", "Цена", "Картинки"}
	index := f.NewSheet("Коллекция")
	f.SetActiveSheet(index)
	lastIndex := 0

	for i, title := range titles {
		f.SetCellValue("Коллекция", clmconv.Itoa(i)+"1", title)

		if title == "Название" {
			f.SetCellValue("Коллекция", clmconv.Itoa(i)+"2", collection.Name)
		}

		if title == "Цена" {
			f.SetCellValue("Коллекция", clmconv.Itoa(i)+"2", collection.Price)
		}

		if title == "Картинки" {
			f.SetCellValue("Коллекция", clmconv.Itoa(i)+"2", collection.Image)
		}

		lastIndex = i
	}

	lastIndex++

	for k, v := range collection.Features {
		f.SetCellValue("Коллекция", clmconv.Itoa(lastIndex)+"1", k)
		f.SetCellValue("Коллекция", clmconv.Itoa(lastIndex)+"2", v)
		lastIndex++
	}

	f.SetActiveSheet(index)
}

func ExcelWrite(collection models.Collection) {
	f := excelize.NewFile()
	collectionSheet(f, collection)
	productsSheet(f, collection.Products)
	if _, err := os.Stat("./data/"); os.IsNotExist(err) {
		if err := os.Mkdir("data", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.SaveAs("./data/" + collection.Name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
