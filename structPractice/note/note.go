package note

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/note/write"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func GetNoteData() (string, string, error) {
	title := InputValue("Title   => ")

	content := InputValue("Content => ")
	return title, content, nil
}

func InputValue(prompt string) string {
	fmt.Print(prompt)

	// var value string
	// fmt.Scanln(&value)
	//we used bufio.newReader because we can write a long text with scanln only write one work o number
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\n")

	return text
}

func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("Invalid Input")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) OutPrint() error {
	// text := fmt.Sprint("\nTHE VALUES ARE:\nNote:   => %v \ncontent => %v\n", n.title, n.content)

	//
	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	fmt.Println(n)
	// write.WriteText(n.title, json)
	return write.WriteText(n.Title, json)

}
