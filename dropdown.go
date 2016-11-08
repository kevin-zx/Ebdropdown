package dropdown

import (
	// "fmt"
	"github.com/kevin-zx/go-util/regexpUtil"
)

type DropDown struct {
	DropDownWord string
	Rank         int
	MainKeyword  string
}

func GetJsonStr(text string) string {
	jsonResult := regexpUtil.FindString("\\{.+\\}", text)
	// fmt.Println(jsonResult[0])
	if len(jsonResult) > 0 {
		return jsonResult[0]
	} else {
		return ""
	}
}
