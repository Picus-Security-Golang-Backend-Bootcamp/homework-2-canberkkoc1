package mod

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var bookName = []string{"In Search of Lost Time", "Ulysses", "Don Quixote",
	"One Hundred Years of Solitude",
	"The Great Gatsby",
	"Moby Dick",
	"War and Peace",
	"Hamlet",
	"The Odyssey",
	"Madame Bovary",
	"The Divine Comedy"}

var authorName = []string{"Agatha Christie",
	"Alan Moore",
	"Albert Camus",
	"Aldous Huxley",
	"Alexander Mccall Smith",
	"Allen Ginsberg",
	"Anaïs Nin",
	"Anne Frank",
	"Anne Rice",
	"Antoine De Saint-Exupéry",
	"Anton Chekhov",
}

type Book struct {
	BookName string
	Author   string
	ID       int
	StockNo  int
	IsbnNo   int
	PageNo   int
	Price    int
	Stock    int
	IsDelete bool
}

func (book *Book) Init(BookName string, Author string, ID int, StockNo int, Price int, IsDelete bool) {

	book.BookName = BookName

	book.Author = Author

	book.StockNo = StockNo

	book.IsDelete = IsDelete

	book.ID = ID

	book.Price = Price

	book.IsbnNo = rand.Intn(100)

	book.PageNo = rand.Intn(100)

	book.Stock = rand.Intn(100)

}

func (b *Book) deleted(deleted bool) {

	b.IsDelete = deleted

}

type DeletedAble interface {
	deleted(deleted bool)
}

func (newBook Book) createBook() *Book {

	book := new(Book)

	for i := 0; i < len(bookName); i++ {

		book.Init(bookName[i], authorName[i], i+1, 12, 100, false)

	}

	return book
}

func DeleteBook(d DeletedAble) {

	args := os.Args

	book := new(Book)

	firstArg := strings.ToLower(args[1])

	if firstArg == "delete" {

		secondArg, _ := strconv.Atoi(strings.ToLower(args[2]))

		for i := 0; i < len(bookName); i++ {

			book.Init(bookName[i], authorName[i], i+1, 12, 100, false)

			if book.ID == secondArg {
				book.IsDelete = true
				fmt.Println(*book)
			} else if len(bookName) < secondArg {
				fmt.Println("invalid id")
				break

			}

		}

	}

}
