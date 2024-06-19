package scrape

import (
    "fmt"
    "log"
	"time"
    "net/url"
    "sync"
    "EndPointer/utils"
    "github.com/gocolly/colly/v2"
)

func EndPoint(domain *string, agent *string, headers *string, proxies *string, threads *int, depth *int, timeout *int) {
    var (
        visited = make(map[string]bool)
        mu      sync.Mutex
    )

    host, err := url.Parse(*domain)
    if err != nil {
        log.Fatal(err)
    }

    var customHeaders []string
    if *headers != "" {
        customHeaders, err = utils.ReadFile(*headers)
        if err != nil {
            log.Fatal(err)
        }
    }

    c := createCollector(host.Hostname(), *depth, *threads, *proxies, time.Duration(*timeout) * time.Second)
    setCollyBehavior(c, *agent, customHeaders)

    checkAndVisit := func(link string) {
        if utils.IsSameDomain(link, host.Hostname()) && utils.IsUnique(link, visited, &mu) {
            c.Visit(link)
        }
    }

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("href"))
        checkAndVisit(link)
    })

    c.OnHTML("form[action]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("action"))
        if utils.IsSameDomain(link, host.Hostname()) && utils.IsUnique(link, visited, &mu) {
            fmt.Println(link)
        }
    })

    c.OnHTML("script[src]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("src"))
        if utils.IsSameDomain(link, host.Hostname()) && utils.IsUnique(link, visited, &mu) {
            fmt.Println(link)
        }
    })

    c.OnHTML("iframe[src]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("src"))
        if utils.IsSameDomain(link, host.Hostname()) && utils.IsUnique(link, visited, &mu) {
            fmt.Println(link)
        }
    })

    // Start scraping
    err = c.Visit(*domain)
    if err != nil {
        log.Fatal(err)
    }

    c.Wait()
}
