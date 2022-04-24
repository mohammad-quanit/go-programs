package main

type Printer interface {
	print()
}

type List []Printer

func (l List) print() {
	for _, it := range l {
		it.print()
	}
}
