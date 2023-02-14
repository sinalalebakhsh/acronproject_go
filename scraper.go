package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type item struct { 
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)
	fmt.Println("", "Name", "\t\t", "Price", "\t\t", "URLs-Link")
	fmt.Println("", "----", "\t\t", "-----", "\t\t", "------------------------------------")

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item {
			Name: h.ChildText("h2.product-title"),
			Price: h.ChildText("div.sale-price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		fmt.Println("", item.Name, "\t", item.Price, "\t", item.ImgUrl)


	
	})

	c.Visit("http://j2store.net/demo/index.php/shop")
}

/************************  RESULT   ****************************************

go run .

 Name            Price           URLs-Link
 ----            -----           ------------------------------------
 Simple          $80.00          http://j2store.net/demo/images/themeparrot/t-shirt/t-shirt_01.png
 Variable        $40.00          http://j2store.net/demo/images/themeparrot/leggings/leggins_01.png
 Configurable    $50.00          http://j2store.net/demo/images/themeparrot/pots/pots_and_pans_01.png
 Downloadable    $30.00          http://j2store.net/demo/images/themeparrot/product_image_07.png
 Blender1        $110.00         http://j2store.net/demo/images/themeparrot/blenders/blenders_01.png
 Blender2        $150.00         http://j2store.net/demo/images/themeparrot/blenders/blenders_02.png
 Blender3        $79.00          http://j2store.net/demo/images/themeparrot/blenders/blenders_03.png
 Blender4        $110.00         http://j2store.net/demo/images/themeparrot/blenders/blenders_05.png
 T-Shirt1        $95.00          http://j2store.net/demo/images/themeparrot/t-shirt/t-shirt_05.png

 ****************************************************************************/ 