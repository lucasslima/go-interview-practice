// Package main contains the implementation for Challenge 9: RESTful Book Management API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// Book represents a book in the database
type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	ISBN          string `json:"isbn"`
	Description   string `json:"description"`
}

// BookRepository defines the operations for book data access
type BookRepository interface {
	GetAll() ([]*Book, error)
	GetByID(id string) (*Book, error)
	Create(book *Book) error
	Update(id string, book *Book) error
	Delete(id string) error
	SearchByAuthor(author string) ([]*Book, error)
	SearchByTitle(title string) ([]*Book, error)
}

// InMemoryBookRepository implements BookRepository using in-memory storage
type InMemoryBookRepository struct {
	idCounter int64
	books     map[string]*Book
	mu        sync.RWMutex
}

func (repo *InMemoryBookRepository) GetAll() ([]*Book, error) {
	books := make([]*Book, 0)
	repo.mu.Lock()
	for _, book := range repo.books {
		books = append(books, book)
	}
	repo.mu.Unlock()
	return books, nil
}

func (repo *InMemoryBookRepository) GetByID(id string) (*Book, error) {
	//TODO implement me
	repo.mu.Lock()
	if book, ok := repo.books[id]; ok {
		repo.mu.Unlock()
		return book, nil
	}
	repo.mu.Unlock()
	return nil, fmt.Errorf("book for id %s not present", id)
}

func (repo *InMemoryBookRepository) Create(book *Book) error {
	book.ID = uuid.NewString()
	// Make sure there's no collisions
	repo.mu.Lock()
	_, exists := repo.books[book.ID]
	for exists {
		book.ID = uuid.NewString()
		_, exists = repo.books[book.ID]
	}
	repo.books[book.ID] = book
	repo.mu.Unlock()
	return nil
}

func (repo *InMemoryBookRepository) Update(id string, book *Book) error {
	existingBook, err := repo.GetByID(id)
	if err != nil {
		return err
	}
	repo.mu.Lock()
	existingBook.Author = book.Author
	existingBook.Title = book.Title
	existingBook.PublishedYear = book.PublishedYear
	existingBook.ISBN = book.ISBN
	existingBook.Description = book.Description
	repo.mu.Unlock()
	return nil
}

func (repo *InMemoryBookRepository) Delete(id string) error {
	_, err := repo.GetByID(id)
	if err != nil {
		return err
	}
	repo.mu.Lock()
	delete(repo.books, id)
	repo.mu.Unlock()
	return nil
}

func (repo *InMemoryBookRepository) SearchByAuthor(author string) ([]*Book, error) {
	books := make([]*Book, 0)
	repo.mu.Lock()
	for _, book := range repo.books {
		if strings.Contains(book.Author, author) {
			books = append(books, book)
		}
	}
	repo.mu.Unlock()
	return books, nil
}

func (repo *InMemoryBookRepository) SearchByTitle(title string) ([]*Book, error) {
	books := make([]*Book, 0)
	repo.mu.Lock()
	for _, book := range repo.books {
		if strings.Contains(book.Title, title) {
			books = append(books, book)
		}
	}
	repo.mu.Unlock()
	return books, nil
}

// NewInMemoryBookRepository creates a new in-memory book repository
func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		idCounter: 0,
		books:     make(map[string]*Book),
	}
}

// Implement BookRepository methods for InMemoryBookRepository
// ...

// BookService defines the business logic for book operations
type BookService interface {
	GetAllBooks() ([]*Book, error)
	GetBookByID(id string) (*Book, error)
	CreateBook(book *Book) error
	UpdateBook(id string, book *Book) error
	DeleteBook(id string) error
	SearchBooksByAuthor(author string) ([]*Book, error)
	SearchBooksByTitle(title string) ([]*Book, error)
}

// DefaultBookService implements BookService
type DefaultBookService struct {
	repo BookRepository
}

func (bs *DefaultBookService) GetAllBooks() ([]*Book, error) {
	books, err := bs.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil

}
func (bs *DefaultBookService) GetBookByID(id string) (*Book, error) {
	book, err := bs.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil

}
func (bs *DefaultBookService) CreateBook(book *Book) error {
	if book.Author == "" {
		return fmt.Errorf("missing author")
	}
	if book.Title == "" {
		return fmt.Errorf("missing title")
	}
	err := bs.repo.Create(book)
	if err != nil {
		return err
	}
	return nil
}

