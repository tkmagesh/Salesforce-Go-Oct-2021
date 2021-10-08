//buffered channels

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	go writeData(ch)
	fmt.Println("[@main] Attempting to read 10")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 10")
	fmt.Println("[@main] Attempting to read 20")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 20")
	fmt.Println("[@main] Attempting to read 30")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 30")
	fmt.Println("[@main] Attempting to read 40")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 40")
	fmt.Println("[@main] Attempting to read 50")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 50")
	fmt.Println("[@main] Attempting to read 60")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 60")
	fmt.Println("[@main] Attempting to read 70")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 70")
	fmt.Println("[@main] Attempting to read 80")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 80")
	fmt.Println("[@main] Attempting to read 90")
	fmt.Println(<-ch)
	fmt.Println("[@main] Completing reading 90")
	fmt.Println("Exiting main")

}

func writeData(ch chan int) {
	fmt.Println("	[@writeData] - Attempting to write 10")
	ch <- 10
	fmt.Println("	[@writeData] - Completed writing 10")
	fmt.Println("	[@writeData] - Attempting to write 20")
	ch <- 20
	fmt.Println("	[@writeData] - Completed writing 20")
	fmt.Println("	[@writeData] - Attempting to write 30")
	ch <- 30
	fmt.Println("	[@writeData] - Completed writing 30")
	fmt.Println("	[@writeData] - Attempting to write 40")
	ch <- 40
	fmt.Println("	[@writeData] - Completed writing 40")
	fmt.Println("	[@writeData] - Attempting to write 50")
	ch <- 50
	fmt.Println("	[@writeData] - Completed writing 50")
	fmt.Println("	[@writeData] - Attempting to write 60")
	ch <- 60
	fmt.Println("	[@writeData] - Completed writing 60")
	fmt.Println("	[@writeData] - Attempting to write 70")
	ch <- 70
	fmt.Println("	[@writeData] - Completed writing 70")
	fmt.Println("	[@writeData] - Attempting to write 80")
	ch <- 80
	fmt.Println("	[@writeData] - Completed writing 80")
	fmt.Println("	[@writeData] - Attempting to write 90")
	ch <- 90
	fmt.Println("	[@writeData] - Completed writing 90")
}
