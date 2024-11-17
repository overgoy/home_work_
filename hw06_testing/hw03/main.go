package hw03

import (
	"fmt"
)

func GenerateBoard(size int) string {
	var board string
	for i := 0; i < size; i++ {
		for s := 0; s < size; s++ {
			if (i+s)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	return board
}

func main() {
	var size int
	fmt.Println("Введите размер доски:")
	fmt.Scan("%d", &size)

	board := GenerateBoard(size)
	fmt.Print(board)
}
