package main

import (
	"fmt"
	"strings"
)

func ConvertQuestionsToQuiz(input *[]RawQuestion) *Quiz {
	result := &Quiz{}

	for _, q := range *input {
		if q.Type == QTAssociation {
			result.Questions = append(result.Questions, *parseAssociation(&q))
		}
	}

	return result
}

func parseAssociation(q *RawQuestion) *XQuestion {
	result := &XQuestion{}

	result.Type = "matching"
	parseAssociationHeader(q, result)
	parseAssociationLine(q, result)

	return result
}

type AssociationSQPair struct {
	Question string
	Answer   string
}

func (p AssociationSQPair) String() string {
	return fmt.Sprintf("[Question]: %v; [Answer]: %v", p.Question, p.Answer)
}

func parseAssociationHeader(q *RawQuestion, result *XQuestion) {
	s := strings.SplitN(q.HeaderText, ".", 2)

	result.Name = Name{
		Text: Text{
			Content: s[0],
		},
	}

	if len(s) == 1 {
		return
	}

	result.QuestionText = QuestionText{
		Format: "html",
		Text: Text{
			Content: s[1],
		},
	}
}

func parseAssociationLine(q *RawQuestion, result *XQuestion) {
	subQuestions := make(map[string]AssociationSQPair)
	for _, line := range q.Lines {
		parseAssociationOption(&line, &subQuestions)
		parseAssociationAnswer(&line, &subQuestions)
	}

	result.SubQuestions = []SubQuestion{}
	for _, q := range subQuestions {
		result.SubQuestions = append(result.SubQuestions, SubQuestion{
			Format: "html",
			Text: Text{
				Content: q.Question,
			},
			Answer: Answer{
				Text: Text{
					Content: q.Answer,
				},
			},
		})
	}
}

func parseAssociationOption(l *Line, sqs *map[string]AssociationSQPair) {
	optionParts := associationOptionMatcher.FindStringSubmatch(l.Style)
	if len(optionParts) == 0 {
		return
	}

	sqKey := optionParts[1]
	val, exists := (*sqs)[sqKey]
	if !exists {
		val = AssociationSQPair{
			Question: l.Content,
		}
	}

	val.Question = l.Content
	(*sqs)[sqKey] = val
}

func parseAssociationAnswer(l *Line, sqs *map[string]AssociationSQPair) {
	answerParts := associationAnswerMatcher.FindStringSubmatch(l.Style)
	if len(answerParts) == 0 {
		return
	}

	sqKey := answerParts[1]
	val, exists := (*sqs)[sqKey]
	if !exists {
		val = AssociationSQPair{
			Answer: l.Content,
		}
	}

	val.Answer = l.Content
	(*sqs)[sqKey] = val
}
