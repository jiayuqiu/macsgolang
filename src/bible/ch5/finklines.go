package ch5

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	fmt.Printf("%v %v \n", n.Type, n.Data)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func Findlinks1() {
	url := strings.NewReader("https://golang.org")
	// doc, err := html.Parse(os.Stdin)
	doc, err := html.Parse(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range Practice51Visit(nil, doc) {
		fmt.Println(link)
	}
}

func Practice51Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	fmt.Printf("%v %v \n", n.Type, n.Data)

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 改为递归方式
	links = Practice51Visit(links, n.FirstChild) // 添加左叶子节点对应的Links
	// links = Practice51Visit(links, n.NextSibling) // 添加右叶子节点对应的links
	return Practice51Visit(links, n.NextSibling)
}
