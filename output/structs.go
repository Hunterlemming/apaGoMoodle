package output

import "encoding/xml"

type Text struct {
	XMLName xml.Name `xml:"text"`
	Content string   `xml:",chardata"`
}

type Name struct {
	XMLName xml.Name `xml:"name"`
	Text    Text     `xml:"text"`
}

type QuestionText struct {
	XMLName xml.Name `xml:"questiontext"`
	Format  string   `xml:"format,attr"`
	Text    Text     `xml:"text"`
}

type GeneralFeedback struct {
	XMLName xml.Name `xml:"generalfeedback"`
	Format  string   `xml:"format,attr"`
	Text    Text     `xml:"text"`
}

type DefaultGrade struct {
	XMLName xml.Name `xml:"defaultgrade"`
	Value   string   `xml:",chardata"`
}

type Penalty struct {
	XMLName xml.Name `xml:"penalty"`
	Value   string   `xml:",chardata"`
}

type Hidden struct {
	XMLName xml.Name `xml:"hidden"`
	Value   byte     `xml:",chardata"`
}

type Single struct {
	XMLName xml.Name `xml:"single"`
	Value   bool     `xml:",chardata"`
}

type ShuffleAnswers struct {
	XMLName xml.Name `xml:"shuffleanswers"`
	Value   bool     `xml:",chardata"`
}

type AnswerNumbering struct {
	XMLName xml.Name `xml:"answernumbering"`
	Value   string   `xml:",chardata"`
}

type CorrectFeedback struct {
	XMLName xml.Name `xml:"correctfeedback"`
	Format  string   `xml:"format,attr"`
	Text    Text     `xml:"text"`
}

type PartiallyCorrectFeedback struct {
	XMLName xml.Name `xml:"partiallycorrectfeedback"`
	Format  string   `xml:"format,attr"`
	Text    Text     `xml:"text"`
}

type IncorrectFeedback struct {
	XMLName xml.Name `xml:"incorrectfeedback"`
	Format  string   `xml:"format,attr"`
	Text    Text     `xml:"text"`
}

type Feedback struct {
	XMLName xml.Name `xml:"feedback"`
	Format  string   `xml:"format,attr"`
	Text    Text     `xml:"text"`
}

type Answer struct {
	XMLName  xml.Name `xml:"answer"`
	Fraction string   `xml:"fraction,attr"`
	Format   string   `xml:"format,attr"`
	Text     Text     `xml:"text"`
	Feedback Feedback `xml:"feedback"`
}

type SubQuestion struct {
	XMLName xml.Name `xml:"subquestion"`
	Format  string   `xml:"format,attr"`
	Text    Text     `xml:"text"`
	Answer  Answer   `xml:"answer"`
}

type Question struct {
	XMLName                  xml.Name                 `xml:"question"`
	Type                     string                   `xml:"type,attr"`
	Name                     Name                     `xml:"name"`
	QuestionText             QuestionText             `xml:"questiontext"`
	GeneralFeedback          GeneralFeedback          `xml:"generalfeedback"`
	DefaultGrade             DefaultGrade             `xml:"defaultgrade"`
	Penalty                  Penalty                  `xml:"penalty"`
	Hidden                   Hidden                   `xml:"hidden"`
	Single                   Single                   `xml:"single"`
	ShuffleAnswers           ShuffleAnswers           `xml:"shuffleanswers"`
	AnswerNumbering          AnswerNumbering          `xml:"answernumbering"`
	CorrectFeedback          CorrectFeedback          `xml:"correctfeedback"`
	PartiallyCorrectFeedback PartiallyCorrectFeedback `xml:"partiallycorrectfeedback"`
	IncorrectFeedback        IncorrectFeedback        `xml:"incorrectfeedback"`
	SubQuestions             []SubQuestion            `xml:"subquestion"`
	Answers                  []Answer                 `xml:"answer"`
}

type Quiz struct {
	XMLName   xml.Name   `xml:"quiz"`
	Questions []Question `xml:"question"`
}
