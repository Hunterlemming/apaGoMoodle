package multichoice

import (
	"strconv"
	"strings"

	"GoMoodle/input"
	"GoMoodle/output"
	"GoMoodle/util/format"
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
			Content: strings.TrimSpace(option + questionText),
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
		result.Answers = append(result.Answers, a)
	}
}
