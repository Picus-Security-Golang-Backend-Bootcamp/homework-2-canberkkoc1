package mod

import (
	"fmt"
	"os"
	"strings"
)

func SearchAndList() {

	book := new(Book)
	args := os.Args

	var bookOrAuthor string

	firstArg := strings.ToLower(args[1])

	if len(args) == 1 || (firstArg != "list" && firstArg != "search") {

		fmt.Println("Please enter 'search' <Book_name> or 'list' ")

	} else {

		if firstArg == "list" {

			for i := 0; i < len(bookName); i++ {

				book.Init(bookName[i], authorName[i], i+1, 12, 100, false)

				fmt.Println(*book)
			}
		}

		if firstArg == "search" {

			secondArg := strings.Join(args[2:], " ")

			secondArg = strings.ToLower(secondArg)

			var result bool = false

			for _, v := range bookName {

				if strings.Contains(strings.ToLower(v), secondArg) {
					result = true
					bookOrAuthor = v
					break
				} else {

					for _, v := range authorName {

						if strings.Contains(strings.ToLower(v), secondArg) {
							result = true
							bookOrAuthor = v
							break
						}

					}
				}

			}
			if result {
				fmt.Printf("found %v\n", bookOrAuthor)

			} else {
				fmt.Printf("not found %v", secondArg)

			}

		}

	}

}
