package main

import (
	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/question/association"
)

// TODO: Replace this with terminal input
const inputName = "./assets/input/association-test.xml"

// TODO: Replace this with terminal input
const outputName = "./assets/output/out.xml"

func main() {
	questions := input.ParseFile(inputName)
	quiz := ConvertQuestionsToQuiz(questions)
	output.WriteFile(outputName, quiz)
}

func init() {
	association.InitRegexMatchers()
}
