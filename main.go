package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
)

var (
	url        = "https://e.intant.ru"
	path       = "//body/div[3]/div[2]/div[1]/div[1]"
	secondPath = "//html/body/div[3]/div/div[2]/div/div[3]/div[2]/div[1]"
)

func test(q string, path string) {
	doc, err := htmlquery.LoadURL(q)

	if err != nil {
		panic(err)
	}

	w, err := htmlquery.QueryAll(doc, path)
	if err != nil {
		fmt.Println(err)
	}

	str := "//a[@href]"

	for _, s := range w {
		list, err := htmlquery.QueryAll(s, str)
		if err != nil {
			panic(err)
		}
		for _, l := range list {

			f := strings.TrimSpace(htmlquery.SelectAttr(l, "href"))
			fmt.Println(f)
			if err != nil {
				panic(err)
			}

			test(url+f, secondPath)

			if err != nil {
				panic(err)
			}

		}

	}

}

func main() {
	test("https://e.intant.ru/catalog/hardware", path)

}
