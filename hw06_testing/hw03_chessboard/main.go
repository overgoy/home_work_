package hw03_chessboard

import "fmt"

func main() {
	var size int
	fmt.Println("Введите размер доски")
	fmt.Scan("%d", size)

	for i := 0; i < size; i++ {
		for s := 0; s < size; s++ {
			if (i+s)%2 == 0 {
				fmt.Println(" ")
			} else {
				fmt.Println("#")
			}
		}
		fmt.Println()
	}
}
