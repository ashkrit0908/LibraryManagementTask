package test

import (
	"testing"

	"github.com/ashkrit0908/library_management_system/internal/library"
)

func TestAddBook(t *testing.T) {
	lib := library.NewLibrary()
	lib.AddBook("Book1", "Author1", 3)

	if len(lib.Books) != 1 {
		t.Errorf("Expected number of books to be 1, got %d", len(lib.Books))
	}

	book := lib.Books[0]
	if book.Title != "Book1" || book.Author != "Author1" || book.Quantity != 3 {
		t.Errorf("Book details mismatch after addition")
	}
}

func TestBorrowBook(t *testing.T) {
	lib := library.NewLibrary()
	lib.AddBook("Book1", "Author1", 3)

	// Borrow a book
	lib.BorrowBook("Book1")

	if lib.Borrowed["Book1"] != 1 {
		t.Errorf("Expected borrowed count for Book1 to be 1, got %d", lib.Borrowed["Book1"])
	}
}

func TestReturnBook(t *testing.T) {
	lib := library.NewLibrary()
	lib.AddBook("Book1", "Author1", 3)

	// Borrow a book
	lib.BorrowBook("Book1")

	// Return the borrowed book
	lib.ReturnBook("Book1")

	if lib.Borrowed["Book1"] != 0 {
		t.Errorf("Expected borrowed count for Book1 to be 0 after return, got %d", lib.Borrowed["Book1"])
	}
}
