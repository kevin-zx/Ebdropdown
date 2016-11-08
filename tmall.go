package dropdown

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kevin-zx/go-util/httpUtil"
	"strconv"
)

var tm_url_format string = "https://suggest.taobao.com/sug?code=utf-8&q=%s&callback=jsonp2869&area=b2c&code=utf-8"

type TmallDropDown struct {
	DropDown
	ProductCount int
}

func TmallDropDownGet(keyword string) []TmallDropDown {
	var tmdropDowns []TmallDropDown
	urlStr := fmt.Sprintf(tm_url_format, keyword)
	result, err := httpUtil.GetWebConFromUrl(urlStr)
	if err != nil {
		return tmdropDowns
	}
	jsonStr := GetJsonStr(result)
	if jsonStr == "" {
		return tmdropDowns
	}
	json, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		return tmdropDowns
	}
	resultArr, err := json.Get("result").Array()
	if err != nil {
		return tmdropDowns
	}

	for rank, item := range resultArr {
		var dropDownkeyword, productCount string
		itemData, ok := item.([]interface{})
		if !ok {
			continue
		}
		if len(itemData) > 0 {
			dropDownkeyword, _ = itemData[0].(string)
		}
		if len(itemData) > 1 {
			productCount, _ = itemData[1].(string)
		}
		var dropdown TmallDropDown
		dropdown.MainKeyword = keyword
		dropdown.DropDownWord = dropDownkeyword
		dropdown.Rank = rank + 1
		dropdown.ProductCount, _ = strconv.Atoi(productCount)
		tmdropDowns = append(tmdropDowns, dropdown)
	}
	return tmdropDowns
}
