package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ashkrit0908/library_management_system/internal/library"
)

func main() {
	library := library.NewLibrary()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Display Available Books")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. Search Book")
		fmt.Println("6. Display Borrowed Books")
		fmt.Println("7. Exit")

		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var title, author string
			var quantity int

			fmt.Print("Enter book title: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter author name: ")
			author, _ = reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Enter quantity available: ")
			quantityInput, _ := reader.ReadString('\n')
			quantityInput = strings.TrimSpace(quantityInput)
			quantity, _ = strconv.Atoi(quantityInput)

			library.AddBook(title, author, quantity)

		case "2":
			library.DisplayBooks()

		case "3":
			var title string
			fmt.Print("Enter the title of the book you want to borrow: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			library.BorrowBook(title)

		case "4":
			var title string
			fmt.Print("Enter the title of the book you want to return: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			library.ReturnBook(title)

		case "5":
			var query string
			fmt.Print("Enter the title or author to search: ")
			query, _ = reader.ReadString('\n')
			query = strings.TrimSpace(query)
			library.SearchBook(query)

		case "6":
			library.DisplayBorrowedBooks()

		case "7":
			fmt.Println("Exiting the library system.")
			return

		default:
			fmt.Println("Invalid choice. Please select again.")
		}
	}
}
