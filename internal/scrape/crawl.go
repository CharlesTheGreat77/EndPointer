package scrape

import (
	"mole/utils"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// function that uses colly to initially scrape for URLs
func EndPoint(domain *string, agent *string, headers *string, proxies *string, threads *int, depth *int, timeout *int) {
	var visited []string

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

	c := createCollector(host.Hostname(), *depth, *threads, *proxies, time.Duration(*timeout)*time.Second)
	setCollyBehavior(c, *agent, customHeaders)

	// regex for paths in HTML content
	regexPatterns := []string{
		`http://(/?%3C=(%22|%27| + "" + ))\/[a-zA-Z0-9_?&=\/\-#\.]*[%22|'|%60]`,
		`http://(/?%3C=(%22|%27| + "" + ))\/[a-zA-Z0-9_?&=\/\-\#\.]*([%22|\'|%60])`,
	}

	regexes := make([]*regexp.Regexp, len(regexPatterns))
	for i, pattern := range regexPatterns {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatal(err)
		}
		regexes[i] = regex
	}

	c.OnHTML("form[action]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("action"))
		if !utils.HasVisited(link, visited) {
			visited = append(visited, link)
			fmt.Println(link)
			e.Request.Visit(link)
		}
	})

	c.OnHTML("a[href], link[href], script[src], iframe[src], img[src]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link == "" {
			link = e.Request.AbsoluteURL(e.Attr("src"))
		}
		if link != "" && !utils.HasVisited(link, visited) {
			visited = append(visited, link)
			fmt.Println(link)
			e.Request.Visit(link)
		}
	})

	c.OnHTML("meta[http-equiv=refresh][content]", func(e *colly.HTMLElement) {
		content := e.Attr("content")
		if urlIdx := strings.Index(content, "url="); urlIdx != -1 {
			link := e.Request.AbsoluteURL(content[urlIdx+4:])
			if !utils.HasVisited(link, visited) {
				visited = append(visited, link)
				fmt.Println(link)
				e.Request.Visit(link)
			}
		}
	})

	err = c.Visit(*domain)
	if err != nil {
		log.Fatal(err)
	}

	c.Wait()
}
