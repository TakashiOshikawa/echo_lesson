package main

import (
	"fmt"
)

type Author struct {
	id int
	name string
	age int
	link string
}

func newAuthor(id int, name string, age int, link string) *Author {
	a := new(Author)
	a.id = id
	a.name = name
	a.age = age
	a.link = link

	return a
}

type Article struct {
	id int
	title string
	body string
	author *Author
}

func newArticle(id int, title string, body string, author *Author) *Article {
	a := new(Article)
	a.id = id
	a.title = title
	a.body = body
	a.author = author

	return a
}

func main() {

	author1 := newAuthor(1, "name1", 11, "link1")
	article1 := newArticle(1, "title1", "body1", author1)

	fmt.Println(article1.author.name)
}
