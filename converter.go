package main

import (
	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/question/association"
	lt "GoMoodle/util/linetype"
)

func ConvertQuestionsToQuiz(input *[]input.RawQuestion) *output.Quiz {
	result := &output.Quiz{}

	for _, q := range *input {
		if q.Type == lt.Association {
			result.Questions = append(result.Questions, *association.Parse(&q))
		}
	}

	return result
}
