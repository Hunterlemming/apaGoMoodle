package main

import "encoding/xml"

type Text struct {
	XMLName xml.Name `xml:"text"`
	Content string   `xml:",chardata"`
}
