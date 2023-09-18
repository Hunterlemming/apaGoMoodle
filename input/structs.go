package input

import "fmt"

type Line struct {
	Style   string
	Content string
}

func (l Line) String() string {
	return fmt.Sprintf("%v -- (%v)\n", l.Content, l.Style)
}

type RawQuestion struct {
	Type       string
	HeaderText string
	Lines      []Line
}

func (q RawQuestion) String() string {
	return fmt.Sprintf("[HeaderText]: %v\n[Type]: %v\n[Line]:\n%v\n", q.HeaderText, q.Type, q.Lines)
}
