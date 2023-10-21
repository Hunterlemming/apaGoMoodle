package association

import (
	"strings"

	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/util/format"
)

func Parse(q *input.RawQuestion) *output.Question {
	result := &output.Question{}
	optionMap, sortedOptions, questions := extractRawSubQuestions(q)

	setDefaults(result)
	parseHeader(q, &sortedOptions, result)
	parseSubQuestions(questions, result, &optionMap)

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

func extractRawSubQuestions(q *input.RawQuestion) (options map[string]SQOption, sortedOptions []string, result []SQPair) {
	options = make(map[string]SQOption)
	sortedOptions = []string{}
	result = []SQPair{}

	// The first few lines are the options, so we can take it for granted
	for _, line := range q.Lines {
		parseOption(&line, &options, &sortedOptions)
		parseAnswer(&line, &result, &options)
	}

	return
}

func parseHeader(q *input.RawQuestion, options *[]string, result *output.Question) {
	result.Name = output.Name{
		Text: output.Text{
			Content: q.HeaderText,
		},
	}

	s := strings.SplitN(q.HeaderText, ".", 2)
	if len(s) == 1 {
		return
	}

	var optionNames []string
	for _, name := range *options {
		fName := format.ToMoodleOptionName(name, subQuestionAnswerMatcher)
		optionNames = append(optionNames, fName)
	}

	result.QuestionText = output.QuestionText{
		Format: "html",
		Text: output.Text{
			Content: format.ToMoodleQuestionText(s[1], optionNames),
		},
	}
}

func parseSubQuestions(rawSubQuestions []SQPair, result *output.Question, optionMap *map[string]SQOption) {
	result.DefaultGrade.Value = format.ToMoodleFloat(float32(len(rawSubQuestions)))

	result.SubQuestions = []output.SubQuestion{}
	for _, q := range rawSubQuestions {
		c := output.SubQuestion{}
		textContent := format.ToMoodleParagraph(q.Answer)
		answerTextContent := format.ToMoodleOptionName(q.Question, subQuestionAnswerMatcher)

		c.Format = "html"
		c.Text.Content = textContent
		c.Answer.Text.Content = answerTextContent
		result.SubQuestions = append(result.SubQuestions, c)
	}

	for _, option := range *optionMap {
		if (option.Used) == true {
			continue
		}

		c := output.SubQuestion{}
		c.Format = "html"
		c.Text.Content = ""
		c.Answer.Text.Content = option.Value
		result.SubQuestions = append(result.SubQuestions, c)
	}
}

func parseOption(l *input.Line, optionMap *map[string]SQOption, sortedOptions *[]string) {
	optionParts := optionDocFormattingMatcher.FindStringSubmatch(l.Style)
	if len(optionParts) == 0 {
		return
	}

	sqKey := optionParts[1]
	(*optionMap)[sqKey] = SQOption{
		Value: l.Content,
		Used:  false,
	}
	(*sortedOptions) = append((*sortedOptions), l.Content)
}

func parseAnswer(l *input.Line, sqs *[]SQPair, optionMap *map[string]SQOption) {
	answerParts := answerDocFormattingMatcher.FindStringSubmatch(l.Style)
	if len(answerParts) == 0 {
		return
	}

	sqKey := answerParts[1]
	answerOption := (*optionMap)[sqKey].Value
	(*optionMap)[sqKey] = SQOption{
		Value: answerOption,
		Used:  true,
	}

	(*sqs) = append((*sqs), SQPair{
		Question: answerOption,
		Answer:   l.Content,
	})
}
