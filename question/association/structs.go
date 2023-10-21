package association

import "fmt"

type SQOption struct {
	Value string
	Used  bool
}

func (o SQOption) String() string {
	return fmt.Sprintf("[Value]: %v; [Used]: %v", o.Value, o.Used)
}

type SQPair struct {
	Question string
	Answer   string
}

func (p SQPair) String() string {
	return fmt.Sprintf("[Question]: %v; [Answer]: %v", p.Question, p.Answer)
}
