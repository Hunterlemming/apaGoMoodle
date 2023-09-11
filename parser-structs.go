package main

import "fmt"

type Line struct {
	Style   string
	Content string
}

func (l Line) String() string {
	return fmt.Sprintf("%v -- (%v)\n", l.Content, l.Style)
}

type Question struct {
	Type    string
	Text    string
	Answers []Line
}

func (q Question) String() string {
	return fmt.Sprintf("[Text]: %v\n[Type]: %v\n[Answers]:\n%v\n", q.Text, q.Type, q.Answers)
}
