package dropdown

import (
	"testing"
)

func testDropDown(t testing.T) {

	tb_dropdowns := TaoBaoDropDownGet("ab")
	if len(tb_dropdowns) > 0 {
		t.Error("get taobao error")
	}

	tm_dropdowns := TmallDropDownGet("ac")
	if len(tm_dropdowns) > 0 {
		t.Error("get tmall error")
	}

	jd_dropdowns := JindongDropDownGet("a")
	if len(jd_dropdowns) > 0 {
		t.Error("get jd error")
	}

}
