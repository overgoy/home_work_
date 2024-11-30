package main

import "fmt"

type ComparisonMode int

const (
	ByYear ComparisonMode = iota
	BySize
	ByRate
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(id int, title, author string, year, size int, rate float64) *Book {
	return &Book{id: id, title: title, author: author, year: year, size: size, rate: rate}
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float64 {
	return b.rate
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

type BookComparator struct {
	mode ComparisonMode
}

func NewBookComparator(mode ComparisonMode) *BookComparator {
	return &BookComparator{mode: mode}
}

func (c *BookComparator) Compare(book1, book2 *Book) bool {
	switch c.mode {
	case ByYear:
		return book1.year > book2.year
	case BySize:
		return book1.size > book2.size
	case ByRate:
		return book1.rate > book2.rate
	default:
		return false
	}
}

func main() {
	book1 := NewBook(1, "Go Programming", "John Doe", 2020, 300, 4.5)
	book2 := NewBook(2, "Learning Go", "Jane Smith", 2021, 250, 4.7)

	comparatorByYear := NewBookComparator(ByYear)
	comparatorBySize := NewBookComparator(BySize)
	comparatorByRate := NewBookComparator(ByRate)

	fmt.Println("Сравнение по году:", comparatorByYear.Compare(book1, book2))
	fmt.Println("Сравнение по размеру:", comparatorBySize.Compare(book1, book2))
	fmt.Println("Сравнение по рейтингу:", comparatorByRate.Compare(book1, book2))
}
