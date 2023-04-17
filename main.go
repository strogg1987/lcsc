package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

type Response struct {
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Result struct {
		BrandHomeVO           interface{} `json:"brandHomeVO"`
		CatalogVOS            interface{} `json:"catalogVOS"`
		ProductSearchResultVO interface{} `json:"productSearchResultVO"`
		LcBookProductResultVO interface{} `json:"lcBookProductResultVO"`
		LcOrderLastPage       int         `json:"lcOrderLastPage"`
		ProductTotalPage      interface{} `json:"productTotalPage"`
		IsToDetail            bool        `json:"isToDetail"`
		IsToBrand             interface{} `json:"isToBrand"`
		TipProductDetailUrlVO struct {
			ProductCode  string `json:"productCode"`
			ProductModel string `json:"productModel"`
			BrandNameEn  string `json:"brandNameEn"`
			CatalogName  string `json:"catalogName"`
		} `json:"tipProductDetailUrlVO"`
		BrandUrlMap interface{} `json:"brandUrlMap"`
	} `json:"result"`
}

type BigSummary struct {
	Code   int `json:"code"`
	Msg    any `json:"msg"`
	Result struct {
		BrandHomeVO any `json:"brandHomeVO"`
		CatalogVOS  []struct {
			CatalogID     int    `json:"catalogId"`
			ParentID      any    `json:"parentId"`
			ParentName    any    `json:"parentName"`
			CatalogName   any    `json:"catalogName"`
			CatalogNameEn string `json:"catalogNameEn"`
			ChildCatelogs []struct {
				CatalogID     int    `json:"catalogId"`
				ParentID      any    `json:"parentId"`
				ParentName    any    `json:"parentName"`
				CatalogName   any    `json:"catalogName"`
				CatalogNameEn string `json:"catalogNameEn"`
				ChildCatelogs any    `json:"childCatelogs"`
				ProductNum    int    `json:"productNum"`
			} `json:"childCatelogs"`
			ProductNum int `json:"productNum"`
		} `json:"catalogVOS"`
		ProductSearchResultVO struct {
			TotalCount  int `json:"totalCount"`
			CurrentPage int `json:"currentPage"`
			PageSize    int `json:"pageSize"`
			ProductList []struct {
				ProductID          int     `json:"productId"`
				ProductCode        string  `json:"productCode"`
				ProductWeight      float64 `json:"productWeight"`
				ForeignWeight      float64 `json:"foreignWeight"`
				Weight             float64 `json:"weight"`
				IsOnsale           bool    `json:"isOnsale"`
				DollarLadderPrice  any     `json:"dollarLadderPrice"`
				IsForeignOnsale    bool    `json:"isForeignOnsale"`
				MinBuyNumber       int     `json:"minBuyNumber"`
				MaxBuyNumber       int     `json:"maxBuyNumber"`
				IsAutoOffsale      bool    `json:"isAutoOffsale"`
				MinPacketUnit      string  `json:"minPacketUnit"`
				ProductUnit        string  `json:"productUnit"`
				ProductArrange     string  `json:"productArrange"`
				MinPacketNumber    int     `json:"minPacketNumber"`
				EncapStandard      string  `json:"encapStandard"`
				ProductModel       string  `json:"productModel"`
				BrandID            int     `json:"brandId"`
				BrandNameEn        string  `json:"brandNameEn"`
				CatalogID          int     `json:"catalogId"`
				CatalogName        string  `json:"catalogName"`
				ParentCatalogID    int     `json:"parentCatalogId"`
				ParentCatalogName  string  `json:"parentCatalogName"`
				ProductDescEn      any     `json:"productDescEn"`
				ProductIntroEn     string  `json:"productIntroEn"`
				IsHasBattery       bool    `json:"isHasBattery"`
				IsDiscount         bool    `json:"isDiscount"`
				IsHot              bool    `json:"isHot"`
				IsEnvironment      bool    `json:"isEnvironment"`
				IsPreSale          bool    `json:"isPreSale"`
				ProductLadderPrice any     `json:"productLadderPrice"`
				LadderDiscountRate any     `json:"ladderDiscountRate"`
				ProductPriceList   []struct {
					Ladder          int     `json:"ladder"`
					ProductPrice    string  `json:"productPrice"`
					DiscountRate    string  `json:"discountRate"`
					CurrencyPrice   float64 `json:"currencyPrice"`
					UsdPrice        float64 `json:"usdPrice"`
					CurrencySymbol  string  `json:"currencySymbol"`
					CnyProductPrice string  `json:"cnyProductPrice"`
				} `json:"productPriceList"`
				StockJs         int `json:"stockJs"`
				StockSz         int `json:"stockSz"`
				StockHk         int `json:"stockHk"`
				WmStockHk       int `json:"wmStockHk"`
				DomesticStockVO struct {
					Total           int `json:"total"`
					ShipImmediately int `json:"shipImmediately"`
					Ship3Days       int `json:"ship3Days"`
				} `json:"domesticStockVO"`
				OverseasStockVO struct {
					Total           int `json:"total"`
					ShipImmediately int `json:"shipImmediately"`
					Ship3Days       int `json:"ship3Days"`
				} `json:"overseasStockVO"`
				ShipImmediately    int      `json:"shipImmediately"`
				Ship3Days          int      `json:"ship3Days"`
				SmtAloneNumberSz   any      `json:"smtAloneNumberSz"`
				SmtAloneNumberJs   any      `json:"smtAloneNumberJs"`
				StockNumber        int      `json:"stockNumber"`
				Split              int      `json:"split"`
				ProductImageURL    string   `json:"productImageUrl"`
				ProductImageURLBig string   `json:"productImageUrlBig"`
				PdfURL             string   `json:"pdfUrl"`
				ProductImages      []string `json:"productImages"`
				ParamVOList        []struct {
					ParamCode             string  `json:"paramCode"`
					ParamName             string  `json:"paramName"`
					ParamNameEn           string  `json:"paramNameEn"`
					ParamValue            string  `json:"paramValue"`
					ParamValueEn          string  `json:"paramValueEn"`
					ParamValueEnForSearch float64 `json:"paramValueEnForSearch"`
					IsMain                bool    `json:"isMain"`
					SortNumber            int     `json:"sortNumber"`
				} `json:"paramVOList"`
				IsReel                bool    `json:"isReel"`
				ReelPrice             float64 `json:"reelPrice"`
				ProductModelHighlight string  `json:"productModelHighlight"`
				ProductCodeHighlight  any     `json:"productCodeHighlight"`
				CatalogCode           string  `json:"catalogCode"`
				ParentCatalogCode     string  `json:"parentCatalogCode"`
				PdfLinkURL            any     `json:"pdfLinkUrl"`
				URL                   string  `json:"url"`
				ShareCostPrice        any     `json:"shareCostPrice"`
				ActivityID            any     `json:"activityId"`
				ActivityName          any     `json:"activityName"`
				ActivityStartTime     any     `json:"activityStartTime"`
				ActivityEndTime       any     `json:"activityEndTime"`
				StockPrice            any     `json:"stockPrice"`
			} `json:"productList"`
		} `json:"productSearchResultVO"`
		LcBookProductResultVO any `json:"lcBookProductResultVO"`
		LcOrderLastPage       int `json:"lcOrderLastPage"`
		ProductTotalPage      int `json:"productTotalPage"`
		IsToDetail            any `json:"isToDetail"`
		IsToBrand             any `json:"isToBrand"`
		TipProductDetailURLVO any `json:"tipProductDetailUrlVO"`
		BrandURLMap           any `json:"brandUrlMap"`
	} `json:"result"`
}

