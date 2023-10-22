package main

import (
	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/question/association"
	"GoMoodle/question/multichoice"
	lt "GoMoodle/util/linetype"
)

func ConvertQuestionsToQuiz(input *[]input.RawQuestion) *output.Quiz {
	result := &output.Quiz{}

	for _, q := range *input {
		if q.Type == lt.Association {
			result.Questions = append(result.Questions, *association.Parse(&q))
		}
		if q.Type == lt.Choice {
			result.Questions = append(result.Questions, *multichoice.Parse(&q))
		}
	}

	return result
}
