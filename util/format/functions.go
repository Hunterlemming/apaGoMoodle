package format

import (
	"fmt"
	"regexp"
	"strings"
)

func ToMoodleFloat(num float32) string {
	formatted := fmt.Sprintf("%.7f", num)
	if len(formatted) >= 9 {
		return formatted[:9]
	}
	return formatted
}

func ToMoodleParagraph(s string) string {
	text := strings.TrimSpace(s)
	// text = replaceNewLines([]byte(text))

	textLines := strings.Split(text, "#xA;")
	text = strings.Join(textLines, "</p><p>")
	return fmt.Sprintf("<![CDATA[<p>%s</p>]]>", text)
}

// func replaceNewLines(textBytes []byte) string {
// 	i := 0
// 	n := len(textBytes)
// 	for {
// 		if i == n {
// 			break
// 		}
// 	}
// }

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
