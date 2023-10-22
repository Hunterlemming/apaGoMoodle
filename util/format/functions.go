package format

import (
	"fmt"
	"regexp"
	"strings"
)

func ToMoodleFloat(num float32) string {
	return fmt.Sprintf("%.7f", num)
}

func ToMoodleParagraph(s string) string {
	return fmt.Sprintf("<![CDATA[<p>%s</p>]]>", strings.TrimSpace(s))
}

func ToMoodleQuestionText(mainText string, subQuestionNames []string) string {
	var fomattedSQNames []string
	for _, sqn := range subQuestionNames {
		fName := fmt.Sprintf("<strong>&nbsp;  %s</strong><br>", sqn)
		fomattedSQNames = append(fomattedSQNames, fName)
	}

	return fmt.Sprintf("<![CDATA[<p>%s</p><p>%s</p>]]>", strings.TrimSpace(mainText), strings.Join(fomattedSQNames, ""))
}

func ToMoodleOptionName(s string, matcher *regexp.Regexp) string {
	parts := matcher.FindStringSubmatch(s)
	if len(parts) == 3 {
		// "A) Good answer" -> "Good answer"
		return parts[2]
	}
	return s
}
