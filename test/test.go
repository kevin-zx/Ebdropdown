package main

import (
	"../../dropdown"
	"fmt"
)

func main() {
	// tb_dropdowns := dropdown.TaoBaoDropDownGet("ab")
	// for _, dropdown := range tb_dropdowns {
	// 	fmt.Printf("tb_dropdownKeyword:%s,rank:%d\n", dropdown.DropDownWord, dropdown.Rank)
	// }

	// tm_dropdowns := dropdown.TmallDropDownGet("ac")
	// for _, dropdown := range tm_dropdowns {
	// 	fmt.Printf("tm_dropdownKeyword:%s,rank:%d\n", dropdown.DropDownWord, dropdown.Rank)
	// }

	// jd_dropdowns := dropdown.JindongDropDownGet("s")
	// for _, dropdown := range jd_dropdowns {
	// 	fmt.Printf("jd_dropdownKeyword:%s,rank:%d,productCount:%d\n", dropdown.DropDownWord, dropdown.Rank, dropdown.ProductCount)
	// }
	al_dropdowns := dropdown.AlibabaDropDownGet("s")
	for _, dropdown := range al_dropdowns {
		fmt.Printf("al_dropdownKeyword:%s,rank:%d,productCount:%d\n", dropdown.DropDownWord, dropdown.Rank, dropdown.ProductCount)
	}
}
