package main

import (
	"module/mod"
)

func main() {

	book := new(mod.Book)

	mod.SearchAndList()

	mod.DeleteBook(book)

}
