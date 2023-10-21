package output

import (
	"encoding/xml"
	"os"
)

func WriteFile(fileName string, quiz *Quiz) {
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

		if i+4 < n && string(original[i:i+5]) == "&#xA;" {
			formatted = append(formatted, byte('\n'))
			i += 5
			continue
		}

		if i+4 < n && string(original[i:i+5]) == "&#34;" {
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
