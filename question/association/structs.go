package association

import "fmt"

type SQPair struct {
	OriginalIndex int
	Question      string
	Answer        string
}

func (p SQPair) String() string {
	return fmt.Sprintf("[OriginalIndex]: %v; [Question]: %v; [Answer]: %v", p.OriginalIndex, p.Question, p.Answer)
}