func main() {
	switch os.Args[1] {
	case "lcscpn":
		resp, err := http.Get("https://wmsc.lcsc.com/wmsc/search/global?keyword=" + os.Args[2])
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			return
		}
		defer resp.Body.Close()

		// Читаем тело ответа
		var part Response
		err = json.NewDecoder(resp.Body).Decode(&part)
		if err != nil {
			fmt.Println("Ошибка чтения тела ответа:", err)
			return
		}
		fmt.Printf("Manufacturer: %+v\n", part.Result.TipProductDetailUrlVO.BrandNameEn)
		fmt.Printf("Manufacturer PN: %+v\n", part.Result.TipProductDetailUrlVO.ProductModel)
		fmt.Printf("LCSC PN: %+v\n", part.Result.TipProductDetailUrlVO.ProductCode)
		resp.Body.Close()
		resp, err = http.Get("https://wmsc.lcsc.com/wmsc/search/global?keyword=" + part.Result.TipProductDetailUrlVO.ProductModel)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса:", err)
			return
		}
		defer resp.Body.Close()

		// Читаем тело ответа
		var summary BigSummary
		err = json.NewDecoder(resp.Body).Decode(&summary)
		if err != nil {
			fmt.Println("Ошибка чтения тела ответа:", err)
			return
		}
		if len(summary.Result.ProductSearchResultVO.ProductList) > 0 {
			fmt.Printf("Summary: %+v\n", StripHTMLTags(summary.Result.ProductSearchResultVO.ProductList[0].ProductIntroEn))

			fmt.Printf("Price max: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ProductPriceList[0].ProductPrice)
			fmt.Printf("Price mix: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ProductPriceList[len(summary.Result.ProductSearchResultVO.ProductList[0].ProductPriceList)-1].ProductPrice)

			fmt.Printf("Ship imm: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ShipImmediately)
		} else {
			fmt.Println("not in stock")
		}
	case "mpn":
		{

			resp, err := http.Get("https://wmsc.lcsc.com/wmsc/search/global?keyword=" + os.Args[2])
			if err != nil {
				fmt.Println("Ошибка выполнения запроса:", err)
				return
			}
			defer resp.Body.Close()

			// Читаем тело ответа
			var summary BigSummary
			err = json.NewDecoder(resp.Body).Decode(&summary)
			if err != nil {
				fmt.Println("Ошибка чтения тела ответа:", err)
				return
			}
			fmt.Printf("Manufacturer: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].BrandNameEn)
			fmt.Printf("Manufacturer PN: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ProductModel)
			fmt.Printf("LCSC PN: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ProductCode)
			fmt.Printf("Summary: %+v\n", StripHTMLTags(summary.Result.ProductSearchResultVO.ProductList[0].ProductIntroEn))
			fmt.Printf("Price max: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ProductPriceList[0].ProductPrice)
			fmt.Printf("Price mix: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ProductPriceList[len(summary.Result.ProductSearchResultVO.ProductList[0].ProductPriceList)-1].ProductPrice)
			fmt.Printf("Ship imm: %+v\n", summary.Result.ProductSearchResultVO.ProductList[0].ShipImmediately)
		}
	default:
		fmt.Println("не так")
	}

}

func StripHTMLTags(html string) string {
	// Regular expression to match HTML tags
	re := regexp.MustCompile(`<(?:.|\n)*?>`)

	// Replace all HTML tags with an empty string
	stripped := re.ReplaceAllString(html, "")

	return stripped
}
