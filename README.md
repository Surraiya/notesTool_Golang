# notesTool
// This tool has functions with which we can manage short one-line notes.
// Thanks to the functions we defined earlier, the user can add, open, view and delete notes.

//Our project starts with library records "bufio", "fmt", "os", "strings" 
//This is needed to read from the user (bufio), output text to the console (fmt), work with files (os), and process string data (strings).

//Starting with the first function func check(err error), it handles errors by taking the error as an argumentnet and checking if it is undefined(nil).
//If the error is not nil (that is, an error occurred), the function prints an error message to the console and calls os.Exit(1), which means that the program is terminated with a return code of 1.

//function func loadCollection(name string) []string, has several steps in it:
// 1. File, err := os.OpenFile(name+".txt", os.O_WRONLY|os.O_TRUNC, 0666) this line opens or creates a file with a name, adding the extension .txt to it.
// 1.1. This line of code is typically used before writing to a file. For example, before writing new notes to a collection, or before saving a modified collection to a file. Opening a file in write mode with the os.O_TRUNC flag truncates the file to zero length, which is useful when overwriting an existing file.
// 2. Function creates a scanner (bufio.NewScanner(file)), which is used to read the file contents line by line.
// 3. "notes = append(notes, line)" - Reads each line from the file using the scanner and adds it to the notes line slice.
// 4. "	check(scanner.Err())" - Checks for an error that occurred when reading lines from a file (scanner.Err()) using the check function. If an error occurs, the program displays an error message and terminates execution.
// 5. "defer file.Close()" - File is closed using defer file.Close(). This ensures that the file will be closed after the function completes. 
// In brief, the function loads the contents of the collection from a text file, and if the file does not exist, creates it, returning notes as a slice of strings. 

//func saveCollection(name string, notes []string) 
// This function opens or creates a text file, truncates it if it exists, writes merged notes to it with newline delimiters, and handles errors in the process.
// In the context of the program, is intended to save the current state of the notes collection in a file. When the user adds, deletes or modifies notes, this function is called to update the notes collection file.

//func showNotes(notes []string)
// This showNotes function is designed to display notes in the console. The showNotes function displays all non-empty notes from the passed notes slice in the console, representing each note in the "index - text" format.
// Suppose a program manages a collection of notes, and the user can add, delete, or view notes. When the user wants to view the current notes, the program calls the showNotes function, which displays them on the screen for the user's convenience.

//func addNote() string
// The addNote function is designed to add a new note to the collection. The function asks the user for text input for a new note, reads this input from the standard input and returns the entered text as a string. 
// Functions such as addNote are commonly used in a program to interact with the user. For example, after calling this function, the text entered by the user can be added to a collection of notes using other functions such as append or saveCollection.

//func deleteNote() int
// The deleteNote function is designed to ask the user for the number of the note he wants to delete or enter 0 to cancel the delete operation. The function is used to interact with the user and get information from him about which note he wants to delete. 
//The entered value (note number) can then be used in other parts of the program to perform a delete or cancel operation depending on the user input.

//func showHelp()
// The showHelp function is designed to output information to the console on how to use the program. fmt.Println("Usage: ./notestool [COLLECTION]"). 
//This information describes the correct syntax for running the program and helps the user understand what arguments can be passed when the program is run.

//func selectOperation() int
//This selectOperation function is responsible for displaying a menu of operations in the console and prompting the user to select one of the operations.
//Function allows the user to select operations from the menu, view the current collection of notes, add new notes, delete notes, view collections and exit the program.

//func getCollection() string
// The getCollection function is designed to ask the user to enter a collection name and returns the name entered by the user as a string.
// The function allows the user to enter a collection name in the program, and this collection name is used in further execution of the program to create, view or delete this note

//func showCollections()
// The showCollections function is intended to display in the console the list of available collections of notes in the current working directory of the program.
// This function displays in the console a list of available note collections in the current working directory of the program, providing the user with information about which collections he can select or switch.

//func changeCollection() string
// This changeCollection function allows the user to change the current collection of notes in the program.
// brief explanation of how the function works
// 1. showCollections(): Outputs a list of available collections to the console.
// 2. Requests the user to select an existing collection (by number) or create a new one.
// 3. If an existing collection is selected, changes the current collection to the selected collection.
// 4. If you choose to create a new collection, prompts the user for a new name and loads an empty collection with that name.
// 5. If an invalid entry is made, displays an error message.
// 6. Returns the name of the current collection after changing or creating it.


//func main()
// Brief Explanation:
// 1. Greeting the user. ("fmt.Println("Welcome to the notes tool!")")
// 2. Gets the name of the current collection of notes. ("collection := getCollection()")
// 3. Loads notes from a text file. ("notes := loadCollection(collection)")
// 4. Starts an infinite loop for user's choice of operations. ("for { choice := selectOperation()")
// 5. Depending on the user's selection, performs operations such as view, add, delete notes, display collections, and change the current collection. ("switch choice {case 1: ... case 6:")
// 6. The program runs until the user chooses to exit the program. ("case 6: fmt.Println("See you later!")os.Exit(0)")
// 7. If an error occurs, the message "invalid choise" is displayed in the user console.


// The data in this code is stored in files with ".txt" extension. Each collection of notes is a text file, where each line contains one note
// It goes like this:
// 1. loadCollection(name string) []string: This function is responsible for loading notes from a file. It opens a file named "name.txt" and reads each line as a separate note. The notes are stored as a slice of strings ([]string), which is returned from the function.
// 2. saveCollection(name string, notes []string): This function saves a collection of notes to a file. It opens a file named "name.txt", combines the notes into a single text separated by newline characters, and writes that text to the file.
// 3. notes []string: In the main function, notes are loaded into the notes variable using loadCollection(collection). The entire for loop in main revolves around this notes slice, changing it as notes are added, deleted, or viewed.