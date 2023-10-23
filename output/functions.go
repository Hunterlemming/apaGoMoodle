package output

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

func WriteFile(quiz *Quiz) {
	os.Mkdir(OutputDir, os.ModeDir)
	fileName := fmt.Sprint(OutputDir, "/", time.Now().Format("2006-01-02T15-04-05"), ".xml")

	header := []byte(xml.Header)
	content, _ := xml.MarshalIndent(*quiz, " ", "  ")

	out := formatOutput(append(header, content...))

	err := os.WriteFile(fileName, out, 0o644)
	if err != nil {
		panic(err)
	}
}

func formatOutput(original []byte) (formatted []byte) {
	i := 0
	n := len(original)

	for {
		if i == n {
			break
		}

		if i+3 < n && string(original[i:i+4]) == "&lt;" {
			formatted = append(formatted, byte('<'))
			i += 4
			continue
		}

		if i+3 < n && string(original[i:i+4]) == "&gt;" {
			formatted = append(formatted, byte('>'))
			i += 4
			continue
		}

		if i+3 < n && string(original[i:i+4]) == "#xA;" {
			formatted = append(formatted, byte('\n'))
			i += 4
			continue
		}

		if i+4 < n && string(original[i:i+5]) == "&#34;" {
			fmt.Println("yo")
			formatted = append(formatted, byte('"'))
			i += 5
			continue
		}

		if i+4 < n && string(original[i:i+5]) == "&amp;" {
			formatted = append(formatted, byte('&'))
			i += 5
			continue
		}

		formatted = append(formatted, original[i])
		i += 1
	}

	return formatted
}
