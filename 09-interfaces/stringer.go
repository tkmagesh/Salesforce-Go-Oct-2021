package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//implementation of the fmr.Stringer interface
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (product *Product) ApplyDiscount(discount float64) {
	product.Cost = product.Cost - (product.Cost * (discount / 100))
}

func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    50,
		Category: "Stationery",
	}

	fmt.Println(pen)
	pen.ApplyDiscount(10)
	fmt.Println(pen)
}
