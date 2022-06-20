package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func main() {
	var Book1 Books
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.book_id = 666
	printBook(Book1)
}
