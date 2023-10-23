package main

import (
	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/question/association"
	"GoMoodle/question/multichoice"
	"GoMoodle/util/docx"
)

const inputDoc = "./input.docx"

func main() {
	sourceName, err := docx.Unzip(inputDoc)
	if err != nil {
		return
	}
	defer docx.Cleanup()

	questions := input.ParseFile(sourceName)
	quiz := ConvertQuestionsToQuiz(questions)
	output.WriteFile(quiz)
}

func init() {
	association.InitRegexMatchers()
	multichoice.InitRegexMatchers()
}
