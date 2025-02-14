package main

import (
	"fmt"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayes interface {
// 	Display()
// }

type outputtable interface {
	// we created interface with names of other interfaces or we can create interface wi name of interface or name interface and new method.
	saver
	// displayes
	Display()
}

// type outputtable interface {
// 	Save()
// 	Display()
// }

type create interface {
	New()
}

func main() {

	printSomething("123")
	printSomething(123)
	printSomething(12.3)
	printSomething(false)

	var appNote *note.Note
	var appTodo *todo.Todo

	fmt.Println("New Note")
	title, content, err := note.GetNoteData()

	appNote, err = note.New(title, content)

	fmt.Println("New Todo")
	todoContent, err := todo.GetTodoData()

	if err != nil {
		fmt.Print("\n", err, "\n\n")
		return
	}

	appTodo, err = todo.New(todoContent)

	fmt.Println("\nInformation of Note")
	outputData(appNote)

	fmt.Println("\nInformation of Todo")
	outputData(appTodo)

	fmt.Printf("\n\n\n")
	printSomething(appTodo)


}

func outputData(data outputtable) {
	data.Display()
	outPrintData(data)
}

func outPrintData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Savin the note faile.")
		return err
	}

	fmt.Println("Saving the note succeded!")
	return nil
}
//alternatively, instead of "insterface{{}}" you can also weite "any" that's an alias for this special "any value is allowed" type.
// func printSomething(value interface{}) {
func printSomething(value any) {

	intVal, ok := value.(int)
	if ok{
		fmt.Println("integer: ", intVal)
		return
	}
	stringVal, ok := value.(string)
	if ok{
		fmt.Println("String: ", stringVal)
		return
	}
	boolVal, ok := value.(bool)
	if ok{
		fmt.Println("Boolean: ", boolVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok{
		fmt.Println("Double: ", floatVal)
		return
	}
	// switch value.(type){
	// case int:
	// 	fmt.Println("Integer:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// case bool:
	// 	fmt.Println("Boolean:", value)
	// case float64:
	// 	fmt.Println("Double:", value)
	// default:
	// 	fmt.Println("don't exist this kind of type")

	// }

}
