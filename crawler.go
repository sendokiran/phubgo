package crawler

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/sendokirandev/phubgo/config"
	"golang.org/x/net/html"
)

const (
	ErrHTTPGet    string = "Failed to get specified URL: %v"
	ErrConfDecode string = "Failed to decode config: %v"
	AssetsFolder  string = "assets"
)

func crawler(b io.Reader) (vidKeys, gifKeys []string) {
	re := regexp.MustCompile(`\/(\S*).php\?(\S*)=ph`)
	z := html.NewTokenizer(b)
	for {
		tokenType := z.Next()
		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			token := z.Token()

			switch token.Data {
			case "a":
				for _, tokenAttr := range token.Attr {
					isHref := tokenAttr.Key == "href" && re.MatchString(tokenAttr.Val)
					if !isHref {
						continue
					}

					vidKey := strings.Split(tokenAttr.Val, "=")
					vidKeys = append(vidKeys, vidKey[len(vidKey)-1])
				}
				continue
			case "img":
				for _, tokenAttr := range token.Attr {
					isMediaBook := tokenAttr.Key == "data-mediabook"
					if !isMediaBook {
						continue
					}

					gifKeys = append(gifKeys, tokenAttr.Val)
				}
				continue
			}
		}
	}
	return
}

func RunCrawler() {
	// Get body response from given url
	url := config.Crawler.URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf(ErrHTTPGet, err)
	}

	// Crawl over tokens and get the values
	b := resp.Body
	defer b.Close()
	vidKeys, gifKeys := crawler(b)

	// filename := "assets/thumbs/%s.webm", vidKeys
	// isFileExist := func() (ok bool, err error) {
	// 	_, err := os.Stat(filename)
	// 	ok := os.IsNotExist(err)
	// }
	// ok, err := isFileExist()

	// fmt.Println("gifKeys")
}
