package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

//Link represents a link (<a href=""> ... </a>) in a HTML document
type Link struct {
	Href string
	Text string
}

// Parse will take in a HTML document and will return a slice of parsed Links
func Parse(r io.Reader) ([]Link, error) {

	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)

	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {

	var link Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}

	link.Text = text(n)

	return link
}

func text(n *html.Node) string {

	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var t string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t += text(c) + " "
	}

	return strings.Join(strings.Fields(t), " ")
}

func linkNodes(n *html.Node) []*html.Node {

	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var nodes []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}

	return nodes
}

// func dfs(n *html.Node, padding string) {

// 	msg := n.Data

// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}

// 	fmt.Println(padding, msg)
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"	")
// 	}

// }
