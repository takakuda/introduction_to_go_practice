package main

import "encoding/xml"

type Post struct {
	XMLName xml.Name `xml:"post"`
}
