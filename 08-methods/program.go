package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//Format and ApplyDiscount as methods
func (p Product) Format() string {
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

	/*
		//using the "Format" and "ApplyDiscount" as "functions"
		fmt.Println(Format(pen))
		ApplyDiscount(&pen, 10)
		fmt.Println(Format(pen))
	*/

	//using the "Format" and "ApplyDiscount" as "methods"
	fmt.Println(pen.Format())
	//(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())
}

//Format and ApplyDiscount as functions
/* func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(product *Product, discount float64) {
	product.Cost = product.Cost - (product.Cost * (discount / 100))
}
*/
