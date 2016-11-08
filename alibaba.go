package dropdown

import (
	// "./httpUtil"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kevin-zx/go-util/httpUtil"
	"strconv"
)

var al_url_format string = "https://suggest.1688.com/bin/suggest?q=%s&encode=utf8"

type AlibabaDropDown struct {
	DropDown
	ProductCount int
}

//get Alibaba dropDown by keyword
func AlibabaDropDownGet(keyword string) []AlibabaDropDown {
	var aldropdowns []AlibabaDropDown
	urlstr := fmt.Sprintf(al_url_format, keyword)
	result, err := httpUtil.GetWebConFromUrl(urlstr)
	if err != nil {
		return aldropdowns
	}
	jsonStr := GetJsonStr(result)
	if jsonStr == "" {
		return aldropdowns
	}
	json, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		return aldropdowns
	}
	resultArr, err := json.Get("result").Array()
	if err != nil {
		return aldropdowns
	}

	for rank, item := range resultArr {
		itemData, ok := item.([]interface{})
		var dropDownkeyword, productCount string
		if !ok {
			continue
		}
		if len(itemData) > 0 {
			dropDownkeyword, _ = itemData[0].(string)
		}
		if len(itemData) > 1 {
			productCount, _ = itemData[1].(string)
		}
		var dropdown AlibabaDropDown
		dropdown.MainKeyword = keyword
		dropdown.DropDownWord = dropDownkeyword
		dropdown.Rank = rank + 1
		dropdown.ProductCount, _ = strconv.Atoi(productCount)
		aldropdowns = append(aldropdowns, dropdown)
	}
	return aldropdowns
}
