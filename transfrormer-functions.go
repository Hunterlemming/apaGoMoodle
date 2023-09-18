package main

func ConvertQuestionsToQuiz(input *[]Question) *Quiz {
	result := &Quiz{}

	for _, q := range *input {
		if q.Type == QTAssociation {
			result.Questions = append(result.Questions, *parseAssociation(&q))
		}
	}

	return result
}

func parseAssociation(q *Question) *XQuestion {
	result := &XQuestion{}

	result.Type = "matching"
	result.Name = Name{
		Text: Text{
			Content: q.Text,
		},
	}
	parseAssociationQT(q, result)

	return result
}

func parseAssociationQT(q *Question, result *XQuestion) {

	result.QuestionText = QuestionText{
		Format: "html",
		Text: Text{
			Content: "",
		},
	}
}
