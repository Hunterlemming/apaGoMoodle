package docx

import "encoding/xml"

type PStyle struct {
	XMLName xml.Name `xml:"pStyle"`
	Val     string   `xml:"val,attr"`
}

type RFonts struct {
	XMLName  xml.Name `xml:"rFonts"`
	EastAsia string   `xml:"eastAsia,attr"`
}

type RPr struct {
	XMLName   xml.Name  `xml:"rPr"`
	RFonts    RFonts    `xml:"rFonts"`
	B         bool      `xml:"b,omitempty"`
	BCs       bool      `xml:"bCs,omitempty"`
	VertAlign VertAlign `xml:"vertAlign"`
}

type VertAlign struct {
	XMLName xml.Name `xml:"vertAlign"`
	Val     string   `xml:"val,attr"`
}

type T struct {
	XMLName   xml.Name `xml:"t"`
	SpacePres string   `xml:"xml:space,attr"`
	Content   string   `xml:",chardata"`
}

type BR struct {
	XMLName xml.Name `xml:"br"`
}

type R struct {
	XMLName xml.Name `xml:"r"`
	RPr     RPr      `xml:"rPr"`
	T       T        `xml:"t"`
	BR      BR       `xml:"br"`
}

type P struct {
	XMLName xml.Name `xml:"p"`
	PStyle  PStyle   `xml:"pPr>pStyle"`
	R       []R      `xml:"r"`
}
