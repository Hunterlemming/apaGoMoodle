package input

import (
	"GoMoodle/util/docx"
	lt "GoMoodle/util/linetype"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func ParseFile(inputName string) *[]RawQuestion {
	f, err := os.Open(inputName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	questions, err := parseQuestions(decoder)
	if err != nil {
		fmt.Println(err)
	}

	return &questions
}

func parseQuestions(decoder *xml.Decoder) (questions []RawQuestion, err error) {
	var question = &RawQuestion{}

	for {
		// Check for EOF
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		// Check for paragraph
		var startElement, ok = token.(xml.StartElement)
		var element = strings.Trim(strings.TrimSuffix(startElement.Name.Local, "\n"), " ")
		if !ok || element != "p" {
			continue
		}

		// Parse paragraph
		line, err := parseLine(decoder, &startElement)
		if err != nil {
			return []RawQuestion{}, err
		}

		// New Question
		if question.HeaderText == "" {
			question.HeaderText = line.Content
			question.Type = line.Style
			continue
		}

		// End Question
		if (line.Style == lt.None || line.Style == lt.Normal) && line.Content == "" {
			questions = append(questions, *question)
			question = &RawQuestion{}
			continue
		}

		// Add answer to Question
		question.Lines = append(question.Lines, line)
	}

	return questions, nil
}

func parseLine(decoder *xml.Decoder, startElement *xml.StartElement) (result Line, err error) {
	var paragraph docx.P
	if err := decoder.DecodeElement(&paragraph, startElement); err != nil {
		return Line{}, err
	}

	var content []string
	for _, element := range paragraph.R {
		content = append(content, element.T.Content)
	}

	result.Style = paragraph.PStyle.Val
	result.Content = strings.Join(content, "")

	return result, nil
}
