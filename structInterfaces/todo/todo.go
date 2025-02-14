package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"example.com/note/write"
)

type Todo struct {
	Content string `json:"text"`
}

func GetTodoData() (string, error) {

	content := InputValue("Content => ")
	return content, nil
}

func InputValue(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\n")

	return text
}

func New(content string) (*Todo, error) {

	return &Todo{
		Content: content,
	}, nil
}

func (todo Todo) Save() error {
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	return write.WriteText("todo", json)
}

func (todo Todo) Display() {
	fmt.Printf("Your content is: \n\n%v\n\n", todo.Content)
}
