package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html>
		<head>
		<title>Title of the document</title>
		</head>

		<body>
		<h1> ` + name + `</h1>
		</body>

		</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating the file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
