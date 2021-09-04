package program2

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
	return &Page{Title: title, Body: body}, nil
}

func ExecIO() {
	p1 := Page{Title: "TestPage", Body: []byte("This is sample body data")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println("File " + string(p2.Title) + " contains text " + string(p2.Body))
}
