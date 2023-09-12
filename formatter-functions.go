package main

import (
	"encoding/xml"
	"os"
)

func writeFile(fileName string) {
	t := &Text{Content: "hi"}
	header := []byte(xml.Header)
	out, _ := xml.MarshalIndent(t, " ", "  ")

	err := os.WriteFile(fileName, append(header, out...), 0644)
	if err != nil {
		panic(err)
	}
}
