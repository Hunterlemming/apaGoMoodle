package format

import "fmt"

func ToMoodleFloat(num float32) string {
	return fmt.Sprintf("%.7f", num)
}
