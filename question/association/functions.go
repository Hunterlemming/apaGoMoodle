package association

import (
	"strings"

	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/util/format"
)

func Parse(q *input.RawQuestion) *output.Question {
	result := &output.Question{}
	rawSubQuestions := extractRawSubQuestions(q)

	setDefaults(result)
	parseHeader(q, rawSubQuestions, result)
	parseSubQuestions(rawSubQuestions, result)

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

func extractRawSubQuestions(q *input.RawQuestion) []SQPair {
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

	return sortedQuestions
}

func parseHeader(q *input.RawQuestion, rawSubQuestions []SQPair, result *output.Question) {
	result.Name = output.Name{
		Text: output.Text{
			Content: q.HeaderText,
		},
	}

	s := strings.SplitN(q.HeaderText, ".", 2)
	if len(s) == 1 {
		return
	}

	var sqNames []string
	for _, question := range rawSubQuestions {
		nameParts := subQuestionAnswerMatcher.FindStringSubmatch(question.Question)
		if len(nameParts) == 3 {
			sqNames = append(sqNames, nameParts[2])
			continue
		}

		sqNames = append(sqNames, question.Question)
	}

	result.QuestionText = output.QuestionText{
		Format: "html",
		Text: output.Text{
			Content: format.ToMoodleQuestionText(s[1], sqNames),
		},
	}
}

func parseSubQuestions(rawSubQuestions []SQPair, result *output.Question) {
	result.DefaultGrade.Value = format.ToMoodleFloat(float32(len(rawSubQuestions)))

	result.SubQuestions = []output.SubQuestion{}
	for _, q := range rawSubQuestions {
		c := output.SubQuestion{}
		textContent := format.ToMoodleParagraph(q.Answer)
		answerTextContent := q.Question
		answerParts := subQuestionAnswerMatcher.FindStringSubmatch(q.Question)
		if len(answerParts) == 3 {
			// "A) Good answer" -> "Good answer"
			answerTextContent = answerParts[2]
		}

		c.Format = "html"
		c.Text.Content = textContent
		c.Answer.Text.Content = answerTextContent
		result.SubQuestions = append(result.SubQuestions, c)
	}
}

func parseOption(l *input.Line, sqs *map[string]SQPair, counter *int) {
	optionParts := optionDocFormattingMatcher.FindStringSubmatch(l.Style)
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
	answerParts := answerDocFormattingMatcher.FindStringSubmatch(l.Style)
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
