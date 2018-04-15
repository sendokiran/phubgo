package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/sendokirandev/phubgo/crawler"
)

type Page struct {
	PageInfo
	Videos []VideoInfo
}

type PageInfo struct {
	Title string
}

type VideoInfo struct {
	Title string
	Thumb string
}

const ()

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := loadPage()
	renderTemplate(w, "index", page)
}

func loadPage() (page *Page) {
	crawler.RunCrawler()

	page := Page{}
	page.PageInfo{Title: "Phubgo"}
	page.VideoInfo{Title: "", Thumb: ""}

	return &page
}

func renderTemplate(w http.ResponseWriter, tmplName string, p *PageInfo) {
	tmpl, _ := template.ParseFiles(tmplName + ".html")
	tmpl.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":1337", nil))
}
