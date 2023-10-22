package multichoice

import (
	"regexp"
)

var questionDocMatcher *regexp.Regexp

const (
	questionDocRE = `^([\p{N}\p{L}]+[.)]* )(.*)$` // 2. Asdada; A) Asdada
)

func InitRegexMatchers() {
	var err error

	questionDocMatcher, err = regexp.Compile(questionDocRE)
	if err != nil {
		panic(err)
	}
}
