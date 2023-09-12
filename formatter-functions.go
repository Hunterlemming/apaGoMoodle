package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

var test = &Quiz{
	Questions: []XQuestion{
		{
			Type: "matching",
			Name: Name{
				Text: Text{
					Content: "12_35-40 Asszociáció",
				},
			},
			QuestionText: QuestionText{
				Format: "html",
				Text: Text{
					Content: "<![CDATA[<p>Párosítsa az egyes állításokhoz a megfelelő anyagokat!</p>\n<p><strong>   SiO<sub>2</sub></strong><br /><strong>   HCl</strong><br /><strong>   Na</strong><br /><strong>   KBr</strong><br /><strong>   egyik sem</strong><br /></p>]]>",
				},
			},
			GeneralFeedback: GeneralFeedback{
				Format: "html",
				Text:   Text{},
			},
			DefaultGrade: DefaultGrade{
				Value: toMoodleFloat(5.0),
			},
			Penalty: Penalty{
				Value: toMoodleFloat(1 / 3),
			},
			Hidden: Hidden{
				Value: 0,
			},
			ShuffleAnswers: ShuffleAnswers{
				Value: false,
			},
			CorrectFeedback: CorrectFeedback{
				Format: "html",
				Text:   Text{},
			},
			PartiallyCorrectFeedback: PartiallyCorrectFeedback{
				Format: "html",
				Text:   Text{},
			},
			IncorrectFeedback: IncorrectFeedback{
				Format: "html",
				Text:   Text{},
			},
			SubQuestions: []SubQuestion{
				{
					Format: "html",
					Text: Text{
						Content: "<![CDATA[<p>atomrácsos kristály</p>]]>",
					},
					Answer: Answer{
						Text: Text{
							Content: "SiO2",
						},
					},
				},
				{
					Format: "html",
					Text: Text{
						Content: "<![CDATA[<p>szobahőmérsékleten gáz halmazállapotú</p>]]>",
					},
					Answer: Answer{
						Text: Text{
							Content: "HCl",
						},
					},
				},
				{
					Format: "html",
					Text: Text{
						Content: "<![CDATA[<p>puha, szilárd anyag</p>]]>",
					},
					Answer: Answer{
						Text: Text{
							Content: "Na",
						},
					},
				},
				{
					Format: "html",
					Text: Text{
						Content: "<![CDATA[<p>szilárd anyag, amely csak olvadt állapotban vezeti az elektromos áramot</p>]]>",
					},
					Answer: Answer{
						Text: Text{
							Content: "KBr",
						},
					},
				},
				{
					Format: "html",
					Text: Text{
						Content: "<![CDATA[<p>molekularácsában a részecskéket hidrogén-kötések rögzítik</p>]]>",
					},
					Answer: Answer{
						Text: Text{
							Content: "egyik sem",
						},
					},
				},
			},
		},
	},
}

func writeFile(fileName string) {
	header := []byte(xml.Header)
	out, _ := xml.MarshalIndent(test, " ", "  ")

	err := os.WriteFile(fileName, append(header, out...), 0644)
	if err != nil {
		panic(err)
	}
}

func toMoodleFloat(num float32) string {
	return fmt.Sprintf("%.7f", num)
}
