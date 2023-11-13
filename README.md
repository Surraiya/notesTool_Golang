# Notes Tool

The Notes Tool is a command-line application built in Go, designed for efficiently managing short, single-line notes. Users can create a new collection or retrieve an existing collection, view existing notes, add new notes, or remove notes from a collection.

## Usage

To use the tool, run the following command in the terminal:

First, You need to build the notesTool file into an executable file. 

```bash
$ go build
```

Then give two argument where first argument would be the name of the executable file and second one would be the name of the collection you want to create or retrieve.

```bash
$ ./notestool [Tag]
```

if you give any wrong argument or if you write help then the project will display usage.

```bash
$ ./notestool help
Usage: ./todotool [TAG]
```

Otherwise, when you will give the accurate argument then the program will welcome you and show you all the available menus so that you can add, show or remove notes in the given collection. Until you exit, the menus will keep prompting.

```bash
$ ./notestool testtag
Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
```

## Operations Example

### Show Notes: 
If you use option 1, It will display all notes in the collection. If there is no existing notes in the collection then It will suggest to add new note using option 2.
```bash
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1
No notes in the collection. Add a note using option 2.
```
If there are existing notes then:
```bash
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - This is the first note in the collection testtag
002 - This is the second note in the collection testtag
003 - This is the third note in the collection testtag
```

### Add a Note: 
If you use option 2, It will add a new note to the collection.
```bash
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2
Enter the note text: This is the first note in the collection testtag
Note added: 001
```

### Delete a Note: 
If you use option 3, It will remove a specific note from the collection.
```bash
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
3
Enter the number of the note to remove or 0 to cancel: 
002
Note with ID 002 is successfully removed.
```
The note will be successfully deleted. You can check that using option 1. 
```bash
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - This is the first note in the collection testtag
003 - This is the third note in the collection testtag
```

### Exit: 
If you use option 4, It will save the given collection and then execute the program.
```bash
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
4
testtag Collection is saved successfully! 
Exiting Notes Tool. Goodbye!
```

## Data Storage

The tool organizes and stores notes as well as collections in a structured JSON format. For each collection, there is a dedicated JSON file that includes information about the collection's name. Inside this file, an array of notes is maintained, where each note is uniquely identified by an ID and accompanied by its respective text. Both notes and collections adhere to a structured format, with notes within a collection presented as an array of Note structures for clarity and coherence.

```go
type Note struct {
	Id   int
	Text string
}

type Collection struct {
	Name  string
	Notes []Note
}

```