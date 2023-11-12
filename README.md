# notesTool
Markdown:

This tool has functions with which we can manage short one-line notes. 
Thanks to the functions we defined earlier, the user can add, open, view and delete notes.

## Check function
1. This function handles errors. If an error occurred, the program terminates. 

## LoadCollection function 
1. This function creates a file with a name and adds .txt format to it in order to save this file in the database in the future.
2. In short, the function loads the contents of the collection from a text file, and if the file does not exist, creates it, returning the records as a fragment of strings.

## SaveCollection function 
1. This function opens or creates a text file, truncates it if it exists, writes merged notes to it with newline delimiters, and handles errors in the process. 
2. In the context of the program, is intended to save the current state of the notes collection in a file. When the user adds, deletes or modifies notes, this function is called to update the notes collection file.

## ShowNotes function 
1. This showNotes function is designed to display notes in the console. The showNotes function displays all non-empty notes from the passed notes slice in the console, representing each note in the "index - text" format. 
2. Suppose a program manages a collection of notes, and the user can add, delete, or view notes. When the user wants to view the current notes, the program calls the showNotes function, which displays them on the screen for the user s convenience.

## AddNotes function 
1. The addNote function is designed to add a new note to the collection. The function asks the user for text input for a new note, reads this input from the standard input and returns the entered text as a string. 
2. Functions such as addNote are commonly used in a program to interact with the user. For example, after calling this function, the text entered by the user can be added to a collection of notes using other functions such as append or saveCollection.

## DeleteNotes function 
1. The deleteNote function is designed to ask the user for the number of the note he wants to delete or enter 0 to cancel the delete operation. The function is used to interact with the user and get information from him about which note he wants to delete.
2. The entered value (note number) can then be used in other parts of the program to perform a delete or cancel operation depending on the user input.

## ShowHelp function 
1. The showHelp function is designed to output information to the console on how to use the program. fmt.Println("Usage: ./notestool [COLLECTION]"). This information describes the correct syntax for running the program and helps the user understand what arguments can be passed when the program is run.

## SelectOperation function
1. This selectOperation function is responsible for displaying a menu of operations in the console and prompting the user to select one of the operations. 
2. Function allows the user to select operations from the menu, view the current collection of notes, add new notes, delete notes, view collections and exit the program.

## GetCollection function 
1.  The getCollection function is designed to ask the user to enter a collection name and returns the name entered by the user as a string. 
2. The function allows the user to enter a collection name in the program, and this collection name is used in further execution of the program to create, view or delete this note

## ShowCollection functions
1. The showCollections function is intended to display in the console the list of available collections of notes in the current working directory of the program.
2. This function displays in the console a list of available note collections in the current working directory of the program, providing the user with information about which collections he can select or switch.

## ChangeCollection
1. This changeCollection function allows the user to change the current collection of notes in the program.

## Main function 
The main function consists of several items, a brief explanation of the function:
	 1. Greeting the user. ("fmt.Println("Welcome to the notes tool!")")
	 2. Gets the name of the current notes collection. ("collection := getCollection()") 
	 3. Loads notes from a text file. ("notes := loadCollection(collection)") 
	 4. Runs an infinite loop to select operations by the user. ("for { choice := selectOperation()") 
	 5. Depending on the user s choice, performs operations such as view, add, delete notes, display collections, and change the current collection. ("switch choice {case 1: ... case 6:") 
	 6. The program runs until the user decides to exit it. ("case 6: fmt.Println("See you later!")os.Exit(0)")") 
	 7. If an error occurs, the message "invalid choise" is printed to the user's console.


## Data Stotage 
Notes for each collection are stored in a separate text file (e.g., [COLLECTION].txt). Each line in the file represents one note. The tool reads and writes data to this file to maintain the state of the collection.

## Usage 
$ ./notestool testtag
Welcome to the notes tool!
### Enter the name of the collection: (waits for users input. Example "Cat")
### Select operation(Cat):
1. Show notes.
2. Add a note.
3. Delete a note.
4. Show collection.
5. Change collection.
6. Exit.
(Waits for users input. Example '2')

### Enter the note text:
(Waits for users input. Example 'meow')

Note added!

### Select operation(Cat):
1. Show notes.
2. Add a note.
3. Delete a note.
4. Show collection.
5. Change collection.
6. Exit.
(Waits for users input. Example '1')

Notes:
1 - meow

End of notes

### Select operation(Cat):
1. Show notes.
2. Add a note.
3. Delete a note.
4. Show collection.
5. Change collection.
6. Exit.
(Waits for users input. Example '3')

Enter the number of the note to delete or 0 to cancel:

(Waits for users input. Example '1')

Note deleted!
 
(if the user presses 0)

Deletion canceled.

### Select operation(Cat):
1. Show notes.
2. Add a note.
3. Delete a note.
4. Show collection.
5. Change collection.
6. Exit.
(Waits for users input. Example '4')

1 - Cat

End of collections

### Select operation(Cat):
1. Show notes.
2. Add a note.
3. Delete a note.
4. Show collection.
5. Change collection.
6. Exit.
(Waits for users input. Example '5')

Collestions:
1 - Cat
### Enter the number of the collection to switch to or 0 to create a new collection:
(Waits for users input. Example '1')

Collection changed

(if the user presses 0)
### Enter the name of the collection:
(Waits for users input. Example 'Dog')

### Select operation(Cat):
1. Show notes.
2. Add a note.
3. Delete a note.
4. Show collection.
5. Change collection.
6. Exit.
(Waits for users input. Example '6')

See you later!

