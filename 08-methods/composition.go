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

//Product "methods"
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (product *Product) ApplyDiscount(discount float64) {
	product.Cost = product.Cost - (product.Cost * (discount / 100))
}

type PerishableProduct struct {
	Expiry string
	Product
}

func main() {
	grapes := PerishableProduct{
		Product: Product{
			Id:       102,
			Name:     "Grapes",
			Cost:     50.0,
			Units:    10,
			Category: "Food",
		},
		Expiry: "2 Days",
	}

	//Accessing the attributes
	//fmt.Println(grapes.Product.Id)
	fmt.Println(grapes.Id)

	//The below becomes better when the "Format" and "ApplyDiscount" are implemented as methods of Product
	/*
		fmt.Println(Format(grapes.Product))
		ApplyDiscount(&grapes.Product, 10)
		fmt.Println(Format(grapes.Product))
	*/
	fmt.Println(grapes.Format())
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}
