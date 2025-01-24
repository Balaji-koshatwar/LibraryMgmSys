***Library Management System in Go***

**Description :**
This is a simple Library Management System implemented in Go. It allows users to manage books and eBooks in a library. The system demonstrates the use of structs, slices, interfaces, and error handling in Go. The application is menu-driven and operates via a text-based interface.

**Features**

1. Add a Book to the library.
2. Add an eBook to the library.
3. Remove a Book or eBook using its ISBN.
4. Search for Books by title.
5. List all Books and eBooks.
6. Display additional details like eBook file sizes.
Prerequisites
Go installed on your system (version 1.16 or later).
Instructions to Run the Program
1. Clone the repository or download the source code.
2. Save the source code file as `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the program using the following command:
   ```
   go run main.go
   ```
5. Follow the on-screen menu to perform library management operations.

***MENU OPTIONS***

1. **Add a Book**
   - Enter the title, author, ISBN, and availability status.
2. **Add an eBook**
   - Enter the title, author, ISBN, file size (in MB), and availability status.
3. **Remove a Book/EBook**
   - Enter the ISBN of the book or eBook to remove.
4. **Search for Books by Title**
   - Enter a keyword or partial title to search for matching books.
5. **List All Books/EBooks**
   - Displays all books and eBooks in the library with their details.
6. **Exit**
   - Terminates the program.
Examples of Input/Output
### Example 1: Adding a Book
**Input:**
```
Enter Title: Go Programming Basics
Enter Author: John Doe
Enter ISBN: 123456789
Is the book available (true/false): true
```
**Output:**
```
Book added successfully!
```

### Example 2: Adding an EBook
**Input:**
```
Enter Title: Advanced Go Programming
Enter Author: Jane Smith
Enter ISBN: 987654321
Enter File Size (in MB): 5
Is the eBook available (true/false): true
```
**Output:**
```
EBook added successfully!
```

### Example 3: Searching for Books by Title
**Input:**
```
Enter Title to search: Go
```
**Output:**
```
Title: Go Programming Basics, Author: John Doe, ISBN: 123456789, Available: true
Title: Advanced Go Programming, Author: Jane Smith, ISBN: 987654321, Available: true, FileSize: 5MB
```

### Example 4: Listing All Books/EBooks
**Output:**
```
Title: Go Programming Basics, Author: John Doe, ISBN: 123456789, Available: true
Title: Advanced Go Programming, Author: Jane Smith, ISBN: 987654321, Available: true, FileSize: 5MB
```

### Example 5: Removing a Book/EBook
**Input:**
```
Enter ISBN to remove: 123456789
```
**Output:**
```
Book/EBook removed successfully!
```

### Example 6: Invalid ISBN during Removal
**Input:**
```
Enter ISBN to remove: 111111111
```
**Output:**
```
Error: book not found
```
**Error Handling :**

- The program checks for duplicate ISBNs when adding books or eBooks.
  
- Attempts to remove a non-existent book result in an error message.
  
- Invalid menu choices prompt the user to try again.
Files

- `main.go`: Contains the source code for the Library Management System.
  
**Author: Balaji Koshatwar**

**License :** This project is licensed under the MIT License. Feel free to use and modify it as needed.


















