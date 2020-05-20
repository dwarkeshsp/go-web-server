package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, err
}

func main() {
	p1 := &Page{"test page", []byte("we wish you a merry christmas")}
	p1.save()
	p2, _ := loadPage("test page")
	fmt.Println(string(p2.Body))
}
