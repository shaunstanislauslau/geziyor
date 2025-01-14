# Geziyor
Geziyor is a blazing fast web crawling and web scraping framework. It can be used to crawl websites and extract structured data from them. Geziyor is useful for a wide range of purposes such as data mining, monitoring and automated testing. 

[![GoDoc](https://godoc.org/github.com/geziyor/geziyor?status.svg)](https://godoc.org/github.com/geziyor/geziyor)
[![report card](https://goreportcard.com/badge/github.com/geziyor/geziyor)](http://goreportcard.com/report/geziyor/geziyor)

## Features
- 1.000+ Requests/Sec
- JS Rendering
- Caching (Memory/Disk)
- Automatic Data Exporting (JSON, CSV, or custom)
- Limit Concurrency (Global/Per Domain)
- Request Delays (Constant/Randomized)
- Automatic response decoding to UTF-8

See scraper [Options](https://godoc.org/github.com/geziyor/geziyor#Options) for all custom settings. 

## Status
Since the project is in **development phase**, **API may change in time**. Thus, we highly recommend you to use Geziyor with go modules.

## Usage
Simple usage 

```go
geziyor.NewGeziyor(geziyor.Options{
    StartURLs: []string{"http://api.ipify.org"},
    ParseFunc: func(r *geziyor.Response) {
        fmt.Println(string(r.Body))
    },
}).Start()
```

Advanced usage

```go
func main() {
    geziyor.NewGeziyor(geziyor.Options{
        StartURLs: []string{"http://quotes.toscrape.com/"},
        ParseFunc: quotesParse,
        Exporters: []geziyor.Exporter{exporter.JSONExporter{}},
    }).Start()
}

func quotesParse(r *geziyor.Response) {
    r.DocHTML.Find("div.quote").Each(func(i int, s *goquery.Selection) {
        r.Exports <- map[string]interface{}{
            "text":   s.Find("span.text").Text(),
            "author": s.Find("small.author").Text(),
        }
    })
    if href, ok := r.DocHTML.Find("li.next > a").Attr("href"); ok {
        r.Geziyor.Get(r.JoinURL(href), quotesParse)
    }
}
```

See [tests](https://github.com/geziyor/geziyor/blob/master/geziyor_test.go) for more usage examples.

## Installation

    go get github.com/geziyor/geziyor
