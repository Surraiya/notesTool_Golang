package collection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Note represents a single note with an ID and text.
type Note struct {
	Id   int
	Text string
}

// Collection represents a collection of notes.
type Collection struct {
	Name  string
	Notes []Note
}

func LoadCollection(collectionName string) Collection {
	var currentCollection Collection
	data, err := os.ReadFile(collectionName)

	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, create a new collection
			currentCollection = Collection{Name: collectionName, Notes: []Note{}}
			return currentCollection
		}
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	//decode JSON-formatted data from a byte slice into a Go data structure.
	err = json.Unmarshal(data, &currentCollection)
	if err != nil {
		fmt.Println("Error unmarshaling collection: ", err)
		os.Exit(1)
	}
	return currentCollection
}

func ShowNotes(currentCollection *Collection) {
	if len(currentCollection.Notes) == 0 {
		fmt.Println("No notes in the collection. Add a note using option 2.")
		return
	}

	fmt.Println("\nNotes:")
	for _, note := range currentCollection.Notes {
		fmt.Printf("%03d - %s\n", note.Id, note.Text)
	}
}

func AddNote(reader *bufio.Reader, currentCollection *Collection) {
	fmt.Print("Enter the note text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if text != "" {
		id := len(currentCollection.Notes) + 1
		newNote := Note{Id: id, Text: text}
		currentCollection.Notes = append(currentCollection.Notes, newNote)
		fmt.Printf("Note added: %03d\n", id)
	} else {
		fmt.Println("Note text cannot be empty. Note not added.")
	}
}

func DeleteNote(currentCollection *Collection) {
	if len(currentCollection.Notes) == 0 {
		fmt.Println("No notes in the collection to delete.")
		return
	}

	fmt.Println("Enter the number of the note to remove or 0 to cancel: ")
	noteId := 0
	_, err := fmt.Scanln(&noteId)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid integer.")
		return
	}

	if noteId == 0 {
		fmt.Println("Operation canceled.")
		return
	}

	found := false
	for i, note := range currentCollection.Notes {
		if note.Id == noteId {
			currentCollection.Notes = append(currentCollection.Notes[:i], currentCollection.Notes[i+1:]...)
			fmt.Printf("Note with ID %03d is successfully removed.\n", noteId)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Note with ID %03d not found.\n", noteId)
	}
}

func SaveCollection(currentCollection *Collection) {
	data, err := json.MarshalIndent(currentCollection, "", "  ")

	if err != nil {
		fmt.Println("Error marshaling collection:", err)
		os.Exit(1)
	}

	err = os.WriteFile(currentCollection.Name, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
	fmt.Printf("%s Collection is saved successfully! \nExiting Notes Tool. Goodbye!\n", currentCollection.Name)
}
