package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sendokirandev/phubgo/crawler"
)

type Page struct {
	PageInfo
	VideoInfo
}

type PageInfo struct {
	Title string
}

type VideoInfo struct {
	Title string
	Thumb string
}

const (
	ErrServer string = "Failed to start server: %v"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := loadPage()
	renderTemplate(w, "index", page)
}

func crawlerHandler(w http.ResponseWriter, r *http.Request) {
	crawler.RunCrawler()
}

func loadPage() (page *Page) {
	page = &Page{}
	page.PageInfo.Title = "Phubgo"
	return
}

func renderTemplate(w http.ResponseWriter, tmplName string, p *Page) {
	tmpl, _ := template.ParseFiles(tmplName + ".html")
	tmpl.Execute(w, p.PageInfo)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/crawler/", crawlerHandler)

	fmt.Println("Server will be running, please visit http://0.0.0.0:1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}
