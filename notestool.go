package main

import (
	"bufio"
	"fmt"
	collection "notesTool/Collection"
	"os"
)

func main() {

	switch {
	case len(os.Args) != 2:
		DisplayUsage()
	case os.Args[1] == "help":
		DisplayUsage()
	default:
		fmt.Println("Welcome to the notes tool!")
		collectionName := os.Args[1]
		currentCollection := collection.LoadCollection(collectionName)
		reader := bufio.NewReader(os.Stdin)

		for {
			DisplayMenu()
			option := 0
			fmt.Scanln(&option)

			switch option {
			case 1:
				collection.ShowNotes(&currentCollection)
			case 2:
				collection.AddNote(reader, &currentCollection)
			case 3:
				collection.DeleteNote(&currentCollection)
			case 4:
				collection.SaveCollection(&currentCollection)
				os.Exit(0)
			default:
				fmt.Println("Invalid option")
			}
		}
	}
}

func DisplayMenu() {
	fmt.Println("\nSelect operation:")
	fmt.Println("1. Show notes.")
	fmt.Println("2. Add a note.")
	fmt.Println("3. Delete a note.")
	fmt.Println("4. Exit.")
}

func DisplayUsage() {
	fmt.Println("Usage: ./notestool [TAG]")
	os.Exit(0)
}
