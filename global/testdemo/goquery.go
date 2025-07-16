package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func main() {
	reader, err := os.Open("uploads/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(reader)
	selection := doc.Find("title")
	fmt.Println(selection)
	fmt.Println(selection.Text())
	selection.SetText("小丁知道")
	fmt.Println(selection.Text())
	selection.Add("title")
	html, err := doc.Html()
	fmt.Println(html)
}
