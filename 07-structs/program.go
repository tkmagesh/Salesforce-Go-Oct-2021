package main

import (
	"fmt"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {

	//var pen Product
	//pen := Product{}
	//pen := Product{100, "Pen", 10, 50, "Stationery"}
	/* pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    50,
		Category: "Stationery",
	} */

	//penPtr := &Product{}
	//penPtr := new(Product)

	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    50,
		Category: "Stationery",
	}

	fmt.Println(pen.Name)
	penPtr := &pen
	//fmt.Println((*penPtr).Name)
	fmt.Println(penPtr.Name)

	fmt.Println(Format(pen))
	//ApplyDiscount(pen, 10)
	ApplyDiscount(&pen, 10)
	fmt.Println("After applying discount")
	fmt.Println(Format(pen))

	pencil := Product{
		Id:       101,
		Name:     "Pencil",
		Cost:     5,
		Units:    50,
		Category: "Stationery",
	}

	anotherPencil := Product{
		Id:       101,
		Name:     "Pencil",
		Cost:     5,
		Units:    50,
		Category: "Stationery",
	}

	fmt.Println(pencil == anotherPencil) //=> ?
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(product *Product, discount float64) {
	product.Cost = product.Cost - (product.Cost * (discount / 100))
}
