package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	var note *note.Note

	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(title, content)
}
