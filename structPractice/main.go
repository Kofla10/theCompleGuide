package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	var appNote *note.Note

	title, content, err := note.GetNoteData()

	appNote, err = note.New(title, content)
	if err != nil {
		fmt.Print("\n", err, "\n\n")
		return
	}

	appNote.OutPrint()

}
