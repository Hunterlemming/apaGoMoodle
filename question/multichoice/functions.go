package multichoice

import (
	"strconv"
	"strings"

	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/util/format"
	"GoMoodle/util/linetype"
)

func Parse(q *input.RawQuestion) *output.Question {
	result := &output.Question{}

	setDefaults(result)
	qType := parseHeader(q, result)
	parseAnswers(&(q.Lines), result, qType)

	return result
}

func setDefaults(result *output.Question) {
	result.Type = "multichoice"
	result.GeneralFeedback.Format = "html"
	result.DefaultGrade.Value = format.ToMoodleFloat(float32(1))
	result.Penalty.Value = format.ToMoodleFloat(float32(1.0 / 3.0))
	result.Hidden.Value = 0
	result.ShuffleAnswers.Value = true
	result.AnswerNumbering.Value = "abc"
	result.CorrectFeedback.Format = "html"
	result.PartiallyCorrectFeedback.Format = "html"
	result.IncorrectFeedback.Format = "html"
}

func parseHeader(q *input.RawQuestion, result *output.Question) QuestionType {
	headerParts := questionDocMatcher.FindStringSubmatch(q.HeaderText)
	questionText := q.HeaderText
	option := ""

	if len(headerParts) == 3 {
		questionText = headerParts[2]
		option = headerParts[1]
		if option[len(option)-1] != ' ' {
			option += " "
		}
	}

	qType := extractQuestionType(&questionText)

	result.Name = output.Name{
		Text: output.Text{
			Content: format.ToMoodleParagraph(strings.TrimSpace(option + questionText)),
		},
	}

	result.QuestionText = output.QuestionText{
		Format: "html",
		Text: output.Text{
			Content: format.ToMoodleParagraph(questionText),
		},
	}

	result.Single.Value = qType.Single

	return qType
}

func extractQuestionType(questionText *string) (qType QuestionType) {
	qType = QuestionType{
		Single: true,
	}

	typeSplit := strings.Split((*questionText), "*#")
	if len(typeSplit) <= 1 {
		return
	}

	(*questionText) = typeSplit[0]
	typeSplit = strings.Split(typeSplit[1], ".")
	goodAnswers, _ := strconv.Atoi(typeSplit[0])

	qType = QuestionType{
		Single:      false,
		Strict:      typeSplit[1] == "sz",
		GoodAnswers: goodAnswers,
	}

	return
}

func parseAnswers(lines *[]input.Line, result *output.Question, qType QuestionType) {
	result.Answers = []output.Answer{}
	for _, l := range *lines {
		a := output.Answer{}
		a.Text.Content = format.ToMoodleParagraph(l.Content)
		a.Format = "html"
		a.Feedback.Format = "html"
		parseAnswerType(&a, l.Style, qType)

		result.Answers = append(result.Answers, a)
	}
}

const numberOfChoices = 4

func parseAnswerType(a *output.Answer, lineStyle string, qType QuestionType) {
	if qType.Single {
		if lineStyle == linetype.ChoiceGood {
			a.Fraction = "100"
			return
		}

		a.Fraction = "0"
		return
	}

	// Multi, Good
	// goodFraction := float32(qType.GoodAnswers) / float32(numberOfChoices-qType.GoodAnswers) * 100
	goodFraction := float32(1) / float32(qType.GoodAnswers) * 100
	if lineStyle == linetype.ChoiceGood {
		a.Fraction = format.ToMoodleFloat(goodFraction)
		return
	}

	// Multi, Bad, Strict
	if qType.Strict {
		a.Fraction = "-100"
		return
	}

	// Multi, Bad, Non-Strict
	badFraction := float32(0.5)
	if qType.GoodAnswers == 1 {
		badFraction = float32(1)
	}
	if qType.GoodAnswers == 3 {
		badFraction = float32(1) / float32(3)
	}
	a.Fraction = format.ToMoodleFloat(-badFraction * 100)
}
