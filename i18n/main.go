package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func main() {
	//先拿到要翻譯的語言 -> 讀取對應json -> 跑回圈去拿到每個詞彙的 localizationUsingJson
	lang := "id"
	// lang2 := "en"
	// lang3 := "zh"
	bundle = i18n.NewBundle(language.Chinese)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile(lang + ".json")

	localizer = i18n.NewLocalizer(bundle, language.Indonesian.String(), language.Chinese.String(), language.English.String())

	odHeaders := []string{"N", "dealN", "PN", "dealPN", "pdtNM", "pdtspec", "demandqty", "punit", "unitprice", "currencytype", "confirmqty", "cancelqty", "returnqty", "itemsumtax", "total", "header"}
	pageHeader_Name := "selema"
	for _, header := range odHeaders {
		localizeConfigOrd := i18n.LocalizeConfig{
			MessageID: header,
			TemplateData: map[string]interface{}{
				"reportType": pageHeader_Name,
			},
		}

		fmt.Println("print configord---------------", localizeConfigOrd)
		localizationUsingJson, _ := localizer.Localize(&localizeConfigOrd)
		fmt.Println(localizationUsingJson)

	}

}
