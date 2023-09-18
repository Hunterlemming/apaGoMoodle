package main

import (
	"regexp"
)

// TODO: Replace this with terminal input
const inputName = "./files/input/association-test.xml"

// TODO: Replace this with terminal input
const outputName = "./files/output/out.xml"

var associationOptionMatcher *regexp.Regexp
var associationAnswerMatcher *regexp.Regexp

func main() {
	questions := ParseFile(inputName)
	quiz := ConvertQuestionsToQuiz(questions)
	writeFile(outputName, quiz)
}

func init() {
	var err error

	associationOptionMatcher, err = regexp.Compile(QTAssociationOptionRE)
	if err != nil {
		panic(err)
	}

	associationAnswerMatcher, err = regexp.Compile(QTAssociationAnswerRE)
	if err != nil {
		panic(err)
	}
}
