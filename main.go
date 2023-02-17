package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

var urls=  []string{
	"https://www.amazon.com/YSSOA-Gaming-Office-Adjustable-Chair/dp/B08XQNSH7B/ref=sr_1_1?keywords=gaming+chairs&pf_rd_i=23508887011&pf_rd_m=ATVPDKIKX0DER&pf_rd_p=434db2ed-6d53-4c59-b173-e8cd550a2e4f&pf_rd_r=MV92YE7BM6SAFXRR30A2&pf_rd_s=merchandised-search-5&pf_rd_t=101&qid=1676578588&sr=8-1",
	
}

type Product struct {
	Name  string
	Price string
}

func main() {
	time_ := time.Now()
	var products []Product

	for _, url := range urls {
		product := scrape(url)
		products = append(products, product)
	}

	file, err :=  os.Create("data.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close() 

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Price"})

	for _, product := range products {
		writer.Write([]string{
			product.Name,
			product.Price,
		})
	}

	elapsed := time.Since(time_).Seconds()
	fmt.Printf("Time %.2f\n", elapsed)
}

func scrape(url string) Product {
	var product Product
	c := colly.NewCollector()

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("I visiting ->", request.URL)
	})

	c.OnHTML("#dp-container", func(e *colly.HTMLElement) {
		Name := e.ChildText("#centerCol #productTitle")
		Price := e.ChildText("#centerCol #color_name_1_price") 

		product = Product{Name, Price}
	})
	
	c.Visit(url)
	
	return product
}
