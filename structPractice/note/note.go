package note

import (
	"errors"
	"fmt"
)

type Note struct {
	title   string
	content string
}

func getNoteData() (string, string, error) {
	title, err := inputValue("Note title: => ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err := inputValue("Note content => ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil
}

func inputValue(prompt string) (string, error) {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Those fields are required")
	}
	return value, nil
}
