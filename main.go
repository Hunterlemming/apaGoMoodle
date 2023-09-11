package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// TODO: Replace this with terminal input
const fileName = "./files/association-test.xml"

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	questions, err := ParseQuestions(decoder)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(questions)
}
