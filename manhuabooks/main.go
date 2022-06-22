package main

import (
	"manhuabooks/service"
	"manhuabooks/utils"
)

const (
	SocksProxy = "socks5://192.168.1.18:8000"
)

func main() {
	//MuLu(JieXi("https://www.iimanhua.cc/comic/2988/index.html"))
	//ZhangJie(JieXi("https://www.iimanhua.cc/comic/2988/945453.html"))

	//JsCeShi()
	//ZhangJie(JieXi("https://www.iimanhua.cc/comic/2988/945453.html"))
	//FenLei(JieXi("https://www.iimanhua.cc/shaonianrexue/"))
	service.ZhangJie(utils.JieXi("https://www.iimanhua.cc/comic/2988/945453.html"))
}

//
//func MuLu(doc *goquery.Document) {
//	doc.Find("li").Each(func(i int, s *goquery.Selection) {
//		//log.Println(s.Find("a").Text())
//		href, _ := s.Find("a").Attr("href")
//		fmt.Println("---href---", href)
//		fmt.Println(s.Find("a").Text())
//	})
//}
//func JieXi(url string) (document *goquery.Document) {
//
//	res, err := http.Get(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//	if res.StatusCode != 200 {
//		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
//	}
//	utf8Body, err := iconv.NewReader(res.Body, "gb2312", "utf-8")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	document, err = goquery.NewDocumentFromReader(utf8Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return
//}
//
//func JsCeShi() (ss string) {
//
//	vm := goja.New()
//	str := "function base64decode(str) {\n    var base64EncodeChars = \"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/\";\n    var base64DecodeChars = new Array(-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1, -1, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1, -1, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1);\n    var c1, c2, c3, c4;\n    var i, len, out;\n    len = str.length;\n    i = 0;\n    out = \"\";\n    while (i < len) {\n        do {\n            c1 = base64DecodeChars[str.charCodeAt(i++) & 255]\n        } while (i < len && c1 == -1);\n        if (c1 == -1) {\n            break\n        }\n        do {\n            c2 = base64DecodeChars[str.charCodeAt(i++) & 255]\n        } while (i < len && c2 == -1);\n        if (c2 == -1) {\n            break\n        }\n        out += String.fromCharCode((c1 << 2) | ((c2 & 48) >> 4));\n        do {\n            c3 = str.charCodeAt(i++) & 255;\n            if (c3 == 61) {\n                return out\n            }\n            c3 = base64DecodeChars[c3]\n        } while (i < len && c3 == -1);\n        if (c3 == -1) {\n            break\n        }\n        out += String.fromCharCode(((c2 & 15) << 4) | ((c3 & 60) >> 2));\n        do {\n            c4 = str.charCodeAt(i++) & 255;\n            if (c4 == 61) {\n                return out\n            }\n            c4 = base64DecodeChars[c4]\n        } while (i < len && c4 == -1);\n        if (c4 == -1) {\n            break\n        }\n        out += String.fromCharCode(((c3 & 3) << 6) | c4)\n    }\n    return out\n};\nvar ret_classurl = '/mh/2988/';\nvar comicname = \"异皇重生\";\nvar viewid = \"945951\";\nvar viewtype = \"1\";\nvar viewname = \"298混乱的真相\";\nvar photosr = new Array();\n\npacked = \"ZXZhbChmdW5jdGlvbihwLGEsYyxrLGUsZCl7ZT1mdW5jdGlvbihjKXtyZXR1cm4gYy50b1N0cmluZygzNil9O2lmKCEnJy5yZXBsYWNlKC9eLyxTdHJpbmcpKXt3aGlsZShjLS0pe2RbZShjKV09a1tjXXx8ZShjKX1rPVtmdW5jdGlvbihlKXtyZXR1cm4gZFtlXX1dO2U9ZnVuY3Rpb24oKXtyZXR1cm4nXFx3Kyd9O2M9MX07d2hpbGUoYy0tKXtpZihrW2NdKXtwPXAucmVwbGFjZShuZXcgUmVnRXhwKCdcXGInK2UoYykrJ1xcYicsJ2cnKSxrW2NdKX19cmV0dXJuIHB9KCdjWzFdPSJlL2IvYS9kL2cvbS5mLzAiO2NbMl09ImUvYi9hL2QvZy9vLmYvMCI7Y1szXT0iZS9iL2EvZC9nL2wuZi8wIjtjWzRdPSJlL2IvYS9kL2cvcC5mLzAiO2NbNV09ImUvYi9hL2QvZy9rLmYvMCI7Y1s2XT0iZS9iL2EvZC9nL2guZi8wIjtjWzddPSJlL2IvYS9kL2cvaS5mLzAiO2NbOF09ImUvYi9hL2QvZy9qLmYvMCI7Y1s5XT0iZS9iL2EvZC9nL24uZi8wIjtjW3FdPSJlL2IvYS9kL2cveS5mLzAiO2NbeF09ImUvYi9hL2QvZy96LmYvMCI7Y1t2XT0iZS9iL2EvZC9nL3cuZi8wIjtjW3JdPSJlL2IvYS9kL2cvcy5mLzAiO2NbdF09ImUvYi9hL2QvZy91LmYvMCI7JywzNiwzNiwnfHx8fHx8fHx8fDA2fDIwMjJ8cGhvdG9zcnwwOHxpbWFnZXMyfGpwZ3wwOXw0NDEwYWI5YjQyfDQ0ZGM3NzRlYTh8NDRlZTNiZjI2Znw0NDA4Y2I4OTY5fDQ0MjNhYzVmMjh8NDQ2OWI5NWY1ZHw0NGRmMDY2MWQyfDQ0NTQwZGQyYTh8NDRhZWEwN2I5MHwxMHwxM3w0NWM0YTU5OTQ3fDE0fDQ1ZTRjYWEzYTF8MTJ8NDU2ZjFiOTAyYnwxMXw0NDlhZDIyNWFmfDQ0YmUxNTFlY2YnLnNwbGl0KCd8JyksMCx7fSkpCg==\";\neval(eval(base64decode(packed).slice(4)));"
//	_, _ = vm.RunString(str)
//
//	value2 := vm.Get("photosr")
//
//	//fmt.Println("value2.String()",value2.String())
//	saar := strings.Split(value2.String(), ",")
//	for _, v := range saar {
//		fmt.Println(v)
//	}
//	return
//}
//
//func FenLei(doc *goquery.Document) {
//	doc.Find("div").Text()
//	//doc.Find("div").Each(func(i int, s *goquery.Selection) {
//	//	//log.Println(s.Find("a").Text())
//	//	href, _ := s.Find("a").Attr("href")
//	//	img:=s.Find("img").Text()
//	//	fmt.Println("---img---", img)
//	//	fmt.Println("---href---", href)
//	//	fmt.Println(s.Find("a").Text())
//	//})
//}