func (bs *DefaultBookService) UpdateBook(id string, book *Book) error {
	err := bs.repo.Update(id, book)
	if err != nil {
		return err
	}
	return nil
}
func (bs *DefaultBookService) DeleteBook(id string) error {
	err := bs.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil

}
func (bs *DefaultBookService) SearchBooksByAuthor(author string) ([]*Book, error) {
	books, err := bs.repo.SearchByAuthor(author)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (bs *DefaultBookService) SearchBooksByTitle(title string) ([]*Book, error) {
	books, err := bs.repo.SearchByTitle(title)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// NewBookService creates a new book service
func NewBookService(repo BookRepository) *DefaultBookService {
	return &DefaultBookService{
		repo: repo,
	}
}

// Implement BookService methods for DefaultBookService
// ...

// BookHandler handles HTTP requests for book operations
type BookHandler struct {
	Service BookService
}

// NewBookHandler creates a new book handler
func NewBookHandler(service BookService) *BookHandler {
	return &BookHandler{
		Service: service,
	}
}

// HandleBooks processes the book-related endpoints
func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method to handle all book endpoints
	// Use the path and method to determine the appropriate action
	// Call the service methods accordingly
	// Return appropriate status codes and JSON responses
	log.Printf("Called path: %v", r.URL.Path)
	idPattern := regexp.MustCompile(`\w+`)
	switch requestPath := strings.TrimPrefix(r.URL.Path, "/api/books"); {
	case requestPath == "" || requestPath == "/":
		switch r.Method {
		case http.MethodGet:
			books, err := h.Service.GetAllBooks()
			if err != nil {
				http.Error(w, fmt.Sprintf("error listing books: %s", err), http.StatusInternalServerError)
			}
			log.Printf("Books: %v", books)
			json.NewEncoder(w).Encode(books)
			break
		case http.MethodPost:
			var bookData Book
			err := json.NewDecoder(r.Body).Decode(&bookData)
			if err != nil {
				http.Error(w, fmt.Sprintf("error creating book: %s", err), http.StatusBadRequest)
			}
			err = h.Service.CreateBook(&bookData)
			if err != nil {
				http.Error(w, fmt.Sprintf("error creating book: %s", err), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(bookData)

		}

	case requestPath == "/search":
		author := r.URL.Query().Get("author")
		title := r.URL.Query().Get("title")
		log.Printf("author = %s, title = %s", author, title)
		if author == "" && title == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		var books []*Book
		if author != "" {
			booksByAuthor, err := h.Service.SearchBooksByAuthor(author)
			if err != nil {
				http.Error(w, fmt.Sprintf("error searching books: %s", err), http.StatusInternalServerError)
				return
			}
			log.Printf("returned books by author: %v", booksByAuthor)
			books = append(books, booksByAuthor...)
		}
		if title != "" {
			booksByTitle, err := h.Service.SearchBooksByTitle(title)
			if err != nil {
				http.Error(w, fmt.Sprintf("error searching books: %s", err), http.StatusInternalServerError)
				return
			}
			log.Printf("returned booksByTitle: %v", booksByTitle)
			books = append(books, booksByTitle...)
		}
		json.NewEncoder(w).Encode(books)
	case idPattern.MatchString(requestPath):
		requestPath = strings.TrimPrefix(requestPath, "/")
		switch r.Method {
		case http.MethodDelete:
			err := h.Service.DeleteBook(requestPath)
			if err != nil {
				http.Error(w, fmt.Sprintf("book not found: %s", err), http.StatusNotFound)
				return
			}
			break
		case http.MethodGet:
			log.Printf("Getting %s", requestPath)
			book, err := h.Service.GetBookByID(requestPath)
			if err != nil {
				http.Error(w, fmt.Sprintf("error getting books: %s", err), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(book)
			break
		case http.MethodPut:
			log.Printf("Patching %s", requestPath)
			var bookData Book
			err := json.NewDecoder(r.Body).Decode(&bookData)
			if err != nil {
				http.Error(w, fmt.Sprintf("error creating book: %s", err), http.StatusBadRequest)
			}
			err = h.Service.UpdateBook(requestPath, &bookData)
			if err != nil {
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bookData)
			break

		}

		break

	default:
		log.Printf("Fell on the default case with %s", requestPath)
		http.Error(
			w, "Not found", http.StatusNotFound,
		)
		return
	}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Error      string `json:"error"`
}

// Helper functions
// ...

func main() {
	// Initialize the repository, service, and handler
	repo := NewInMemoryBookRepository()
	service := NewBookService(repo)
	handler := NewBookHandler(service)

	// Create a new router and register endpoints
	http.HandleFunc("/api/books", handler.HandleBooks)
	http.HandleFunc("/api/books/", handler.HandleBooks)

	// Start the server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
