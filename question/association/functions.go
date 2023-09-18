package association

import (
	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/util/format"
	"strings"
)

func Parse(q *input.RawQuestion) *output.Question {
	result := &output.Question{}

	setDefaults(result)
	parseHeader(q, result)
	parseSubQuestions(q, result)

	return result
}

func setDefaults(result *output.Question) {
	result.Type = "matching"
	result.GeneralFeedback.Format = "html"
	result.Penalty.Value = format.ToMoodleFloat(float32(1.0 / 3.0))
	result.Hidden.Value = 0
	result.ShuffleAnswers.Value = true
	result.CorrectFeedback.Format = "html"
	result.PartiallyCorrectFeedback.Format = "html"
	result.IncorrectFeedback.Format = "html"
}

func parseHeader(q *input.RawQuestion, result *output.Question) {
	s := strings.SplitN(q.HeaderText, ".", 2)

	result.Name = output.Name{
		Text: output.Text{
			Content: s[0],
		},
	}

	if len(s) == 1 {
		return
	}

	result.QuestionText = output.QuestionText{
		Format: "html",
		Text: output.Text{
			Content: s[1],
		},
	}
}

func parseSubQuestions(q *input.RawQuestion, result *output.Question) {
	counter := 0
	subQuestions := make(map[string]SQPair)
	for _, line := range q.Lines {
		parseOption(&line, &subQuestions, &counter)
		parseAnswer(&line, &subQuestions, &counter)
	}

	sortedQuestions := make([]SQPair, counter)
	for _, q := range subQuestions {
		sortedQuestions[q.OriginalIndex] = q
	}

	result.DefaultGrade.Value = format.ToMoodleFloat(float32(len(subQuestions)))

	result.SubQuestions = []output.SubQuestion{}
	for _, q := range sortedQuestions {
		c := output.SubQuestion{}
		c.Format = "html"
		c.Text.Content = q.Question
		c.Answer.Text.Content = q.Answer
		result.SubQuestions = append(result.SubQuestions, c)
	}
}

func parseOption(l *input.Line, sqs *map[string]SQPair, counter *int) {
	optionParts := optionMatcher.FindStringSubmatch(l.Style)
	if len(optionParts) == 0 {
		return
	}

	sqKey := optionParts[1]
	val, exists := (*sqs)[sqKey]
	if !exists {
		val = SQPair{
			OriginalIndex: *counter,
			Question:      l.Content,
		}
		*counter += 1
	}

	val.Question = l.Content
	(*sqs)[sqKey] = val
}

func parseAnswer(l *input.Line, sqs *map[string]SQPair, counter *int) {
	answerParts := answerMatcher.FindStringSubmatch(l.Style)
	if len(answerParts) == 0 {
		return
	}

	sqKey := answerParts[1]
	val, exists := (*sqs)[sqKey]
	if !exists {
		val = SQPair{
			OriginalIndex: *counter,
			Answer:        l.Content,
		}
		*counter += 1
	}

	val.Answer = l.Content
	(*sqs)[sqKey] = val
}
