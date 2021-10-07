package main

import "fmt"

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//iteration using indexer
	fmt.Println("iteration using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	//iteration using range
	fmt.Println("iteration using range")
	for _, value := range nos {
		fmt.Println(value)
	}

	//copy of an array
	var newNos [5]int
	newNos = nos
	fmt.Println(newNos)
	newNos[0] = 200
	fmt.Println(newNos, nos)

	//Slices
	fmt.Printf("\nSlice\n")

	/*
		var products []string
		fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))
		products = append(products, "Pen")
		fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))
		products = append(products, "Pencil")
		fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))
		products = append(products, "Marker", "Scribble-Pad", "Sharpner")
		fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))
	*/

	//var products []string = []string{"Pen", "Pencil", "Marker"}
	var products []string = make([]string, 2, 5)
	products[0] = "Pen"
	products[1] = "Pencil"
	//products[2] = "Marker"
	products = append(products, "Marker")
	//products = append(products, "Pen", "Pencil", "Marker")
	fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))
	products = append(products, "Scribble-Pad")
	fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))
	products = append(products, "Sharpner", "Stapler")
	fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))
	products = append(products, "Stapler-Pins")
	fmt.Printf("%#v, len = %d, cap = %d\n", products, len(products), cap(products))

	/*
		newProducts := append(products, "P1", "P2")
		fmt.Println(products, newProducts)

		products[0] = "Dummy Product"
		fmt.Println(products, newProducts)

		products = append(products, "Mouse")
		fmt.Println()
		fmt.Println(products, newProducts)
	*/
	/*
		//iteration using indexer
		fmt.Println("iteration using indexer")
		for idx := 0; idx < len(products); idx++ {
			fmt.Println(products[idx])
		}

		//iteration using range
		fmt.Println("iteration using range")
		for _, value := range products {
			fmt.Println(value)
		}
	*/

	//copy of a slice
	/*
		n := []int{1, 2, 3}
		n2 := make([]int, 3)
		copy(n2, n)
		fmt.Println(n, n2)
		n[0] = 100
		fmt.Println(n, n2)
	*/

	//slicing a slice
	/*
		[lo : hi] => from index lo to hi-1
		[lo : ] => from index lo to end
		[ : hi] => from 0 to hi-1
		[lo : lo] => []
		[lo : lo+1] => [lo]
		[:] => from start to end
	*/
	fmt.Println("Slicing a slice")
	fmt.Println(products)
	productsSlice := products[2:5]
	fmt.Println("products[2:5] => ", productsSlice)
	fmt.Println(cap(products), cap(productsSlice))

	//Maps

	fmt.Println("Maps")
	//productRanks := map[string]int{}
	//productRanks := make(map[string]int)
	productRanks := map[string]int{
		"Pen":    3,
		"Pencil": 1,
		"Marker": 2,
	}
	fmt.Println(productRanks)

	//accessing a value using the key
	fmt.Println("Rank of Pencil =>", productRanks["Pencil"])

	//adding a new key-value pair
	productRanks["Scribble-Pad"] = 5
	fmt.Println(productRanks)

	//iteration using range
	fmt.Println("iteration using range")
	for key, value := range productRanks {
		fmt.Println(key, value)
	}

	//checking if a key exists
	fmt.Println("Checking if a key exists")
	//rank, exists := productRanks["Pen"]
	//fmt.Println("Rank of Pen =>", rank, "Exists =>", exists)

	if rank, exists := productRanks["Stapler"]; exists {
		fmt.Println("Rank of Stapler =>", rank, "Exists =>", exists)
	} else {
		fmt.Println("Rank of Stapler =>", "Not Found")
	}

	//deleting a key-value pair
	fmt.Println("Deleting a key-value pair")
	delete(productRanks, "Pen")
	fmt.Println(productRanks)
}
