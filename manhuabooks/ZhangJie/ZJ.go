package ZhangJie

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
	"strings"
)

func ZhangJie(doc *goquery.Document) {
	vm := goja.New()
	script := doc.First().Find("script").Text()
	//fmt.Println(doc.First().Find("script").Text())
	_, _ = vm.RunString(script)

	value2 := vm.Get("photosr")

	//fmt.Println("value2.String()",value2.String())
	saar := strings.Split(value2.String(), ",")
	for _, v := range saar {
		fmt.Println(v)
	}
}
