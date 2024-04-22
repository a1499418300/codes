package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
func main() {
	website := "https://cdn.7277.cn/b/N347I18230q23"
	resp, err := http.Get(website)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	root, err := url.Parse(website)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("src")
		if exists {
			url, err := url.Parse(val)
			if err != nil {
				panic(err)
			}
			path := root.ResolveReference(url).Path
			dir, file := filepath.Split(path)
			os.MkdirAll(dir, os.ModePerm)
			err = downloadFile(filepath.Join(dir, file), root.ResolveReference(url).String())
			if err != nil {
				panic(err)
			}
			s.SetAttr("src", filepath.Join(dir, file))
		}
	})
	html, err := doc.Html()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("index.html", []byte(html), os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("Website downloaded successfully")
}
