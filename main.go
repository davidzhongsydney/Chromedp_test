package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var gStockFullName string
var gPriceStr string

func main() {

	gStockFullName = "uninitialized"
	gPriceStr = "uninitialized"

	dir, _ := ioutil.TempDir("", "chromedp-example")

	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:]) // chromedp.Flag("headless", false),

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	//var example string

	var res2 string
	var res3 string
	code := "bse"
	urlBase := `https://www.marketindex.com.au/asx/`

	fmt.Println(urlBase + code)

	err := chromedp.Run(ctx,
		chromedp.Navigate(urlBase+code),
		chromedp.WaitVisible(`#quoteapi-default-announcement`),
		chromedp.Sleep(5*time.Second),
		chromedp.Text(`#quote-summary-root`, &res2, chromedp.NodeVisible, chromedp.ByQueryAll),
		chromedp.Text(`.mi-table tbody`, &res3, chromedp.NodeVisible, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Println(res2)
	var stockFullName string
	temp := strings.Split(res2, "\n")

	fullNameRawStr := temp[0]

	fullNameEndIndex := strings.IndexByte(fullNameRawStr, '(') - 1
	stockFullName = fullNameRawStr[0:fullNameEndIndex]

	fmt.Println(stockFullName)

	gStockFullName = stockFullName

	stockDetailArr := strings.Split(res3, "\n")

	// last price
	lastPriceRaw := stockDetailArr[0]
	priceStartIndex := strings.IndexByte(lastPriceRaw, '$') + 1
	priceStr := lastPriceRaw[priceStartIndex:]

	price, _ := strconv.ParseFloat(priceStr, 64)
	fmt.Println(price)

	gPriceStr = priceStr
	//-------------------
	changeRaw := stockDetailArr[1]
	changePerStartIndex := strings.IndexByte(changeRaw, '(') + 1
	changePerEndIndex := strings.IndexByte(changeRaw, '%')
	changePerStr := changeRaw[changePerStartIndex:changePerEndIndex]
	fmt.Println(changePerStr)

	changeEndIndex := changePerStartIndex - 2
	changeStartIndex := strings.LastIndex(changeRaw[0:changeEndIndex], " ") + 1
	changeStr := changeRaw[changeStartIndex:changeEndIndex]
	fmt.Println(changeStr)

	// high low
	// "Day Range $1.52 - $1.60"
	dayRangeRaw := stockDetailArr[7]
	lowStartIndex := 11
	lowEndIndex := strings.IndexByte(dayRangeRaw, '-') - 1
	lowStr := dayRangeRaw[lowStartIndex:lowEndIndex]
	fmt.Println(lowStr)

	highStartIndex := lowEndIndex + 4
	highStr := dayRangeRaw[highStartIndex:]
	fmt.Println(highStr)

	// Volume
	// "Volume    20,278"
	volumeRaw := strings.ReplaceAll(stockDetailArr[3], ",", "")
	volumeStartIndex := strings.LastIndex(volumeRaw, " ") + 1
	volumeStr := volumeRaw[volumeStartIndex:]
	fmt.Println(volumeStr)

	//----------------------------------

	http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, gStockFullName, gPriceStr)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
