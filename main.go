package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Book struct
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// EBook struct
type EBook struct {
	Book
	FileSize int
}

// BookInterface interface
type BookInterface interface {
	DisplayDetails()
	GetISBN() string
	GetTitle() string
}

// Implement DisplayDetails for Book
func (b Book) DisplayDetails() {
	fmt.Printf("Title: %s, Author: %s, ISBN: %s, Available: %t\n", b.Title, b.Author, b.ISBN, b.Available)
}

// Implement GetISBN for Book
func (b Book) GetISBN() string {
	return b.ISBN
}

// Implement GetTitle for Book
func (b Book) GetTitle() string {
	return b.Title
}

// Implement DisplayDetails for EBook
func (e EBook) DisplayDetails() {
	fmt.Printf("Title: %s, Author: %s, ISBN: %s, Available: %t, FileSize: %dMB\n", e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}

// Implement GetISBN for EBook
func (e EBook) GetISBN() string {
	return e.ISBN
}

// Implement GetTitle for EBook
func (e EBook) GetTitle() string {
	return e.Title
}

// Library struct
type Library struct {
	Books []BookInterface
}

// AddBook method
func (l *Library) AddBook(book BookInterface) error {
	// Check for duplicate ISBN
	for _, b := range l.Books {
		if b.GetISBN() == book.GetISBN() {
			return errors.New("book with the same ISBN already exists")
		}
	}
	// Add the book to the library
	l.Books = append(l.Books, book)
	return nil
}

// RemoveBook method
func (l *Library) RemoveBook(isbn string) error {
	for i, book := range l.Books {
		if book.GetISBN() == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

// SearchBookByTitle method
func (l *Library) SearchBookByTitle(title string) {
	found := false
	for _, book := range l.Books {
		if strings.Contains(strings.ToLower(book.GetTitle()), strings.ToLower(title)) {
			book.DisplayDetails()
			found = true
		}
	}
	if !found {
		fmt.Println("No books found with the given title.")
	}
}

// ListBooks method
func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	for _, book := range l.Books {
		book.DisplayDetails()
	}
}

func main() {
	library := &Library{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Add an EBook")
		fmt.Println("3. Remove a Book/EBook")
		fmt.Println("4. Search for Books by Title")
		fmt.Println("5. List all Books/EBooks")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choice := 0
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Enter ISBN: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)

			fmt.Print("Is the book available (true/false): ")
			var available bool
			fmt.Scanln(&available)

			book := Book{Title: title, Author: author, ISBN: isbn, Available: available}
			if err := library.AddBook(book); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book added successfully!")
			}

		case 2:
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Enter ISBN: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)

			fmt.Print("Enter File Size (in MB): ")
			var fileSize int
			fmt.Scanln(&fileSize)

			fmt.Print("Is the eBook available (true/false): ")
			var available bool
			fmt.Scanln(&available)

			ebook := EBook{Book: Book{Title: title, Author: author, ISBN: isbn, Available: available}, FileSize: fileSize}
			if err := library.AddBook(ebook); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("EBook added successfully!")
			}

		case 3:
			fmt.Print("Enter ISBN to remove: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)
			if err := library.RemoveBook(isbn); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book/EBook removed successfully!")
			}

		case 4:
			fmt.Print("Enter Title to search: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			library.SearchBookByTitle(title)

		case 5:
			library.ListBooks()

		case 6:
			fmt.Println("Exiting the Library Management System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}