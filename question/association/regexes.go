package association

import "regexp"

var (
	optionDocFormattingMatcher *regexp.Regexp
	answerDocFormattingMatcher *regexp.Regexp
	subQuestionAnswerMatcher   *regexp.Regexp
)

const (
	optionDocFormattingRE = `^Vlaszthat([A-Z])$` // Ex.: VlaszthatA = option A
	answerDocFormattingRE = `^Vlasz([A-Z])$`
	subQuestionAnswerRE   = `^([A-Z])\) ([\p{L} ]+)$`
)

func InitRegexMatchers() {
	var err error

	optionDocFormattingMatcher, err = regexp.Compile(optionDocFormattingRE)
	if err != nil {
		panic(err)
	}

	answerDocFormattingMatcher, err = regexp.Compile(answerDocFormattingRE)
	if err != nil {
		panic(err)
	}

	subQuestionAnswerMatcher, err = regexp.Compile(subQuestionAnswerRE)
	if err != nil {
		panic(err)
	}
}
