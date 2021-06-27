package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

// func main() {

// 	dir, _ := ioutil.TempDir("", "chromedp-example")

// 	defer os.RemoveAll(dir)

// 	opts := append(chromedp.DefaultExecAllocatorOptions[:]) // chromedp.Flag("headless", false),
// 	//	chromedp.DisableGPU,
// 	//	chromedp.UserDataDir(dir),

// 	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
// 	defer cancel()

// 	// create chrome instance
// 	ctx, cancel := chromedp.NewContext(
// 		allocCtx,
// 		//chromedp.WithDebugf(log.Printf),
// 	)
// 	defer cancel()

// 	// create a timeout
// 	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
// 	defer cancel()

// 	// navigate to a page, wait for an element, click
// 	//var example string

// 	var res2 string

// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate(`https://www.marketindex.com.au/all-ordinaries`),
// 		// wait for footer element is visible (ie, page is loaded)
// 		chromedp.WaitVisible(`#quoteapi-default-announcement`),
// 		// find and click "Example" link
// 		// chromedp.Click(`img`, chromedp.NodeVisible),

// 		// chromedp.WaitVisible(`#quoteapi-default-announcement`),
// 		//chromedp.Sleep(10*time.Second),

// 		//chromedp.Click(`#sticky-table-root > p.text-center.mt-1 > button`, chromedp.NodeVisible, chromedp.ByID),
// 		//chromedp.Sleep(5*time.Second),
// 		//chromedp.Click(`#sticky-table-root > p.text-center.mt-1 > button`, chromedp.NodeVisible, chromedp.ByID),
// 		//chromedp.Click(`button[@data-quoteapi='pageSize=0']`, chromedp.NodeVisible),
// 		//chromedp.Click(`Show All Companies`, chromedp.NodeVisible, chromedp.BySearch),
// 		//chromedp.Sleep(10*time.Second),
// 		chromedp.Sleep(10*time.Second),
// 		//chromedp.OuterHTML("#sticky-table-root", &html),
// 		//chromedp.WaitVisible(`#quoteapi-default-announcement`),
// 		// retrieve the text of the textarea
// 		//chromedp.Value(`#sticky-table-root tr td a`, &example),
// 		// chromedp.Text(`#sticky-table-root tbody tr`, &res, chromedp.NodeVisible, chromedp.ByID),
// 		// chromedp.Text(`#sticky-table-root > p.text-center.mt-1 > button`, &res2, chromedp.NodeVisible, chromedp.ByID),

// 		chromedp.Text(`#sticky-table-root tbody`, &res2, chromedp.NodeVisible, chromedp.ByID),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//	fmt.Println(res2)

// 	temp := strings.Split(res2, "\n")

// 	for _, stockStr := range temp {
// 		var price float32
// 		var change float32
// 		var changePercent float32
// 		var high float32
// 		var low float32
// 		var volume int64

// 		code := stockStr[1:4]

// 		stockNoCommaStr := strings.ReplaceAll(stockStr, ",", "")
// 		priceIndexStart := strings.IndexByte(stockNoCommaStr[0:], '$')
// 		newline := stockNoCommaStr[priceIndexStart:]

// 		// AT1	Atomo Diagnostics Ltd  	$0.255	-0.02	-7.27%	$0.28	$0.25	2,864,904	$104,012,036	-34.62%
// 		_, err := fmt.Sscanf(newline, "$%f	%f	%f%%	$%f	$%f	%d", &price, &change, &changePercent, &high, &low, &volume)

// 		if err != nil {
// 			_, err := fmt.Sscanf(newline, "$%f	%f	%f	$%f	$%f	%d", &price, &change, &changePercent, &high, &low, &volume)

// 			if err != nil {
// 				continue
// 			}
// 		}
// 		fmt.Println(code, price, change, changePercent, high, low, volume)
// 	}

