package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) String() string {
	var str string
	for _, p := range products {
		str += fmt.Sprintf("%s\n", p)
	}
	return str
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	}
}

//implementations for the sort.Interface interface

func (products Products) Len() int { return len(products) }
func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

//by name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//by name
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("\nSorting by Id (default)")
	sort.Sort(products)
	fmt.Println(products)

	fmt.Println("\nSorting by Name")
	//sort.Sort(ByName{products})
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println("\nSorting by Cost")
	//sort.Sort(ByCost{products})
	products.Sort("Cost")
	fmt.Println(products)

}

// implement the sort functionality that will allow the user to sort the products by any of the attributes (Id, Name, Cost, Units, Category)
// IMPORTANT : use sort.Sort() function
