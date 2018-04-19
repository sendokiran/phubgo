package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/sendokirandev/phubgo/config"
	"golang.org/x/net/html"
)

const (
	ErrHTTPGet   string = "Failed to get specified URL: %v"
	AssetsFolder string = "assets"
)

var (
	conf = config.New()
)

type Item struct {
	Thumbs []Video
}

type Video struct {
	Key string
	Gif string
}

func crawler(b io.Reader) (data *Item) {
	re := regexp.MustCompile(`\/(\S*).php\?(\S*)=ph`)
	z := html.NewTokenizer(b)
	data = &Item{Thumbs: []Video{}}
	video := Video{}
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
					// vidKeys = append(vidKeys, vidKey[len(vidKey)-1])
					video.Key = vidKey[len(vidKey)-1]
				}
				continue
			case "img":
				for _, tokenAttr := range token.Attr {
					isMediaBook := tokenAttr.Key == "data-mediabook"
					if !isMediaBook {
						continue
					}

					// gifKeys = append(gifKeys, tokenAttr.Val)
					video.Gif = tokenAttr.Val
				}
				continue
			}
		}
		data.Thumbs = append(data.Thumbs, video)
	}
	return
}

func RunCrawler() {
	// Get body response from given url
	url := conf.CrawlerInfo.URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf(ErrHTTPGet, err)
	}

	// Crawl over tokens and get the values
	b := resp.Body
	defer b.Close()
	data := crawler(b)
	fmt.Println(data)

	// filename := fmt.Sprintf("assets/thumbs/%s.webm", vidKeys)
	// isFileExist := func() bool {
	// 	_, err = os.Stat(filename)
	// 	if err != nil {
	// 		return os.IsNotExist(err)
	// 	}
	// 	return true
	// }

	// if !isFileExist() {
	// 	fmt.Println("buat")
	// }
}
