package main

import (
	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/question/association"
	"GoMoodle/question/multichoice"
	"GoMoodle/util/docx"
)

// TODO: Replace this with terminal input
// const inputName = "./assets/input/assoc-test-2.xml"
const inputName = "./assets/input/mult-ans-test.xml"

// TODO: Replace this with terminal input
const outputName = "./assets/output/out.xml"

const (
	inputDoc  = "./input.docx"
	outputXML = "./output.xml"
)

func main() {
	sourceName, err := docx.Unzip(inputDoc)
	if err != nil {
		return
	}
	defer docx.Cleanup()

	questions := input.ParseFile(sourceName)
	quiz := ConvertQuestionsToQuiz(questions)
	output.WriteFile(outputXML, quiz)
}

func init() {
	association.InitRegexMatchers()
	multichoice.InitRegexMatchers()
}