// 	// d1 := []byte(res2)
// 	// err = ioutil.WriteFile("./dat1.txt", d1, 0644)
// 	// check(err)
// }

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {

	dir, _ := ioutil.TempDir("", "chromedp-example")

	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:]) // chromedp.Flag("headless", false),
	//	chromedp.DisableGPU,
	//	chromedp.UserDataDir(dir),

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
		//chromedp.Navigate(`https://www.investsmart.com.au/shares/indices/`),
		chromedp.Navigate(urlBase+code),

		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`#quoteapi-default-announcement`),
		// find and click "Example" link
		// chromedp.Click(`img`, chromedp.NodeVisible),

		// chromedp.WaitVisible(`#intercom-frame`),
		chromedp.Sleep(5*time.Second),

		//chromedp.Click(`#sticky-table-root > p.text-center.mt-1 > button`, chromedp.NodeVisible, chromedp.ByID),
		//chromedp.Sleep(5*time.Second),
		//chromedp.Click(`#sticky-table-root > p.text-center.mt-1 > button`, chromedp.NodeVisible, chromedp.ByID),
		//chromedp.Click(`button[@data-quoteapi='pageSize=0']`, chromedp.NodeVisible),
		//chromedp.Click(`Show All Companies`, chromedp.NodeVisible, chromedp.BySearch),
		//chromedp.Sleep(10*time.Second),
		chromedp.Sleep(5*time.Second),
		//chromedp.OuterHTML("#sticky-table-root", &html),
		//chromedp.WaitVisible(`#quoteapi-default-announcement`),
		// retrieve the text of the textarea
		//chromedp.Value(`#sticky-table-root tr td a`, &example),
		// chromedp.Text(`#sticky-table-root tbody tr`, &res, chromedp.NodeVisible, chromedp.ByID),
		// chromedp.Text(`#sticky-table-root > p.text-center.mt-1 > button`, &res2, chromedp.NodeVisible, chromedp.ByID),

		//chromedp.Text(`#main tbody`, &res2, chromedp.NodeVisible, chromedp.ByID),
		chromedp.Text(`#quote-summary-root`, &res2, chromedp.NodeVisible, chromedp.ByQueryAll),
		chromedp.Text(`.mi-table tbody`, &res3, chromedp.NodeVisible, chromedp.ByQueryAll),
		//chromedp.Text(`#mob-sidebar-elements-root`, &res2, chromedp.NodeVisible, chromedp.ByID),
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

	stockDetailArr := strings.Split(res3, "\n")

	// last price
	lastPriceRaw := stockDetailArr[0]
	priceStartIndex := strings.IndexByte(lastPriceRaw, '$') + 1
	priceStr := lastPriceRaw[priceStartIndex:]

	price, _ := strconv.ParseFloat(priceStr, 64)
	fmt.Println(price)
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
	// //------ ORDS -----------------------------
	// ordsRaw_1 := temp[0]

	// var indexValueFloat float32
	// var changePercentFloat float32
	// var changeAmountFloat float32

	// //"All Ordinaries	7312.0	risen by 0.73% 53.2"
	// ordsRaw_2 := strings.ReplaceAll(ordsRaw_1, "All Ordinaries	", "")
	// ordsRaw_3 := strings.ReplaceAll(ordsRaw_2, "risen by ", "")
	// ordsRaw_4 := strings.ReplaceAll(ordsRaw_3, "fallen by ", "")

	// //"7312.0	0.73% 53.2"
	// _, err = fmt.Sscanf(ordsRaw_4, "%f	%f%% %f", &indexValueFloat, &changePercentFloat, &changeAmountFloat)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //--------- ASX200 --------------------------------------
	// //"S&P/ASX 200	7055.4	risen by 0.83% 58.0"
	// asx200Raw_1 := temp[5]

	// //5401.5	0.42% 22.7
	// asx200Raw_2 := strings.ReplaceAll(asx200Raw_1, "S&P/ASX 200	", "")
	// asx200Raw_3 := strings.ReplaceAll(asx200Raw_2, "risen by ", "")
	// asx200Raw_4 := strings.ReplaceAll(asx200Raw_3, "fallen by ", "")

	// //"7312.0	0.73% 53.2"
	// _, err = fmt.Sscanf(asx200Raw_4, "%f	%f%% %f", &indexValueFloat, &changePercentFloat, &changeAmountFloat)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//---------All Tech --------------

	// 	// AT1	Atomo Diagnostics Ltd  	$0.255	-0.02	-7.27%	$0.28	$0.25	2,864,904	$104,012,036	-34.62%
	// 	_, err := fmt.Sscanf(newline, "$%f	%f	%f%%	$%f	$%f	%d", &price, &change, &changePercent, &high, &low, &volume)

	// d1 := []byte(res2)
	// err = ioutil.WriteFile("./dat1.txt", d1, 0644)
	// check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func main() {

// 	dir, _ := ioutil.TempDir("", "chromedp-example")

// 	defer os.RemoveAll(dir)

// 	opts := append(chromedp.DefaultExecAllocatorOptions[:]) // chromedp.Flag("headless", false),
// 	//	chromedp.DisableGPU,
// 	//	chromedp.UserDataDir(dir),

// 	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
// 	defer cancel()

// 	// create chrome instance
// 	ctx, cancel := chromedp.NewContext(
// 		allocCtx,
// 		//chromedp.WithDebugf(log.Printf),
// 	)
// 	defer cancel()

// 	// create a timeout
// 	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
// 	defer cancel()

// 	// navigate to a page, wait for an element, click
// 	//var example string

// 	var res2 string

// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate(`https://www.marketindex.com.au`),
// 		// wait for footer element is visible (ie, page is loaded)
// 		chromedp.WaitVisible(`#quoteapi-default-announcement`),

// 		chromedp.Sleep(5*time.Second),
// 		chromedp.Text(`#mob-sidebar-elements-root`, &res2, chromedp.NodeVisible, chromedp.ByID),
// 		//chromedp.Text(`#mob-sidebar-elements-root`, &res2, chromedp.NodeVisible, chromedp.ByID),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//	fmt.Println(res2)

// 	temp := strings.Split(res2, "\n")

// 	//------ ORDS -----------------------------
// 	tradeStatus := temp[0]
// 	nextOpenDate := temp[1]
// 	action := temp[3]

// 	fmt.Println(tradeStatus, nextOpenDate, action)

// 	// 	// AT1	Atomo Diagnostics Ltd  	$0.255	-0.02	-7.27%	$0.28	$0.25	2,864,904	$104,012,036	-34.62%
// 	// 	_, err := fmt.Sscanf(newline, "$%f	%f	%f%%	$%f	$%f	%d", &price, &change, &changePercent, &high, &low, &volume)

// 	// d1 := []byte(res2)
// 	// err = ioutil.WriteFile("./dat1.txt", d1, 0644)
// 	// check(err)
// }
