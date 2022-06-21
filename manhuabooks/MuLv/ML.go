package MuLv

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func MuLu(doc *goquery.Document) {
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		//log.Println(s.Find("a").Text())
		href, _ := s.Find("a").Attr("href")
		fmt.Println("---href---", href)
		fmt.Println(s.Find("a").Text())
	})
}
