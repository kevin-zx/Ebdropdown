package dropdown

import (
	// "./httpUtil"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kevin-zx/go-util/httpUtil"
	"strconv"
)

var jd_url_format = "http://dd-search.jd.com/?ver=2&zip=1&key=%s&pvid=4a6j59vi.tnm8dd&t=1478588720776&curr_url=www.jd.com%2F&callback=jQuery1880422"

type JingdongDropDown struct {
	DropDown
	ProductCount int
}

func JindongDropDownGet(keyword string) []JingdongDropDown {
	var jddropdowns []JingdongDropDown
	urlstr := fmt.Sprintf(jd_url_format, keyword)
	header := make(map[string]string)
	header["Accept-Encoding"] = "gzip, deflate, sdch"
	header["Accept-Language"] = "en-US,en;q=0.8"
	header["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36"
	header["Accept"] = "*/*"
	header["Referer"] = "http://www.jd.com/"
	header["Connection"] = "keep-alive"
	result, err := httpUtil.GetWebConFromUrlWithHeader(urlstr, header)
	if err != nil {
		return jddropdowns
	}

	jsonStr := "[" + GetJsonStr(result) + "]"
	if jsonStr == "" {
		return jddropdowns
	}
	jsonObj, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		return jddropdowns
	}
	jsonArr, _ := jsonObj.Array()
	for rank, v := range jsonArr {
		var jddropdown JingdongDropDown
		datamap, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		dropDownKeywordData := datamap["key"]
		if dropDownKeywordData == nil {
			continue
		}
		ProductCountData, _ := datamap["qre"].(json.Number)
		productCount, _ := strconv.Atoi(ProductCountData.String())
		dropDownKeyword, _ := dropDownKeywordData.(string)
		jddropdown.Rank = rank + 1
		jddropdown.DropDownWord = dropDownKeyword
		jddropdown.ProductCount = productCount
		jddropdowns = append(jddropdowns, jddropdown)
	}
	return jddropdowns
}
