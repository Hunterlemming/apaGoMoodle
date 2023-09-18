package association

import "regexp"

var optionMatcher *regexp.Regexp
var answerMatcher *regexp.Regexp

const (
	optionRE = `^Vlaszthat([A-Z])$` // Ex.: VlaszthatA = option A
	answerRE = `^Vlasz([A-Z])$`
)

func InitRegexMatchers() {

	var err error

	optionMatcher, err = regexp.Compile(optionRE)
	if err != nil {
		panic(err)
	}

	answerMatcher, err = regexp.Compile(answerRE)
	if err != nil {
		panic(err)
	}
}
