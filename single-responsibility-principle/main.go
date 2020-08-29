package main

import (
	"fmt"
)

type book struct {
	title  string
	author string
}

var books []*book

// TODO fungsi untuk menambahkan buku
func (b *book) AddBook(bk *book) []*book {
	books = append(books, bk)
	return books
}

// TODO fungsi untuk menampilkan buku yang telah di tambahkan
func (b *book) display() {
	for _, v := range books {
		fmt.Println(fmt.Sprintf("title : %s, author : %s", v.title, v.author))
	}
}

func main() {
	b := &book{}
	b.AddBook(&book{"buku 1", "eko"})
	b.AddBook(&book{"buku 2", "eko 2"})
	b.display()
}
