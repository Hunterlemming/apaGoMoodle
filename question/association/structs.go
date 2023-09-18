package association

import "fmt"

type SQPair struct {
	Question string
	Answer   string
}

func (p SQPair) String() string {
	return fmt.Sprintf("[Question]: %v; [Answer]: %v", p.Question, p.Answer)
}
