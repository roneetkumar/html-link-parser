package main

import (
	"fmt"
	"strings"

	link "github.com/roneetkumar/html-link-parser"
)

var exampleHTML = `
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<a href="first-link"> Link Text
			<span> Span Text </span>
		</a>
		<a href="second-link">Hello</a>
	</body>
</html>
`

func main() {

	r := strings.NewReader(exampleHTML)

	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", links)

}
