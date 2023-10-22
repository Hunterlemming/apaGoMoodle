package multichoice

import "fmt"

type QuestionType struct {
	Single      bool
	Strict      bool
	GoodAnswers int
}

func (t QuestionType) String() string {
	return fmt.Sprintf("[Single]: %v; [Strict]: %v; [GoodAnswers]: %v;", t.Single, t.Strict, t.GoodAnswers)
}
