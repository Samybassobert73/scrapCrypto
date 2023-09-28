package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
	"strconv"
	"time"
)

func main() {

	baseURL := "https://fr.finance.yahoo.com/crypto-monnaies/?guccounter=1&guce_referrer=aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8&guce_referrer_sig=AQAAALb-o5VX8La4lBa52ZOUSMXD_3h0OAy0_15dQYnTjVSpsjv189bNzkPPm2Sf4XIcd1ck4FvTbuB3LbLP8OdjHDtkXa8g9jYIpH_hlKFfrXPCfKmZGLkcHwgtYYOUy8vLUpFDKI0dGxUO__Va5aqOpU3vH2UsRghnzLtmKj6K3L96"

    offset := 0

    increment := 100

	c := colly.NewCollector()

	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {

		name := e.ChildText("td:nth-child(2)")
        price := e.ChildText("td:nth-child(3)")

        fmt.Printf("Name: %s\n", name)
        fmt.Printf("Price: %s\n", price)

    })

    for {
		if offset == 0 {
            url := baseURL
            // Print the URL for debugging
            fmt.Println("*****************************")
            fmt.Printf("offset: %s\n", offset)
            fmt.Println("*****************************")
            // Start scraping by visiting the URL
            err := c.Visit(url)
            if err != nil {
                fmt.Println("Error scraping URL:", err)
            }
        } else {
            // Scrape with the current offset
            url := baseURL + "&offset=" + strconv.Itoa(offset)
            // Print the URL for debugging
            fmt.Println("*****************************")
            fmt.Printf("offset: %s\n", offset)
            fmt.Println("*****************************")
            // Start scraping by visiting the URL
            err := c.Visit(url)
            if err != nil {
                fmt.Println("Error scraping URL:", err)
            }
        }

        offset += increment

        time.Sleep(5 * time.Second)
    }
}
