package dropdown

import (
	// "./httpUtil"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kevin-zx/go-util/httpUtil"
	"strconv"
)

var tb_url_format string = "https://suggest.taobao.com/sug?code=utf-8&q=%s&callback=jsonp1728&area=c2c"

type TaobaoDropDown struct {
	DropDown
	ProductCount int
}

//get taobao dropDown by keyword
func TaoBaoDropDownGet(keyword string) []TaobaoDropDown {
	var tbdropdowns []TaobaoDropDown
	urlstr := fmt.Sprintf(tb_url_format, keyword)
	result, err := httpUtil.GetWebConFromUrl(urlstr)
	if err != nil {
		return tbdropdowns
	}
	jsonStr := GetJsonStr(result)
	if jsonStr == "" {
		return tbdropdowns
	}
	json, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		return tbdropdowns
	}
	resultArr, err := json.Get("result").Array()
	if err != nil {
		return tbdropdowns
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
		var dropdown TaobaoDropDown
		dropdown.MainKeyword = keyword
		dropdown.DropDownWord = dropDownkeyword
		dropdown.Rank = rank + 1
		dropdown.ProductCount, _ = strconv.Atoi(productCount)
		tbdropdowns = append(tbdropdowns, dropdown)
	}
	return tbdropdowns
}
