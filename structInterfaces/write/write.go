package write

import (
	"os"
	"strings"
)

var file = "structPractice.txt"

func WriteText(title string, text []byte) error {

	fileName := strings.ReplaceAll(title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".json"

	os.WriteFile(fileName, []byte(text), 0644)

	return nil
}
