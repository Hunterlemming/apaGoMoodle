package main

import (
	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/question/association"
	"GoMoodle/question/multichoice"
)

// TODO: Replace this with terminal input
// const inputName = "./assets/input/assoc-test-2.xml"
const inputName = "./assets/input/mult-ans-test.xml"

// TODO: Replace this with terminal input
const outputName = "./assets/output/out.xml"

func main() {
	questions := input.ParseFile(inputName)
	quiz := ConvertQuestionsToQuiz(questions)
	output.WriteFile(outputName, quiz)
}

func init() {
	association.InitRegexMatchers()
	multichoice.InitRegexMatchers()
}
