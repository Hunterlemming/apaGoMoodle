package main

import (
	"regexp"
)

// TODO: Replace this with terminal input
const inputName = "./files/input/association-test.xml"

// TODO: Replace this with terminal input
const outputName = "./files/output/out.xml"

var associationOptionMatcher *regexp.Regexp

func main() {
	questions := ParseFile(inputName)
	quiz := ConvertQuestionsToQuiz(questions)
	writeFile(outputName, quiz)
}

func init() {
	associationOptionMatcher, _ = regexp.Compile(QTAssociationOptionRE)
}
